DOCKER?=docker
SERVICE := loyalty-system
GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	$(DOCKER) build --tag $(SERVICE) .

run: build
	$(DOCKER) run -d --publish 8080:8080 $(SERVICE) -d