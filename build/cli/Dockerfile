ARG GO_VERSION='1.19'
ARG ALPINE_VERSION='3.15'
ARG SHERLOCK_BUILD_VERSION='development'

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build
WORKDIR /build
ENV CGO_ENABLED=0
ENV GOBIN=/bin

# Install git because go build will embed VCS info with it
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -ldflags="-X 'main.BuildVersion=${SHERLOCK_BUILD_VERSION}'" -o /bin/sherlock ./cmd/cli/...

FROM alpine:${ALPINE_VERSION} as runtime
COPY --from=build /bin/sherlock /bin/sherlock
ENTRYPOINT [ "/bin/sherlock" ]
