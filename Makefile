# Makefile para MV3: Continente Server y Data Node 2

build:
	sudo docker-compose build continenteserver datanode2

up:
	sudo docker-compose up -d continenteserver datanode2

down:
	sudo docker-compose down

logs:
	sudo docker-compose logs -f continenteserver datanode2

.PHONY: build up down logs
