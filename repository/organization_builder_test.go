package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
)

func TestBuildOrganizationWithoutCountry(t *testing.T) {
	name := "Stanford University"
	params := operations.FindOrCreateOrganizationParams{Name: name}
	builder := &OrganizationBuilder{}
	person, err := builder.Build(params)
	assert.Nil(t, err)
	assert.Contains(t, person.Triples[0], "a <http://xmlns.com/foaf/0.1/Organization>")
	assert.Contains(t, person.Triples[1], "<http://www.w3.org/2008/05/skos#prefLabel> \"Stanford University\"")
}

func TestBuildOrganizationWithCountry(t *testing.T) {
	name := "Stanford University"
	country := "USA"
	params := operations.FindOrCreateOrganizationParams{Name: name, Country: &country}
	builder := &OrganizationBuilder{}
	person, err := builder.Build(params)
	assert.Nil(t, err)
	assert.Contains(t, person.Triples[0], "a <http://xmlns.com/foaf/0.1/Organization>")
	assert.Contains(t, person.Triples[1], "<http://www.w3.org/2008/05/skos#prefLabel> \"Stanford University\"")
	assert.Contains(t, person.Triples[2], "<http://purl.org/dc/terms/spatial> \"USA\"")
}
