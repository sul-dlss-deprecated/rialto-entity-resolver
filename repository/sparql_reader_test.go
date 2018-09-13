package repository

import (
	"strings"
	"testing"

	"github.com/jcoyne/sparql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockedRepo is a mocked object that implements the SPARQLRepository interface
type MockedRepo struct {
	mock.Mock
}

func (f *MockedRepo) Query(q string) (*sparql.Results, error) {
	args := f.Called(q)
	return args.Get(0).(*sparql.Results), args.Error(1)
}

func (f *MockedRepo) Update(q string) error {
	return nil
}

func TestQueryByTypePredicateAndObject(t *testing.T) {
	fakeRepo := new(MockedRepo)

	institutionJSON := strings.NewReader(`{
	    "head": { "vars": [ "id" ] } ,
	    "results": {
	      "bindings": [
	        {
	          "id": { "type": "uri" , "value": "http://sul.stanford.edu/rialto/agents/people/8de0ce5e-a2a4-4e61-974e-df6c213cf220" }
	        }
	      ]
	    }
	  }`)
	fakeRepo.On("Query", "SELECT ?id\n\t\t\tWHERE {\n\t\t\t\t?id a <http://xmlns.com/foaf/0.1/Person> .\n\t\t\t\t?id <http://www.w3.org/2006/vcard/ns#fn> \"Giarlo, Mike\" .\n\t\t\t}").
		Return(sparql.ParseJSON(institutionJSON))

	reader := &SparqlReader{
		repo: fakeRepo,
	}
	result, _ := reader.QueryByTypePredicateAndObject("http://xmlns.com/foaf/0.1/Person", "http://www.w3.org/2006/vcard/ns#fn", "Giarlo, Mike")
	assert.Equal(t, *result, "http://sul.stanford.edu/rialto/agents/people/8de0ce5e-a2a4-4e61-974e-df6c213cf220")
}
