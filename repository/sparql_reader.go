package repository

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/knakk/rdf"
	"github.com/knakk/sparql"
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
	_, err := r.Update(query)
	log.Println("Inserted data")
	return err
}

func (r *SparqlReader) endpoint() string {
	return os.Getenv("SPARQL_ENDPOINT")
}

// Update does a SPARQL update
// See https://github.com/knakk/sparql/issues/9
func (r *SparqlReader) Update(q string) ([]rdf.Triple, error) {
	form := url.Values{}
	form.Set("update", q)
	form.Set("format", "text/turtle")
	b := form.Encode()

	req, err := http.NewRequest(
		"POST",
		r.endpoint(),
		bytes.NewBufferString(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(b)))
	req.Header.Set("Accept", "text/turtle")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		var msg string
		if err != nil {
			msg = "Failed to read response body"
		} else {
			if strings.TrimSpace(string(b)) != "" {
				msg = "Response body: \n" + string(b)
			}
		}
		return nil, fmt.Errorf("Construct: SPARQL request failed: %s. "+msg, resp.Status)
	}
	dec := rdf.NewTripleDecoder(resp.Body, rdf.Turtle)
	return dec.DecodeAll()
}
