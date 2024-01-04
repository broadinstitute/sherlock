ARG GO_VERSION='1.21'
ARG ALPINE_VERSION='3.16'

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build
ARG BUILD_VERSION='development'
WORKDIR /build/sherlock
ENV CGO_ENABLED=0
ENV GOBIN=/bin

COPY sherlock/go.mod sherlock/go.sum ./
COPY go-shared ../go-shared/
RUN go mod download && go mod verify

COPY sherlock ./
RUN go build -buildvcs=false -ldflags="-X 'github.com/broadinstitute/sherlock/go-shared/pkg/version.BuildVersion=${BUILD_VERSION}'" -o /bin/sherlock ./cmd/...

# FROM alpine:${ALPINE_VERSION} as runtime <-- use this if you hit issues
FROM gcr.io/distroless/static:nonroot as runtime

COPY --from=build /bin/sherlock /bin/sherlock
ENTRYPOINT [ "/bin/sherlock" ]
