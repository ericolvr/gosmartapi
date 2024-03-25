start:
	docker-compose up -d

run:
	go run cmd/main.go

build:
	GOOS=darwin GOARCH=amd64 go build -o build/app ./cmd/main.go

run-builded:
	./build/app

.PHONY:  start, run, build