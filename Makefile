.PHONY: dependency listen test docker-up docker-down clear run

build:
	go build -o bin/tada-runner-generator

install:
	go install