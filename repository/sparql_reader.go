package repository

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/knakk/sparql"
)

// Reader reads from the data store
type Reader interface {
	QueryByTypePredicateAndObject(entityType string, predicate string, object string) (*string, error)
	QueryByEntityTypeIdentifierTypeAndObject(entityType string, identifierType string, object string) (*string, error)
}

// SPARQLRepository is an interface we are making on the sparql library we are using,
// so that we can mock it in tests.
type SPARQLRepository interface {
	Query(q string) (*sparql.Results, error)
}

// SparqlReader represents the functions we do on the triplestore
type SparqlReader struct {
	repo SPARQLRepository
}

// NewSparqlReader creates a new instance of the sparqlReader for the provided endpoint
func NewSparqlReader(url string) *SparqlReader {
	repo, err := sparql.NewRepo(url,
		sparql.Timeout(time.Second*20),
	)
	if err != nil {
		log.Fatal(err)
	}
	return &SparqlReader{repo: repo}
}

// QueryByEntityTypeIdentifierTypeAndObject issues the query for the provided type with the given identifier type and object value
func (r *SparqlReader) QueryByEntityTypeIdentifierTypeAndObject(entityType string, identifierType string, object string) (*string, error) {
	query := fmt.Sprintf(`SELECT ?id
			WHERE {
				?id a <%s> .
				?id <http://purl.org/dc/terms/identifier> "%s"^^<%s> .
			}`, entityType, r.escapeLiteral(object), identifierType)
	results, err := r.repo.Query(query)
	if err != nil {
		log.Printf("[SPARQL] %s returned error: %v", query, err)
		return nil, err
	}

	if len(results.Solutions()) == 0 {
		log.Printf("[SPARQL] %s returned no results", query)
		return nil, nil
	}

	id := results.Solutions()[0]["id"].String()
	log.Printf("[SPARQL] %s returned %s from %d results", query, id, len(results.Solutions()))
	return &id, nil
}

// QueryByTypePredicateAndObject issues the query for the provided type with the given predicate and object assertion
func (r *SparqlReader) QueryByTypePredicateAndObject(entityType string, predicate string, object string) (*string, error) {
	query := fmt.Sprintf(`SELECT ?id
			WHERE {
				?id a <%s> .
				?id <%s> "%s" .
			}`, entityType, predicate, r.escapeLiteral(object))
	results, err := r.repo.Query(query)
	if err != nil {
		log.Printf("[SPARQL] %s returned error: %v", query, err)
		return nil, err
	}

	if len(results.Solutions()) == 0 {
		log.Printf("[SPARQL] %s returned no results", query)
		return nil, nil
	}

	id := results.Solutions()[0]["id"].String()
	log.Printf("[SPARQL] %s returned %s from %d results", query, id, len(results.Solutions()))
	return &id, nil
}

func (r *SparqlReader) escapeLiteral(literal string) string {
	return strings.Replace(literal, "\"", "\\\"", -1)
}
