# Build the manager binary
FROM golang:1.22 AS builder
ARG TARGETOS
ARG TARGETARCH
ARG TARGETPLATFORM

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# since we use vendoring we don't need to redownload our dependencies every time. Instead we can simply
# reuse our vendored directory and verify everything is good. If not we can abort here and ask for a revendor.
COPY vendor vendor/
RUN go mod verify

# Copy the go source
COPY cmd cmd/
COPY api api/
COPY internal internal/

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
RUN GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o ksan-storage cmd/main.go

FROM --platform=$TARGETPLATFORM fedora:latest

RUN dnf update -y && \
    dnf clean all

WORKDIR /
COPY --from=builder /workspace/ksan-storage .
USER 65532:65532

ENTRYPOINT ["/ksan-storage"]
