package handlers

import (
	"net/http"
	"os"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"github.com/sul-dlss/rialto-entity-resolver/runtime"
)

func TestLookupTopicByName(t *testing.T) {
	os.Setenv("API_KEY", "abcdefg")
	r := gofight.New()
	repo := new(MockedRepo)
	id := "http://sul.stanford.edu/rialto/concepts/8de0ce5e-a2a4-4e61-974e-df6c213cf220"
	repo.On("QueryForTopicByName", "Computer Science").
		Return(&id, nil)
	registry := &runtime.Registry{Repository: repo}
	r.GET("/topic?name=Computer+Science").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
			})
}

func TestTopicNotFound(t *testing.T) {
	r := gofight.New()
	repo := new(MockedRepo)
	repo.On("QueryForTopicByName", "Computer Science").
		Return(nil, nil)
	registry := &runtime.Registry{Repository: repo}
	r.GET("/topic?name=Computer+Science").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusNotFound, r.Code)
			})
}
