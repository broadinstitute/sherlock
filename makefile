local-up:
	docker-compose -f build/local/server/docker-compose.yaml up --build

local-stop:
	docker-compose -f build/local/server/docker-compose.yaml stop

local-down:
	docker-compose -f build/local/server/docker-compose.yaml down --volumes

integration-test:
	docker-compose -f build/local/server/docker-compose.test.yaml up --build --abort-on-container-exit
	docker-compose -f build/local/server/docker-compose.test.yaml down --volumes
	
integration-down:
	docker-compose -f build/local/server/docker-compose.test.yaml down --volumes

integration-local:
	docker run --name test-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=sherlock -d -p 5432:5432 postgres:13
	go test -v -race ./...
	docker stop test-postgres
	docker rm test-postgres

unit-test:
	go test -v -short -race ./...

tests-with-coverage:
	docker-compose -f build/local/server/docker-compose.test.yaml up --build --abort-on-container-exit
	docker-compose -f build/local/server/docker-compose.test.yaml down --volumes
