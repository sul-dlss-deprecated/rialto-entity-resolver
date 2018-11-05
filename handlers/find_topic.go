package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sul-dlss/rialto-entity-resolver/generated/models"
	"github.com/sul-dlss/rialto-entity-resolver/generated/restapi/operations"
	"github.com/sul-dlss/rialto-entity-resolver/runtime"
)

// NewFindTopic will query Neptune for a topic record, or create a new one if no existing record can be found.
func NewFindTopic(registry *runtime.Registry) operations.FindTopicHandler {
	return &findTopic{
		registry: registry,
	}
}

// findTopic handles a request for finding & returning an entry
type findTopic struct {
	registry *runtime.Registry
}

// Handle the retrieve resource request
func (d *findTopic) Handle(params operations.FindTopicParams, principal interface{}) middleware.Responder {
	uri, err := d.registry.Repository.QueryForTopicByName(params.Name)

	if err != nil {
		return d.handleError(err)
	}
	if uri != nil {
		return operations.NewFindTopicOK().WithPayload(*uri)
	}

	problem := &models.Error{
		Title:  "Not Found",
		Detail: "Unable to find a topic matching the criteria you provided",
		Status: "404",
	}
	errors := []*models.Error{problem}

	return operations.NewFindTopicNotFound().WithPayload(&models.ErrorResponse{Errors: errors})
}

func (d *findTopic) handleError(err error) middleware.Responder {
	log.Printf("%s", err)
	problem := &models.Error{
		Title:  "Server error",
		Detail: err.Error(),
		Status: "500",
	}
	errors := []*models.Error{problem}

	return operations.NewFindTopicInternalServerError().WithPayload(&models.ErrorResponse{Errors: errors})
}
