local-up:
	docker-compose -f build/local/server/docker-compose.yaml up --build

local-stop:
	docker-compose -f build/local/server/docker-compose.yaml stop

local-down:
	docker-compose -f build/local/server/docker-compose.yaml down --volumes

# not sure if this is good make style but it works for now
functional-test:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5432:5432 postgres:13 -c max_connections=200
	export SHERLOCK_DB_PASSWORD="password" && go test -p 1 -v -race ./...
	docker stop test-postgres
	docker rm test-postgres

unit-test:
	go test -v -short -race ./...

tests-with-coverage:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5432:5432 postgres:13 -c max_connections=200
	export SHERLOCK_DB_PASSWORD="password" && go test -p 1 -v -race -coverpkg=./... -coverprofile=cover.out -covermode=atomic ./...
	docker stop test-postgres
	docker rm -f -v test-postgres

make pg-up:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5432:5432 postgres:13 -c max_connections=200

pg-down:
	docker stop test-postgres
	docker rm -f -v test-postgres
