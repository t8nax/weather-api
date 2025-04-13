run:
	swag init -g main.go
	go run main.go
.PHONY: run

run-docker:
	docker run -d -it -p 3001:8080 --env-file .env  weather-api