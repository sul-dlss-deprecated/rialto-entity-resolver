package repository

import (
	"fmt"
)

// Repository is an interface that rialto-entity-resolver reads from as its source
type Repository interface {
	QueryForPersonByOrcid(orcid string) (*string, error)
	QueryForPersonByName(firstName string, lastName string) (*string, error)
	QueryForOrganizationByName(name string) (*string, error)
	QueryForTopicByName(name string) (*string, error)
}

// Service is the Neptune implementation of the repository
type Service struct {
	reader Reader
}

const vcardFn = "http://www.w3.org/2006/vcard/ns#fn"
const subject = "http://purl.org/dc/terms/subject"
const personType = "http://xmlns.com/foaf/0.1/Person"
const organizationType = "http://xmlns.com/foaf/0.1/Organization"
const topicType = "http://www.w3.org/2008/05/skos#Concept"
const orcidPredicate = "http://vivoweb.org/ontology/core#orcidId"
const prefLabel = "http://www.w3.org/2008/05/skos#prefLabel"

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
