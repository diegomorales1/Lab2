# Makefile para MV2: Continente Folder y Nodo Tai

build:
	sudo docker-compose build continentefolder nodotai

up:
	sudo docker-compose up continentefolder nodotai

tai:
	sudo docker-compose up -d continentefolder nodotai
	sudo docker attach nodotai

stop:
	sudo docker-compose stop

down:
	sudo docker-compose down

logs:
	sudo docker-compose logs -f continentefolder nodotai

.PHONY: build up attach_nodotai down logs


