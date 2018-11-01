package runtime

import (
	"github.com/sul-dlss/rialto-entity-resolver/repository"
)

// Registry is the object that holds all the external services
type Registry struct {
	Repository repository.Repository
}

// NewRegistry creates a new instance of the service registry
func NewRegistry(repo repository.Repository) *Registry {
	return &Registry{
		Repository: repo,
	}
}
