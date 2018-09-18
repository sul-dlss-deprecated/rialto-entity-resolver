package repository

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
)

// TopicBuilder builds the triples that create an topic in the repository
type TopicBuilder struct {
}

// Build creates a Topic which contains the triples that define it.
func (m *TopicBuilder) Build(params operations.FindOrCreateTopicParams) (*Entity, error) {
	u1 := uuid.NewV4()
	id := fmt.Sprintf("http://sul.stanford.edu/rialto/concepts/%s", u1)
	topic := &Entity{ID: id}
	triples := []string{
		fmt.Sprintf("<%s> a <%s>", id, topicType),
		fmt.Sprintf("<%s> <%s> \"%s\"", id, subject, params.Name),
	}
	topic.Triples = triples
	return topic, nil
}
