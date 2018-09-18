package repository

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
)

// PersonBuilder builds the triples that create a person in the repository
type PersonBuilder struct {
	Service Repository
}

// Entity encapsulates all we know about an Entity: it's identifier and triples
type Entity struct {
	Triples []string
	ID      string
}

// Build creates a Person which contains the triples that define it.
func (m *PersonBuilder) Build(params operations.FindOrCreatePersonParams) (*Entity, error) {
	u1 := uuid.NewV4()
	id := fmt.Sprintf("http://sul.stanford.edu/rialto/agents/people/%s", u1)
	person := &Entity{ID: id}
	triples := []string{
		fmt.Sprintf("<%s> a <%s>", id, personType),
		fmt.Sprintf("<%s> <%s> \"%s\"", id, vcardFn, m.fullName(params)),
	}
	if params.Orcid != nil {
		triples = append(triples, fmt.Sprintf("<%s> <%s> <https://orcid.org/%s>", id, orcidPredicate, *params.Orcid))
	}
	if params.Organization != nil {
		orgURI, err := m.Service.QueryForOrganizationByName(*params.Organization)
		if err != nil {
			return nil, err
		}
		if orgURI == nil {
			orgParams := operations.FindOrCreateOrganizationParams{Name: *params.Organization, Country: params.Country}
			orgURI, err = m.Service.CreateOrganization(orgParams)
			if err != nil {
				return nil, err
			}
		}
		triples = append(triples, fmt.Sprintf("<%s> <%s> <%s>", id, obo50, *orgURI))

	}
	person.Triples = triples
	return person, nil
}

// Use either the passed in full_name or we will construct one by joining the first and last name.
func (m *PersonBuilder) fullName(params operations.FindOrCreatePersonParams) string {
	if params.FullName != nil {
		return *params.FullName
	}
	if params.FirstName != nil && params.LastName != nil {
		return fmt.Sprintf("%s, %s", *params.LastName, *params.FirstName)
	}
	panic("No names were provided")
}
