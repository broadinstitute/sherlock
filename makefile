BUILD_VERSION ?= development

## Local development

local-up:
	docker-compose -f dev/local-with-pg.yaml up --build

local-stop:
	docker-compose -f dev/local-with-pg.yaml stop

local-down:
	docker-compose -f dev/local-with-pg.yaml down --volumes

## Normal testing (depends on install-pact)

test:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5431:5432 postgres:15 -c max_connections=200
	cd go-shared && go test -p 1 -race ./...
	cd sherlock && go test -p 1 -race ./...
	docker stop test-postgres
	docker rm test-postgres

test-with-coverage:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5431:5432 postgres:15 -c max_connections=200
	cd go-shared && go test -p 1 -race -coverprofile=cover.out -covermode=atomic ./...
	cd sherlock && go test -p 1 -race -coverprofile=cover.out -covermode=atomic ./...
	docker stop test-postgres
	docker rm -f -v test-postgres

## Pact testing

install-pact:
	cd sherlock && go install $$(grep 'github.com/pact-foundation/pact-go/v2' go.mod | sed -ne 's/^[[:blank:]]*//p' | sed -ne 's/[[:blank:]]/@/p')
	sudo -s $$(which pact-go) -l DEBUG install -f

test-pact-provider:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5431:5432 postgres:15 -c max_connections=200
	cd sherlock && go test -p 1 -v -race -ldflags="-X 'github.com/broadinstitute/sherlock/go-shared/pkg/version.BuildVersion=${BUILD_VERSION}'" ./internal/pact/...

document-pact-provider:
	echo "# Sherlock Provider State Handlers" > sherlock/internal/pact/README.md
	echo "### [More Information](https://docs.pact.io/getting_started/provider_states)" >> sherlock/internal/pact/README.md
	sed -nr 's/.*stateHandlers\["([^"]+)"\].*/- `\1`/p' sherlock/internal/pact/provider_test.go >> sherlock/internal/pact/README.md
	sed -nr 's/^\t*\s*([a-zA-Z_0-9]+)\(\).+/- `\1 exists`/p' sherlock/internal/models/test_data.go >> sherlock/internal/pact/README.md

## Long running test database (so tests can be run from GoLand)

pg-up:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5431:5432 postgres:15 -c max_connections=200

pg-down:
	docker stop test-postgres
	docker rm -f -v test-postgres

## Swagger generation

install-swagger:
	cd sherlock && go install $$(grep 'github.com/swaggo/swag' go.mod | sed -ne 's/^[[:blank:]]*//p' | sed -ne 's/.*[[:blank:]]/github.com\/swaggo\/swag\/cmd\/swag@/p')

format-swagger:
	cd sherlock && swag fmt -d ./ -g internal/boot/router.go

generate-swagger:
	cd sherlock && swag init -d ./ -g internal/boot/router.go

## Mock generation

install-mockery:
# 	(As a backup, use `go install github.com/vektra/mockery/v2@v2.32.4` but check https://github.com/vektra/mockery/releases for versions)
	brew install mockery

generate-mocks:
	cd sherlock && mockery

## Cloud function deployment

# We use `go mod vendor` to package local dependencies in a way that the gcloud CLI can upload, but
# we don't keep that directory around because it'll confuse GoLand (we want our source of truth to be
# go.mod)
deploy-webhook-proxy:
	cd sherlock-webhook-proxy \
		&& go mod tidy \
		&& go mod vendor \
		&& gcloud functions deploy sherlock-webhook-proxy \
			--gen2 \
			--project=dsp-devops-super-prod \
			--region us-central1 \
			--runtime go121 \
			--trigger-http \
			--allow-unauthenticated \
			--source . \
			--entry-point HandleWebhook \
			--set-env-vars SHERLOCK_URL='https://sherlock.dsp-devops-prod.broadinstitute.org',IAP_AUDIENCE='257801540345-1gqi6qi66bjbssbv01horu9243el2r8b.apps.googleusercontent.com',ALLOWED_GITHUB_ORGS='broadinstitute DataBiosphere CancerDataAggregator hcholab' \
			--set-secrets GITHUB_WEBHOOK_SECRET=sherlock-github-webhook-secret:latest \
			--service-account sherlock-webhook-proxy@dsp-devops-super-prod.iam.gserviceaccount.com \
		; rm -rf vendor
