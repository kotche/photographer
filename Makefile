.PHONY: docker-up rebuild

docker-up:
	docker-compose up --build -d

rebuild:
	docker-compose down
	docker-compose up --build -d

swag:
	swag init -g cmd/main.go