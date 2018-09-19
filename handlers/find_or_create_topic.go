package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
	"github.com/sul-dlss-labs/rialto-entity-resolver/runtime"
)

// NewFindOrCreateTopic will query Neptune for a topic record, or create a new one if no existing record can be found.
func NewFindOrCreateTopic(registry *runtime.Registry) operations.FindOrCreateTopicHandler {
	return &findOrCreateTopic{
		registry: registry,
	}
}

// findOrCreateTopic handles a request for finding & returning an entry
type findOrCreateTopic struct {
	registry *runtime.Registry
}

// Handle the retrieve resource request
func (d *findOrCreateTopic) Handle(params operations.FindOrCreateTopicParams, principal interface{}) middleware.Responder {
	uri, err := d.registry.Repository.QueryForTopicByName(params.Name)

	if err != nil {
		log.Printf("%s", err)
		panic(err)
	}
	if uri != nil {
		return operations.NewFindOrCreateTopicOK().WithPayload(*uri)
	}

	uri, err = d.registry.Repository.CreateTopic(params)
	if err != nil {
		panic(err)
	}
	return operations.NewFindOrCreateTopicOK().WithPayload(*uri)
}
