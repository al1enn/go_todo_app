migrate create -ext sql -dir ./migrations -seq init

migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up

migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' force 1

migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' down


docker run --name=todo -e POSTGRES_PASSWORD='postgres' -p 5432:5432 -d --rm postgres   