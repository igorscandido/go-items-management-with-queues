up:
	docker compose -f build/docker-compose.yaml up -d

down:
	docker compose -f build/docker-compose.yaml down

logs-api:
	docker logs -f items-management-api

logs-consumer:
	docker logs -f items-management-consumer