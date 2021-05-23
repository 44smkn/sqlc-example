.PHONY: build
build:
	sqlc generate
	go build $(shell pwd)/cmd/toy-isuumo