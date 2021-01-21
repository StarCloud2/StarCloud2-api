all:
	echo "Start Doker Compose"
	go fmt ./...
	docker-compose up --build