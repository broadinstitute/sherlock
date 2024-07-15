ARG GO_VERSION='1.21'
ARG DEBIAN_VERSION='bookworm'

FROM mcr.microsoft.com/oss/go/microsoft/golang:${GO_VERSION}-fips-${DEBIAN_VERSION} AS build
ARG BUILD_VERSION='development'
WORKDIR /build/sherlock

COPY sherlock/go.mod sherlock/go.sum ./
COPY go-shared ../go-shared/
RUN go mod download && go mod verify

COPY sherlock ./
RUN go build -buildvcs=false -ldflags="-X 'github.com/broadinstitute/sherlock/go-shared/pkg/version.BuildVersion=${BUILD_VERSION}'" -o /bin/sherlock ./cmd/...

FROM mcr.microsoft.com/oss/go/microsoft/golang:${GO_VERSION}-fips-${DEBIAN_VERSION} AS runtime

COPY --from=build /bin/sherlock /bin/sherlock
ENTRYPOINT [ "/bin/sherlock" ]
