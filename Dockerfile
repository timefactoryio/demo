# Build stage
FROM --platform=$BUILDPLATFORM golang:latest AS go_builder
WORKDIR /src

ARG TARGETARCH

COPY . ./

# Cache Go modules
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Build the binary
RUN --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /out/demo .

# Final stage
FROM scratch
WORKDIR /app

# Copy the Go binary
COPY --from=go_builder /out/demo /demo

USER 1001
ENTRYPOINT ["/demo"]