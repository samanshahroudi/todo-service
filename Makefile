# Makefile

.PHONY: run test benchmark

run:
	docker compose up --build

test:
	go test ./... -v

benchmark:
	go test ./tests/ -bench=.
