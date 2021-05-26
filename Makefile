SHELL := /bin/sh

# Variables definitions
# -----------------------------------------------------------------------------
PROJECT_NAME := hasty-challenge-manager
IMAGE_NAME := arthurhenrique/${PROJECT_NAME}
GOPACKAGES = $(shell go list ./...)
VERSION = "v0.1.0"

# Target section
# -----------------------------------------------------------------------------
.PHONY: clean run install build test

vendor:
	go mod vendor

install: clean vendor

run:
	go run main.go api

run-schedule:
	go run main.go schedule-checker

test:
	ENVIRONMENT="test" go test -count=1 -v $(GOPACKAGES)

test-deployment-yaml:
	wget https://github.com/instrumenta/kubeval/releases/latest/download/kubeval-linux-amd64.tar.gz && \
	tar xf kubeval-linux-amd64.tar.gz && \
	sudo cp kubeval /usr/local/bin && \
	kubeval deploy/*.yaml

docker/registry: docker/build docker/tag docker/push

docker/build:
	docker build --network=host -t $(IMAGE_NAME) .

docker/tag:
	docker tag $(IMAGE_NAME) $(IMAGE_NAME):$(VERSION)

docker/push:
	docker push $(IMAGE_NAME):$(VERSION) && \
	docker push $(IMAGE_NAME):latest

docker/up:
	docker-compose up -d
	sleep 5 # wait db be ready
	docker run --rm --network=host -v "$(PWD)/migrations:/flyway/sql:ro"  \
		boxfuse/flyway:5.2.4-alpine \
		-driver="org.postgresql.Driver" \
		-user="master" \
		-schemas="job_manager" \
		-password="123456" \
		-url="jdbc:postgresql://localhost:5432/job_manager" \
		migrate

docker/down:
	docker-compose down

clean:
	rm -rf vendor/
	rm -rf challenge/__pycache__
