ARG GO_VERSION='1.24'

# Non-root CBL-Mariner 2.0 runtime image
# https://mcr.microsoft.com/en-us/product/cbl-mariner/distroless/minimal/about
#
# When building amd64 images for deployed use, we typically use appsec-blessed images,
# like us.gcr.io/broad-dsp-gcr-public/base/cbl-mariner:2.0-nonroot-fips, defined at
# the time of building. See /.github/workflows/sherlock-build.yaml
# https://github.com/broadinstitute/dsp-appsec-blessed-images/blob/main/cbl-mariner/Dockerfile.2.0-nonroot-fips
#
# We do this swap because there aren't ARM-based blessed images (so for native local
# dev, we need to use a non-blessed image). The key is that whatever image is set
# here needs to be non-root and CBL-Mariner 2.0 with OpenSSL available.
ARG RUNTIME_BASE_IMAGE='mcr.microsoft.com/cbl-mariner/distroless/base:2.0-nonroot'

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

FROM ${RUNTIME_BASE_IMAGE} AS runtime

COPY --from=build /bin/sherlock /bin/sherlock
ENTRYPOINT [ "/bin/sherlock" ]
