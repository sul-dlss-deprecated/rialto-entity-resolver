package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
)

func TestBuildTopic(t *testing.T) {
	name := "Computer Science"
	params := operations.FindOrCreateTopicParams{Name: name}
	builder := &TopicBuilder{}
	topic, err := builder.Build(params)
	assert.Nil(t, err)
	assert.Contains(t, topic.Triples[0], "a <http://www.w3.org/2008/05/skos#Concept>")
	assert.Contains(t, topic.Triples[1], "<http://purl.org/dc/terms/subject> \"Computer Science\"")
}
