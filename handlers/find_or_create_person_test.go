package handlers

import (
	"net/http"
	"os"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
	"github.com/sul-dlss-labs/rialto-entity-resolver/runtime"
)

// MockedRepo is a mocked object that implements the Repository interface
type MockedRepo struct {
	mock.Mock
}

func (m *MockedRepo) QueryForPersonByOrcid(orcid string) (*string, error) {
	args := m.Called(orcid)
	result := args.Get(0)
	if result != nil {
		return result.(*string), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockedRepo) QueryForPersonByName(firstName string, lastName string) (*string, error) {
	args := m.Called(firstName, lastName)
	result := args.Get(0)
	if result != nil {
		return result.(*string), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockedRepo) CreatePerson(ops operations.FindOrCreatePersonParams) (*string, error) {
	args := m.Called(ops)
	return args.Get(0).(*string), args.Error(1)
}

func TestLookupUserByName(t *testing.T) {
	os.Setenv("API_KEY", "abcdefg")
	r := gofight.New()
	repo := new(MockedRepo)
	id := "http://sul.stanford.edu/rialto/agents/people/8de0ce5e-a2a4-4e61-974e-df6c213cf220"
	repo.On("QueryForPersonByName", "Aaron", "Collier").
		Return(&id, nil)
	registry := &runtime.Registry{Repository: repo}
	r.GET("/person?last_name=Collier&first_name=Aaron").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
			})
}

func TestLookupUserByOrcid(t *testing.T) {
	r := gofight.New()
	repo := new(MockedRepo)
	id := "http://sul.stanford.edu/rialto/agents/people/8de0ce5e-a2a4-4e61-974e-df6c213cf220"
	repo.On("QueryForPersonByOrcid", "0000-0000-0000-0012").
		Return(&id, nil)
	registry := &runtime.Registry{Repository: repo}
	r.GET("/person?last_name=Collier&first_name=Aaron&orcid=0000-0000-0000-0012").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
			})
}

func TestCreateUser(t *testing.T) {
	r := gofight.New()
	repo := new(MockedRepo)
	repo.On("QueryForPersonByOrcid", "0000-0000-0000-0012").
		Return(nil, nil)
	repo.On("QueryForPersonByName", "Aaron", "Collier").
		Return(nil, nil)
	id := "http://sul.stanford.edu/rialto/agents/people/8de0ce5e-a2a4-4e61-974e-df6c213cf220"
	repo.On("CreatePerson", mock.Anything).
		Return(&id, nil)
	registry := &runtime.Registry{Repository: repo}
	r.GET("/person?last_name=Collier&first_name=Aaron&orcid=0000-0000-0000-0012").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
			})
}
