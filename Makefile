.PHONY: help up down restart

help:
	echo "you dont need help :-)"

up:
	docker-compose up --build -d

down:
	docker-compose down

restart:
	docker-compose build starcloud-api
	docker-compose up --no-deps -d starcloud-api
