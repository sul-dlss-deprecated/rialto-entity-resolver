package repository

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
)

// OrganizationBuilder builds the triples that create an organization in the repository
type OrganizationBuilder struct {
}

// Build creates a Person which contains the triples that define it.
func (m *OrganizationBuilder) Build(params operations.FindOrCreateOrganizationParams) (*Entity, error) {
	u1 := uuid.NewV4()
	id := fmt.Sprintf("http://sul.stanford.edu/rialto/agents/organizations/%s", u1)
	person := &Entity{ID: id}
	triples := []string{
		fmt.Sprintf("<%s> a <%s>", id, organizationType),
		fmt.Sprintf("<%s> <%s> \"%s\"", id, prefLabel, params.Name),
	}
	if params.Country != nil {
		// TODO: URI this
		triples = append(triples, fmt.Sprintf("<%s> <%s> \"%s\"", id, spatial, *params.Country))
	}
	person.Triples = triples
	return person, nil
}
