package repository

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jcoyne/sparql"
)

// Reader reads from the data store
type Reader interface {
	QueryByTypePredicateAndObject(entityType string, predcate string, object string) (*string, error)
	Insert(triples []string) error
}

// SPARQLRepository is an interface we are making on the sparql library we are using,
// so that we can mock it in tests.
type SPARQLRepository interface {
	Query(q string) (*sparql.Results, error)
	Update(q string) error
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

// QueryByTypePredicateAndObject issues the query for the provided type with the given predicate and object assertion
func (r *SparqlReader) QueryByTypePredicateAndObject(entityType string, predicate string, object string) (*string, error) {
	query := fmt.Sprintf(`SELECT ?id
			WHERE {
				?id a <%s> .
				?id <%s> "%s" .
			}`, entityType, predicate, object)
	log.Printf("[SPARQL] %s", query)
	results, err := r.repo.Query(query)
	if err != nil {
		return nil, err
	}

	if len(results.Solutions()) == 0 {
		return nil, nil
	}

	id := results.Solutions()[0]["id"].String()
	return &id, nil
}

// Insert does a SPARQL update with the given triples
func (r *SparqlReader) Insert(triples []string) error {
	query := fmt.Sprintf(`INSERT DATA {
				%s
			}`, strings.Join(triples, ".\n"))
	err := r.repo.Update(query)
	log.Println("Inserted data")
	return err
}

func (r *SparqlReader) endpoint() string {
	return os.Getenv("SPARQL_ENDPOINT")
}
