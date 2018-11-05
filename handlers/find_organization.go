package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sul-dlss/rialto-entity-resolver/generated/models"
	"github.com/sul-dlss/rialto-entity-resolver/generated/restapi/operations"
	"github.com/sul-dlss/rialto-entity-resolver/runtime"
)

// NewFindOrganization will query Neptune for a organization record, or create a new one if no existing record can be found.
func NewFindOrganization(registry *runtime.Registry) operations.FindOrganizationHandler {
	return &findOrganization{
		registry: registry,
	}
}

// findOrganization handles a request for finding & returning an entry
type findOrganization struct {
	registry *runtime.Registry
}

// Handle the retrieve resource request
func (d *findOrganization) Handle(params operations.FindOrganizationParams, principal interface{}) middleware.Responder {
	uri, err := d.registry.Repository.QueryForOrganizationByName(params.Name)

	if err != nil {
		return d.handleError(err)
	}
	if uri != nil {
		return operations.NewFindOrganizationOK().WithPayload(*uri)
	}

	problem := &models.Error{
		Title:  "Not Found",
		Detail: "Unable to find an organization matching the criteria you provided",
		Status: "404",
	}
	errors := []*models.Error{problem}

	return operations.NewFindOrganizationNotFound().WithPayload(&models.ErrorResponse{Errors: errors})
}

func (d *findOrganization) handleError(err error) middleware.Responder {
	log.Printf("%s", err)
	problem := &models.Error{
		Title:  "Server error",
		Detail: err.Error(),
		Status: "500",
	}
	errors := []*models.Error{problem}

	return operations.NewFindOrganizationInternalServerError().WithPayload(&models.ErrorResponse{Errors: errors})
}
