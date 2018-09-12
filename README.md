# Rialto Entity Resolver

Use this service if you want to resolve some properties about an entity to a known URI.
If the entity is not found in RIALTO, one will be created.


## Generate server

```
swagger generate server -t generated --exclude-main
```

## Run Server

```
SPARQL_ENDPOINT=http://localhost:9999/blazegraph/namespace/kb/sparql go run cmd/server/main.go --port 3001
```

## Make a request
```
curl http://localhost:3001/person?last_name=Giarlo&first_name=Mike&orcid=0000-0002-2100-6108
```
