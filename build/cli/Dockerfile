ARG GO_VERSION='1.16'
ARG ALPINE_VERSION='3.14'
ARG SHERLOCK_BUILD_VERSION='development'

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build
WORKDIR /build
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOBIN=/bin
COPY . .
RUN go test -short ./... && go build -ldflags="-X 'main.BuildVersion=${SHERLOCK_BUILD_VERSION}'" -o /bin/sherlock ./cmd/cli/...

FROM alpine:${ALPINE_VERSION} as runtime
COPY --from=build /bin/sherlock /bin/sherlock
ENTRYPOINT ["/bin/sherlock"]
