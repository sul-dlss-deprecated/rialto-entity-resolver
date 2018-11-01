FROM golang:alpine as BASE
WORKDIR /go/src/github.com/sul-dlss/rialto-entity-resolver
COPY . .
RUN apk update && apk add --no-cache ca-certificates && \
    apk add --no-cache --virtual .build-deps git && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    apk del .build-deps
RUN CGO_ENABLED=0 GOOS=linux go build -o server -ldflags "-s" -a -installsuffix cgo cmd/server/main.go

FROM scratch
EXPOSE 8080
COPY --from=BASE /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BASE /go/src/github.com/sul-dlss/rialto-entity-resolver/server .
CMD ["./server"]