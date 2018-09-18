# RIALTO Entity Resolver

Use this service if you want to resolve some properties about an entity to a known URI. If the entity is not found in RIALTO, one will be created.

## Generate server

**NOTE**: Only do this when you change `swagger.yml`!

```
rm -r generated/*
swagger generate server -t generated --exclude-main
```

## Run Server

**NOTE**: the example value for `SPARQL_ENDPOINT` below assumes you're already running Blazegraph locally on port 9999 (default). To spin up Blazegraph as a local SPARQL endpoint, follow the [Blazegraph quick start](https://wiki.blazegraph.com/wiki/index.php/Quick_Start).

```
API_KEY=abc123 SPARQL_ENDPOINT=http://localhost:9999/blazegraph/namespace/kb/sparql go run cmd/server/main.go --port 3001
```

If this command bombs out, you may need to run `dep ensure` first to ensure you've got all the resolver's dependencies in place.

## Make a request

Make sure to pass in the API key specified in the `API_KEY` environment variable when the server was started.

```
curl -H "X-API-Key: abc123" "http://localhost:3001/person?last_name=Giarlo&first_name=Mike&orcid=0000-0002-2100-6108"
```

## Docker

### Build

```
docker build -t suldlss/rialto-entity-resolver:latest .
```

### Run

```
docker run -p 3000:3000 \
-e SPARQL_ENDPOINT=http://10.35.38.143:9999/blazegraph/namespace/kb/sparql \
-e API_KEY=<key> \
suldlss/rialto-entity-resolver:latest
```

### Deploy

```
docker push suldlss/rialto-entity-resolver:latest
```
