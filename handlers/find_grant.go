package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sul-dlss/rialto-entity-resolver/generated/models"
	"github.com/sul-dlss/rialto-entity-resolver/generated/restapi/operations"
	"github.com/sul-dlss/rialto-entity-resolver/runtime"
)

// NewFindGrant will query Neptune for a grant record, or create a new one if no existing record can be found.
func NewFindGrant(registry *runtime.Registry) operations.FindGrantHandler {
	return &findGrant{
		registry: registry,
	}
}

// findGrant handles a request for finding & returning an entry
type findGrant struct {
	registry *runtime.Registry
}

// Handle the retrieve resource request
func (d *findGrant) Handle(params operations.FindGrantParams, principal interface{}) middleware.Responder {
	uri, err := d.registry.Repository.QueryForGrantByIdentifier(params.Identifier)

	if err != nil {
		return d.handleError(err)
	}
	if uri != nil {
		return operations.NewFindGrantOK().WithPayload(*uri)
	}

	problem := &models.Error{
		Title:  "Not Found",
		Detail: "Unable to find a grant matching the criteria you provided",
		Status: "404",
	}
	errors := []*models.Error{problem}

	return operations.NewFindGrantNotFound().WithPayload(&models.ErrorResponse{Errors: errors})
}

func (d *findGrant) handleError(err error) middleware.Responder {
	log.Printf("%s", err)
	problem := &models.Error{
		Title:  "Server error",
		Detail: err.Error(),
		Status: "500",
	}
	errors := []*models.Error{problem}

	return operations.NewFindGrantInternalServerError().WithPayload(&models.ErrorResponse{Errors: errors})
}
