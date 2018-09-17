# Rialto Entity Resolver

Use this service if you want to resolve some properties about an entity to a known URI.
If the entity is not found in RIALTO, one will be created.


## Generate server

```
rm -r generated/*
swagger generate server -t generated --exclude-main
```

## Run Server

```
API_KEY=abc123 SPARQL_ENDPOINT=http://localhost:9999/blazegraph/namespace/kb/sparql go run cmd/server/main.go --port 3001
```

## Make a request
```
curl http://localhost:3001/person?last_name=Giarlo&first_name=Mike&orcid=0000-0002-2100-6108
```

## Docker
### Build
```
docker build -t suldlss/rialto-entity-resolver:latest .
```
### Run
```
docker run -p 3000:3001 \
-e SPARQL_ENDPOINT=http://10.35.38.143:9999/blazegraph/namespace/kb/sparql \
-e API_KEY=<key> \
suldlss/rialto-entity-resolver:latest
```

### Deploy
```
docker push suldlss/rialto-entity-resolver:latest
```
