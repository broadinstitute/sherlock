ARG GO_VERSION='1.16'
ARG ALPINE_VERSION='3.14'

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as build
WORKDIR /build
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOBIN=/bin
COPY . .
RUN go test -short ./... && go build -o /bin/sherlock ./cmd/server/...

FROM alpine:${ALPINE_VERSION} as runtime
ENV APP_NAME=sherlock
COPY --from=build /bin/${APP_NAME} /bin/${APP_NAME}
# Copy Database migration changelogs into final image
COPY --from=build /build/db/migrations/ /db/migrations
# The sleep is a quick fix to ensure the postgres container has time
# to start when running locally via docker-compose
ENTRYPOINT [ "sh", "-c", "sleep 10 && ./bin/${APP_NAME}" ]
