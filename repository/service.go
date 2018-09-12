package repository

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
)

// Repository is an interface that rialto-entity-resolver reads from as its source
type Repository interface {
	QueryForPersonByOrcid(orcid string) (*string, error)
	QueryForPersonByName(firstName string, lastName string) (*string, error)
	CreatePerson(params operations.FindOrCreatePersonParams) (*string, error)
}

// Service is the Neptune implementation of the repository
type Service struct {
	reader Reader
}

const vcardFn = "http://www.w3.org/2006/vcard/ns#fn"
const personType = "http://xmlns.com/foaf/0.1/Person"
const orcidPredicate = "http://vivoweb.org/ontology/core#orcidId"

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

// CreatePerson creates a new person with the given parameters
func (m *Service) CreatePerson(params operations.FindOrCreatePersonParams) (*string, error) {
	u1 := uuid.NewV4()
	id := fmt.Sprintf("http://sul.stanford.edu/rialto/agents/people/%s", u1)
	triples := []string{fmt.Sprintf("<%s> a <%s>", id, personType)}
	if params.Orcid != nil {
		triples = append(triples, fmt.Sprintf("<%s> <%s> <%s>", id, orcidPredicate, *params.Orcid))
	}
	if params.FirstName != nil && params.LastName != nil {
		triples = append(triples, fmt.Sprintf("<%s> <%s> \"%s, %s\"", id, vcardFn, *params.LastName, *params.FirstName))
	}
	err := m.reader.Insert(triples)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
