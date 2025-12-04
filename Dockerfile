# Build stage
FROM --platform=$BUILDPLATFORM golang:latest AS go_builder
WORKDIR /src

ARG TARGETARCH

# Copy only go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./

# Cache Go modules
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy main.go and build the binary
COPY main.go ./
RUN --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /out/demo .

# Final stage
FROM scratch
COPY --from=go_builder /out/demo /demo

USER 1001
ENTRYPOINT ["/demo"]