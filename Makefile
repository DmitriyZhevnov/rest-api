drun:
	docker-compose up -d
run:
	docker-compose up 
brun:
	docker-compose up --build
migrate:
	migrate -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable -path db/migrations up