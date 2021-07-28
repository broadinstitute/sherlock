local-up:
	docker-compose -f build/local/server/docker-compose.yaml up -d --build

local-stop:
	docker-compose -f build/local/server/docker-compose.yaml stop

local-down:
	docker-compose -f build/local/server/docker-compose.yaml down --volumes
