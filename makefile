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
