CONTAINER=golang

run-infra:
	docker-compose -f .\docker-local\docker-compose.yaml up -d

run-all-docker:
	docker-compose -f .\docker-local\docker-compose.yaml -f .\docker-local\docker-compose-app.yaml up -d

run-ms-person:
	cd ms-person && go run main.go