package handlers

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"github.com/sul-dlss/rialto-entity-resolver/runtime"
)

func TestLookupGrantByIdentifier(t *testing.T) {
	os.Setenv("API_KEY", "abcdefg")
	r := gofight.New()
	repo := new(MockedRepo)
	id := "http://sul.stanford.edu/rialto/grants/000134677"
	repo.On("QueryForGrantByIdentifier", "abcd1234").
		Return(&id, nil)
	registry := &runtime.Registry{Repository: repo}
	r.GET("/grant?identifier=abcd1234").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
			})
}

func TestGrantNotFound(t *testing.T) {
	r := gofight.New()
	repo := new(MockedRepo)
	repo.On("QueryForGrantByIdentifier", "abcd1234").
		Return(nil, nil)
	registry := &runtime.Registry{Repository: repo}
	r.GET("/grant?identifier=abcd1234").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusNotFound, r.Code)
			})
}

func TestGrantServerError(t *testing.T) {
	r := gofight.New()
	repo := new(MockedRepo)
	repo.On("QueryForGrantByIdentifier", "abcd1234").
		Return(nil, errors.New("ooops"))
	registry := &runtime.Registry{Repository: repo}
	r.GET("/grant?identifier=abcd1234").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusInternalServerError, r.Code)
			})
}
