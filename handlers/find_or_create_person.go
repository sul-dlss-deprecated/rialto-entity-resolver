package handlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
)

// NewFindOrCreatePerson will query Neptune for a person record, or create a new one if no existing record can be found.
func NewFindOrCreatePerson() operations.FindOrCreatePersonHandler {
	return &findOrCreatePerson{}
}

// findOrCreatePerson handles a request for finding & returning an entry
type findOrCreatePerson struct {
}

// Handle the retrieve resource request
func (d *findOrCreatePerson) Handle(params operations.FindOrCreatePersonParams) middleware.Responder {
	name := swag.StringValue(params.LastName)

	greeting := fmt.Sprintf("http://rialto.stanford.org/person/%s", name)
	return operations.NewFindOrCreatePersonOK().WithPayload(greeting)
}
