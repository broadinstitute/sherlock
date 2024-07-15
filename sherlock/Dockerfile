ARG GO_VERSION='1.22'

FROM mcr.microsoft.com/oss/go/microsoft/golang:${GO_VERSION}-fips-cbl-mariner2.0
ARG BUILD_VERSION='development'
WORKDIR /build/sherlock
ENV CGO_ENABLED=1

COPY sherlock/go.mod sherlock/go.sum ./
COPY go-shared ../go-shared/
RUN go mod download && go mod verify

COPY sherlock ./
RUN go build -buildvcs=false -ldflags="-X 'github.com/broadinstitute/sherlock/go-shared/pkg/version.BuildVersion=${BUILD_VERSION}'" -o /bin/sherlock ./cmd/...

ENTRYPOINT [ "/bin/sherlock" ]
