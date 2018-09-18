package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
)

// MockedRepository is a mocked object that implements the Repository interface
type MockedRepository struct {
	mock.Mock
}

func (m *MockedRepository) QueryForPersonByOrcid(orcid string) (*string, error) {
	args := m.Called(orcid)
	result := args.Get(0)
	if result != nil {
		return result.(*string), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockedRepository) QueryForPersonByName(firstName string, lastName string) (*string, error) {
	args := m.Called(firstName, lastName)
	result := args.Get(0)
	if result != nil {
		return result.(*string), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockedRepository) QueryForOrganizationByName(name string) (*string, error) {
	args := m.Called(name)
	result := args.Get(0)
	if result != nil {
		return result.(*string), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockedRepository) CreatePerson(ops operations.FindOrCreatePersonParams) (*string, error) {
	args := m.Called(ops)
	return args.Get(0).(*string), args.Error(1)
}

func (m *MockedRepository) CreateOrganization(ops operations.FindOrCreateOrganizationParams) (*string, error) {
	args := m.Called(ops)
	return args.Get(0).(*string), args.Error(1)
}

func TestBuildPersonWithFullname(t *testing.T) {
	fn := "Wilson, Jennifer L."
	orcid := "0000-0002-2328-2018"
	params := operations.FindOrCreatePersonParams{FullName: &fn, Orcid: &orcid}
	builder := &PersonBuilder{}
	person, err := builder.Build(params)
	assert.Nil(t, err)
	assert.Contains(t, person.Triples[0], "a <http://xmlns.com/foaf/0.1/Person>")
	assert.Contains(t, person.Triples[1], "<http://www.w3.org/2006/vcard/ns#fn> \"Wilson, Jennifer L.\"")
	assert.Contains(t, person.Triples[2], "<http://vivoweb.org/ontology/core#orcidId> <https://orcid.org/0000-0002-2328-2018>")
}

func TestBuildPersonWithExistingOrganization(t *testing.T) {
	mockService := new(MockedRepository)
	id := "http://sul.stanford.edu/rialto/agents/organizations/8de0ce5e-a2a4-4e61-974e-df6c213cf220"
	mockService.On("QueryForOrganizationByName", "Ghent University").
		Return(&id, nil)
	fn := "Wilson, Jennifer L."
	orcid := "0000-0002-2328-2018"
	organization := "Ghent University"
	country := "BEL"
	params := operations.FindOrCreatePersonParams{FullName: &fn, Orcid: &orcid, Organization: &organization, Country: &country}
	builder := &PersonBuilder{Service: mockService}
	person, err := builder.Build(params)
	assert.Nil(t, err)
	assert.Contains(t, person.Triples[0], "a <http://xmlns.com/foaf/0.1/Person>")
	assert.Contains(t, person.Triples[1], "<http://www.w3.org/2006/vcard/ns#fn> \"Wilson, Jennifer L.\"")
	assert.Contains(t, person.Triples[2], "<http://vivoweb.org/ontology/core#orcidId> <https://orcid.org/0000-0002-2328-2018>")
}

func TestBuildPersonWithNameParts(t *testing.T) {
	lastName := "Wilson"
	firstName := "Jennifer L."
	orcid := "0000-0002-2328-2018"
	params := operations.FindOrCreatePersonParams{LastName: &lastName, FirstName: &firstName, Orcid: &orcid}
	builder := &PersonBuilder{}
	person, err := builder.Build(params)
	assert.Nil(t, err)
	assert.Contains(t, person.Triples[0], "a <http://xmlns.com/foaf/0.1/Person>")
	assert.Contains(t, person.Triples[1], "<http://www.w3.org/2006/vcard/ns#fn> \"Wilson, Jennifer L.\"")
	assert.Contains(t, person.Triples[2], "<http://vivoweb.org/ontology/core#orcidId> <https://orcid.org/0000-0002-2328-2018>")
}
