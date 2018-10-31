package repository

import (
	"fmt"
)

// Repository is an interface that rialto-entity-resolver reads from as its source
type Repository interface {
	QueryForPersonBySunetid(sunetid string) (*string, error)
	QueryForPersonByOrcid(orcid string) (*string, error)
	QueryForPersonByName(firstName string, lastName string) (*string, error)
	QueryForOrganizationByName(name string) (*string, error)
	QueryForTopicByName(name string) (*string, error)
	QueryForGrantByIdentifier(identifier string) (*string, error)
}

// Service is the Neptune implementation of the repository
type Service struct {
	reader Reader
}

const altLabel = "http://www.w3.org/2004/02/skos/core#altLabel"
const grantType = "http://vivoweb.org/ontology/core#Grant"
const orcidPredicate = "http://vivoweb.org/ontology/core#orcidId"
const organizationType = "http://xmlns.com/foaf/0.1/Organization"
const personType = "http://xmlns.com/foaf/0.1/Person"
const prefLabel = "http://www.w3.org/2004/02/skos/core#prefLabel"
const subject = "http://purl.org/dc/terms/subject"
const sunetidType = "http://sul.stanford.edu/rialto/context/identifiers/Sunetid"
const topicType = "http://www.w3.org/2004/02/skos/core#Concept"

// NewService creates a new Service instance
func NewService(reader Reader) Repository {
	return &Service{reader: reader}
}

// QueryForPersonBySunetid returns the Person's URI that has the given SUNet ID
func (m *Service) QueryForPersonBySunetid(sunetid string) (*string, error) {
	uri, err := m.reader.QueryByEntityTypeIdentifierTypeAndObject(
		personType,
		sunetidType,
		sunetid)

	if err != nil {
		return nil, err
	}

	return uri, nil
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
		altLabel,
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
		prefLabel,
		name)

	if err != nil {
		return nil, err
	}

	return uri, nil
}

// QueryForTopicByName returns the Topic's URI that has the given name
func (m *Service) QueryForTopicByName(name string) (*string, error) {
	uri, err := m.reader.QueryByTypePredicateAndObject(
		topicType,
		subject,
		name)

	if err != nil {
		return nil, err
	}

	return uri, nil
}

// QueryForGrantByIdentifier returns the Grant's URI that has the given identifier
func (m *Service) QueryForGrantByIdentifier(identifier string) (*string, error) {
	uri, err := m.reader.QueryByTypePredicateAndObject(
		grantType,
		altLabel,
		identifier)

	if err != nil {
		return nil, err
	}

	return uri, nil
}
