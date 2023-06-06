local-up:
	docker-compose -f sherlock/dev/local-with-pg.yaml up --build

local-stop:
	docker-compose -f sherlock/dev/local-with-pg.yaml stop

local-down:
	docker-compose -f sherlock/dev/local-with-pg.yaml down --volumes

# We bump the max connections because I (Jack) was sloppy with database connections but not quantity of tests
functional-test:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5432:5432 postgres:13 -c max_connections=200
	export SHERLOCK_DB_PASSWORD="password" && go test -p 1 -v -race sherlock/...
	docker stop test-postgres
	docker rm test-postgres

unit-test:
	go test -v -short -race sherlock/...

tests-with-coverage:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5432:5432 postgres:13 -c max_connections=200
	export SHERLOCK_DB_PASSWORD="password" && go test -p 1 -v -race -coverprofile=sherlock/cover.out -covermode=atomic sherlock/...
	docker stop test-postgres
	docker rm -f -v test-postgres

make pg-up:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5432:5432 postgres:13 -c max_connections=200

pg-down:
	docker stop test-postgres
	docker rm -f -v test-postgres

# To install swag, `go install github.com/swaggo/swag/cmd/swag@latest`
generate-swagger:
	swag fmt -d sherlock -g sherlock/internal/boot/router.go
	swag init -d sherlock -g sherlock/internal/boot/router.go
