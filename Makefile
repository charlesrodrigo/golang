CONTAINER=golang

run-infra:
	docker-compose -f .\docker-local\docker-compose.yaml up -d

run-ms-person:
	cd ms-person && go run main.go