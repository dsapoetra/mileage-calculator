APP_NAME=go-mileage-calculator

build:
	go build -o bin/$(APP_NAME)

run.build :
	sudo chmod -R 777 bin/* && ./bin/$(APP_NAME)