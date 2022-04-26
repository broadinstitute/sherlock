ARG GO_VERSION='1.18'
ARG ALPINE_VERSION='3.15'


FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build
ARG SHERLOCK_BUILD_VERSION='development'
WORKDIR /build
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOBIN=/bin
COPY . .
RUN go test -short ./... && go build -ldflags="-X 'main.BuildVersion=${SHERLOCK_BUILD_VERSION}'" -o /bin/sherlock ./cmd/cli/...

FROM alpine:${ALPINE_VERSION} as runtime
COPY --from=build /bin/sherlock /bin/sherlock
ENTRYPOINT ["/bin/sherlock"]
