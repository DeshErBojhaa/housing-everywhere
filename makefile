SHELL := /bin/bash

dns:
	docker build \
		-f dockerfile \
		-t dns:1.0 \
		.

up:
	docker-compose up

down:
	docker-compose down

test:
	go test ./... -v

clean:
	docker system prune -f

stop:
	docker stop $(docker ps -aq)

remove:
	docker rm $(docker ps -aq)