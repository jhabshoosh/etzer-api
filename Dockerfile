ARG GO_VERSION=1.18
ARG ALPINE_VERSION=3.12

FROM golang:${GO_VERSION}-alpine AS builder

# ARG SEABOLT_VERSION=v1.7.4

# RUN apk add --update --no-cache ca-certificates cmake make g++ openssl-dev openssl-libs-static git curl pkgconfig libcap
# RUN git clone -b ${SEABOLT_VERSION} https://github.com/neo4j-drivers/seabolt.git /seabolt
# RUN update-ca-certificates 2>/dev/null || true

# WORKDIR /seabolt/build

# RUN cmake -D CMAKE_BUILD_TYPE=Release -D CMAKE_INSTALL_LIBDIR=lib .. && cmake --build . --target install

RUN mkdir -p /go/src/github.com/jhabshoosh/etzer-api
RUN mkdir /build
ADD . /go/src/github.com/jhabshoosh/etzer-api
WORKDIR /go/src/github.com/jhabshoosh/etzer-api

RUN go generate ./...

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /server cmd/main.go

# Create alpine runtime image
FROM alpine:${ALPINE_VERSION} as server

# Environment variables
ENV ETZER_DEBUG 'TRUE'
ENV ETZER_PORT '8080'
ENV ETZER_NEO4J_HOST 'localhost'
ENV ETZER_NEO4J_PORT '7687'
ENV ETZER_NEO4J_USER 'neo4j'
ENV ETZER_NEO4J_PASS 'test'

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /server /server

USER 1000

EXPOSE 80

ENTRYPOINT ["/server"]