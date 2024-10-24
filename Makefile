# Makefile para MV4: Primary Node y Diaboromon

build:
	sudo docker-compose build primarynode diaboromon

up:
	sudo docker-compose up -d primarynode diaboromon

stop:
	sudo docker-compose stop

down:
	sudo docker-compose down

logs:
	sudo docker-compose logs -f primarynode diaboromon

.PHONY: build up down logs

