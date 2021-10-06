local-up:
	docker-compose -f build/local/server/docker-compose.yaml up --build

local-stop:
	docker-compose -f build/local/server/docker-compose.yaml stop

local-down:
	docker-compose -f build/local/server/docker-compose.yaml down --volumes

# not sure if this is good make style but it works for now
integration-test:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5432:5432 postgres:13
	export SHERLOCK_DBHOST="localhost" && export SHERLOCK_DBPASSWORD="password" && go test -v -race ./...
	docker stop test-postgres
	docker rm test-postgres

unit-test:
	go test -v -short -race ./...

tests-with-coverage:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5432:5432 postgres:13
	export SHERLOCK_DBHOST="localhost" && export SHERLOCK_DBPASSWORD="password" && go test -v -race -coverprofile=cover.out -covermode=atomic ./...
	docker stop test-postgres
	docker rm test-postgres

make pg-up:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5432:5432 postgres:13

pg-down:
	docker stop test-postgres
	docker rm test-postgres
