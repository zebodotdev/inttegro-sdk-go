# syntax=docker/dockerfile:1.7

FROM golang:1.23-alpine AS base
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

# Package-manager distribution artifact for Go modules is source + tags.
# This target runs checks and emits a source tarball in /out.
FROM base AS dist
RUN go test ./...
RUN mkdir -p /out && tar -czf /out/inttegro-sdk-go.tar.gz .

# CI target (use in GitHub Actions)
FROM base AS ci
RUN go test ./...

# Local/development target
FROM base AS dev
CMD ["sh"]
