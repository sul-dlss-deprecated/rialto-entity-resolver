package handlers

import (
	"net/http"
	"os"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"github.com/sul-dlss-labs/rialto-entity-resolver/runtime"
)

func TestLookupOrgByName(t *testing.T) {
	os.Setenv("API_KEY", "abcdefg")
	r := gofight.New()
	repo := new(MockedRepo)
	id := "http://sul.stanford.edu/rialto/agents/organizations/8de0ce5e-a2a4-4e61-974e-df6c213cf220"
	repo.On("QueryForOrganizationByName", "Stanford University").
		Return(&id, nil)
	registry := &runtime.Registry{Repository: repo}
	r.GET("/organization?name=Stanford+University").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
			})
}

func TestOrganizationNotFound(t *testing.T) {
	r := gofight.New()
	repo := new(MockedRepo)
	repo.On("QueryForOrganizationByName", "Stanford University").
		Return(nil, nil)

	registry := &runtime.Registry{Repository: repo}
	r.GET("/organization?name=Stanford+University").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusNotFound, r.Code)
			})
}
