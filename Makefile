start:
	docker-compose up -d

run:
	go run cmd/main.go

.PHONY:  start, run 