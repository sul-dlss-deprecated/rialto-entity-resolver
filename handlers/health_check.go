package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/models"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
)

// NewHealthCheck will return the service health
func NewHealthCheck() operations.HealthCheckHandler {
	return &healthCheck{}
}

// findOrCreatePerson handles a request for finding & returning an entry
type healthCheck struct{}

// Handle the retrieve resource request
func (d *healthCheck) Handle(params operations.HealthCheckParams) middleware.Responder {
	return operations.NewHealthCheckOK().WithPayload(&models.HealthCheckResponse{Status: "OK"})
}
