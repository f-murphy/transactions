run:
	docker-compose up -d

stop:
	docker-compose down

tests:
	go test -v ./...