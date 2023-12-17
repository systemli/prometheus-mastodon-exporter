# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.21.4 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /prometheus-mastodon-exporter

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /prometheus-mastodon-exporter /prometheus-mastodon-exporter

EXPOSE 13120

USER nonroot:nonroot

ENTRYPOINT ["/prometheus-mastodon-exporter"]
