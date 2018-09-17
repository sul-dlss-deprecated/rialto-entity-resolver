package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
)

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
