container_golang="ruskotwo/emotional-analyzer/golang:latest"
container_python="ruskotwo/emotional-analyzer/python:latest"

container_tests_telegram_bot="ruskotwo/emotional-analyzer/telegram-bot:latest"

generate:
	cd cmd/factory && wire ; cd ../..

build_golang:
	docker build -t ${container_golang} -f ./docker/golang.Dockerfile .

dev_golang: generate
	docker build -t ${container_golang} -f ./docker/golang.dev.Dockerfile .
	docker-compose up -d --profile app

build_python:
	docker build -t ${container_python} -f ./docker/python.Dockerfile .

all: generate build_golang build_python
	docker-compose --profile app up -d

tests_telegram_bot:
	docker build -t ${container_tests_telegram_bot} -f ./tests/telegram-bot/Dockerfile ./tests/telegram-bot
	docker-compose --profile tests up -d