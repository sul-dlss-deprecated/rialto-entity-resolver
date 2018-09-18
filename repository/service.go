package repository

import (
	"fmt"
	"log"

	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
)

// Repository is an interface that rialto-entity-resolver reads from as its source
type Repository interface {
	QueryForPersonByOrcid(orcid string) (*string, error)
	QueryForPersonByName(firstName string, lastName string) (*string, error)
	QueryForOrganizationByName(name string) (*string, error)
	CreatePerson(params operations.FindOrCreatePersonParams) (*string, error)
	CreateOrganization(params operations.FindOrCreateOrganizationParams) (*string, error)
}

// Service is the Neptune implementation of the repository
type Service struct {
	reader Reader
}

const vcardFn = "http://www.w3.org/2006/vcard/ns#fn"
const obo50 = "http://purl.obolibrary.org/obo/BFO_0000050"
const personType = "http://xmlns.com/foaf/0.1/Person"
const organizationType = "http://xmlns.com/foaf/0.1/Organization"
const orcidPredicate = "http://vivoweb.org/ontology/core#orcidId"
const prefLabel = "http://www.w3.org/2008/05/skos#prefLabel"
const spatial = "http://purl.org/dc/terms/spatial"

// NewService creates a new Service instance
func NewService(reader Reader) Repository {
	return &Service{reader: reader}
}

// QueryForPersonByOrcid returns the Person's URI that has the given OrcID
func (m *Service) QueryForPersonByOrcid(orcid string) (*string, error) {
	uri, err := m.reader.QueryByTypePredicateAndObject(
		personType,
		orcidPredicate,
		orcid)

	if err != nil {
		return nil, err
	}

	return uri, nil
}

// QueryForPersonByName returns the Person's URI that has the given name
func (m *Service) QueryForPersonByName(firstName string, lastName string) (*string, error) {
	uri, err := m.reader.QueryByTypePredicateAndObject(
		personType,
		vcardFn,
		fmt.Sprintf("%s, %s", lastName, firstName))

	if err != nil {
		return nil, err
	}

	return uri, nil
}

// QueryForOrganizationByName returns the Organization's URI that has the given name
func (m *Service) QueryForOrganizationByName(name string) (*string, error) {
	uri, err := m.reader.QueryByTypePredicateAndObject(
		organizationType,
		obo50,
		name)

	if err != nil {
		return nil, err
	}

	return uri, nil
}

// CreatePerson creates a new person with the given parameters
func (m *Service) CreatePerson(params operations.FindOrCreatePersonParams) (*string, error) {
	builder := &PersonBuilder{Service: m}
	person, err := builder.Build(params)
	if err != nil {
		return nil, err
	}
	log.Printf("Writing triples %v", person.Triples)
	if err = m.reader.Insert(person.Triples); err != nil {
		return nil, err
	}
	return &person.ID, nil
}

// CreateOrganization creates a new organization with the given parameters
func (m *Service) CreateOrganization(params operations.FindOrCreateOrganizationParams) (*string, error) {
	builder := &OrganizationBuilder{}
	person, err := builder.Build(params)
	if err != nil {
		return nil, err
	}
	log.Printf("Writing triples %v", person.Triples)
	if err = m.reader.Insert(person.Triples); err != nil {
		return nil, err
	}
	return &person.ID, nil
}
