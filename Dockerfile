# syntax=docker/dockerfile:1

# Build stage: compile the Go binary with modules cached.
FROM golang:1.23.4-bookworm AS build
WORKDIR /src

# Leverage go modules cache while keeping the final image small.
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Allow buildx to inject target platform info.
ARG TARGETOS=linux
ARG TARGETARCH=amd64
ENV CGO_ENABLED=0
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -trimpath -ldflags="-s -w" \
    -o /out/tencent-cert-updater ./

# Runtime stage: use a minimal distroless base with CA certificates.
FROM gcr.io/distroless/base-debian12 AS runtime
WORKDIR /app
COPY --from=build /out/tencent-cert-updater ./tencent-cert-updater

ENTRYPOINT ["/app/tencent-cert-updater"]
