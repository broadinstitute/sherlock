local-up:
	docker-compose -f dev/local-with-pg.yaml up --build

local-stop:
	docker-compose -f dev/local-with-pg.yaml stop

local-down:
	docker-compose -f dev/local-with-pg.yaml down --volumes

test:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5431:5432 postgres:13 -c max_connections=200
	cd go-shared && go test -p 1 -v -race ./...
	cd sherlock && go test -p 1 -v -race ./...
	docker stop test-postgres
	docker rm test-postgres

test-with-coverage:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5431:5432 postgres:13 -c max_connections=200
	cd go-shared && go test -p 1 -v -race -coverprofile=cover.out -covermode=atomic ./...
	cd sherlock && go test -p 1 -v -race -coverprofile=cover.out -covermode=atomic ./...
	docker stop test-postgres
	docker rm -f -v test-postgres

make pg-up:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5431:5432 postgres:13 -c max_connections=200

pg-down:
	docker stop test-postgres
	docker rm -f -v test-postgres

# To install swag, `go install github.com/swaggo/swag/cmd/swag@latest`
generate-swagger:
	cd sherlock && swag fmt -d ./ -g internal/boot/router.go && swag init -d ./ -g internal/boot/router.go

# To install mockery, `brew install mockery`
# 	(As a backup, use `go install github.com/vektra/mockery/v2@v2.32.4` but check https://github.com/vektra/mockery/releases for versions)
generate-mocks:
	cd sherlock && go generate ./...

# We use `go mod vendor` to package local dependencies in a way that the gcloud CLI can upload, but
# we don't keep that directory around because it'll confuse GoLand (we want our source of truth to be
# go.mod)
deploy-webhook-proxy:
	cd sherlock-webhook-proxy \
		&& go mod vendor \
		&& gcloud functions deploy sherlock-webhook-proxy \
			--gen2 \
			--project=dsp-tools-k8s \
			--region us-central1 \
			--runtime go119 \
			--trigger-http \
			--allow-unauthenticated \
			--source . \
			--entry-point HandleWebhook \
			--set-env-vars SHERLOCK_URL='https://sherlock.dsp-devops.broadinstitute.org',IAP_AUDIENCE="1038484894585-k8qvf7l876733laev0lm8kenfa2lj6bn.apps.googleusercontent.com",ALLOWED_GITHUB_ORGS="broadinstitute DataBiosphere CancerDataAggregator" \
			--set-secrets GITHUB_WEBHOOK_SECRET=sherlock-prod-webhook-secret:latest \
			--service-account sherlock-webhook-proxy@dsp-tools-k8s.iam.gserviceaccount.com \
		; rm -rf vendor
