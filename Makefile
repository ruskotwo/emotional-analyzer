container_golang="ruskotwo/emotional-analyzer/golang:latest"
container_python="ruskotwo/emotional-analyzer/python:latest"

generate:
	cd cmd/factory && wire ; cd ../..

build_golang:
	docker build -t ${container_golang} -f ./docker/golang.Dockerfile .

build_python:
	docker build -t ${container_python} -f ./docker/python.Dockerfile .

all: generate build_golang
	docker-compose up -d