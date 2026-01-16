# Build stage
FROM --platform=$BUILDPLATFORM golang:latest AS go_builder
WORKDIR /src

ARG TARGETARCH

# Copy only go.mod and go.sum for dependency resolution
COPY main.go go.mod go.sum ./

# Cache Go modules
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# build the binary
RUN --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /out/demo .

# Copy CA certificates from builder
RUN mkdir -p /out/etc/ssl/certs && \
    cp /etc/ssl/certs/ca-certificates.crt /out/etc/ssl/certs/

# Final stage
FROM scratch
WORKDIR /app
COPY --from=go_builder /out/demo /demo
COPY --from=go_builder /out/etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY ./timefactory.svg /logo.svg
COPY ./slides /slides

USER 1001
ENTRYPOINT ["/demo"]