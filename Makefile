.PHONY: build
build:
	go build -v ./cmd/site-checker-tg-bot

run:
	SERVICE_CONFIG_PATH="./config/service.yaml" ./site-checker-tg-bot

.DEFAULT_GOAL := build