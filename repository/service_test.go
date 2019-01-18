package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockedReader is a mocked object that implements the Reader interface
type MockedReader struct {
	mock.Mock
}

func (f *MockedReader) QueryByEntityTypeIdentifierTypeAndObject(entityType string, identifierType string, object string) (*string, error) {
	return nil, nil
}

func (f *MockedReader) QueryByTypePredicateAndObject(entityType string, predicate string, object string) (*string, error) {
	args := f.Called(entityType, predicate, object)
	if args.Get(0) != nil {
		return args.Get(0).(*string), nil
	}
	return nil, args.Error(1)

}

func TestQueryForPersonByName(t *testing.T) {
	fakeReader := new(MockedReader)

	uri := "http://sul.stanford.edu/rialto/agents/people/8de0ce5e-a2a4-4e61-974e-df6c213cf220"

	fakeReader.On("QueryByTypePredicateAndObject", personType, altLabel, "Littman, Justin").Return(nil, nil)
	fakeReader.On("QueryByTypePredicateAndObject", personType, altLabel, "littman, justin").Return(&uri, nil)

	service := NewService(fakeReader)

	result, _ := service.QueryForPersonByName("Justin", "Littman")
	assert.Equal(t, *result, uri)
	fakeReader.AssertExpectations(t)
}
