APP_NAME=mileage-calculator
GO_FILES = $(shell go list ./...)


.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	go build -o bin/$(APP_NAME)

.PHONY: cover
cover:
	go test -race $(GO_FILES) -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func coverage.out