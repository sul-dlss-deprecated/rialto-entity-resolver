package repository

import (
	"fmt"
	"os"
)

// BuildRepository instantiates the repository from the environment variables
func BuildRepository() Repository {
	endpoint := os.Getenv("SPARQL_ENDPOINT")
	if endpoint == "" {
		fmt.Fprintln(os.Stderr, "Required environment variable 'SPARQL_ENDPOINT' was not set.")
		os.Exit(0)
	}
	sparqlReader := NewSparqlReader(endpoint)
	return NewService(sparqlReader)
}
