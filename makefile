local-up:
	docker-compose -f dev/local-with-pg.yaml up --build

local-stop:
	docker-compose -f dev/local-with-pg.yaml stop

local-down:
	docker-compose -f dev/local-with-pg.yaml down --volumes

# We bump the max connections because I (Jack) was sloppy with database connections but not quantity of tests
functional-test:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5432:5432 postgres:13 -c max_connections=200
	cd sherlock && export SHERLOCK_DB_PASSWORD="password" && go test -p 1 -v -race ./...
	docker stop test-postgres
	docker rm test-postgres

unit-test:
	cd sherlock && go test -v -short -race ./...

tests-with-coverage:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5432:5432 postgres:13 -c max_connections=200
	cd sherlock && export SHERLOCK_DB_PASSWORD="password" && go test -p 1 -v -race -coverprofile=cover.out -covermode=atomic ./...
	docker stop test-postgres
	docker rm -f -v test-postgres

make pg-up:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5432:5432 postgres:13 -c max_connections=200

pg-down:
	docker stop test-postgres
	docker rm -f -v test-postgres

# To install swag, `go install github.com/swaggo/swag/cmd/swag@latest`
generate-swagger:
	cd sherlock && swag fmt -d ./ -g internal/boot/router.go && swag init -d ./ -g internal/boot/router.go

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
			--set-env-vars SHERLOCK_URL='https://sherlock.dsp-devops.broadinstitute.org',IAP_AUDIENCE="1038484894585-k8qvf7l876733laev0lm8kenfa2lj6bn.apps.googleusercontent.com" \
			--set-secrets GITHUB_WEBHOOK_SECRET=sherlock-prod-webhook-secret:latest \
			--service-account sherlock-webhook-proxy@dsp-tools-k8s.iam.gserviceaccount.com \
		; rm -rf vendor
