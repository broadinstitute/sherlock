ARG GO_VERSION='1.19'
ARG ALPINE_VERSION='3.16'

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build
ARG BUILD_VERSION='development'
WORKDIR /build
ENV CGO_ENABLED=0
ENV GOBIN=/bin

COPY go.mod go.sum ./
RUN go mod download && go mod verify

# See the .dockerignore, it ignores by default
COPY . ./
RUN go build -buildvcs=false -ldflags="-X 'main.BuildVersion=${BUILD_VERSION}'" -o /bin/sherlock ./cmd/sherlock/...

# FROM alpine:${ALPINE_VERSION} as runtime <-- use this if you hit issues
FROM gcr.io/distroless/static:nonroot as runtime
COPY --from=build /bin/sherlock /bin/sherlock
ENTRYPOINT [ "/bin/sherlock" ]
