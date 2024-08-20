ARG GO_VERSION='1.22'

# https://github.com/microsoft/go-images/tree/microsoft/main
FROM mcr.microsoft.com/oss/go/microsoft/golang:${GO_VERSION}-fips-cbl-mariner2.0 AS build
ARG BUILD_VERSION='development'
WORKDIR /build/sherlock
ENV CGO_ENABLED=1

COPY sherlock/go.mod sherlock/go.sum ./
COPY go-shared ../go-shared/
RUN go mod download && go mod verify

COPY sherlock ./
RUN go build -buildvcs=false -ldflags="-X 'github.com/broadinstitute/sherlock/go-shared/pkg/version.BuildVersion=${BUILD_VERSION}'" -o /bin/sherlock ./cmd/...

# Check that the binary is FIPS-capable, fail if not
RUN go get github.com/acardace/fips-detect && \
    go run github.com/acardace/fips-detect /bin/sherlock \
    | grep -E 'FIPS-capable Go binary.*Yes'

# https://mcr.microsoft.com/en-us/product/cbl-mariner/distroless/minimal/about
#
# Broad appsec-blessed image from https://github.com/broadinstitute/dsp-appsec-blessed-images/blob/main/cbl-mariner/Dockerfile.2.0-nonroot-fips
FROM us.gcr.io/broad-dsp-gcr-public/base/cbl-mariner:2.0-nonroot-fips AS runtime

COPY --from=build /bin/sherlock /bin/sherlock
ENTRYPOINT [ "/bin/sherlock" ]
