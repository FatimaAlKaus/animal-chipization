.PHONY: migrate-new
migrate-new:
	migrate create -ext sql -dir db/migrations -seq $$name  

.PHONY: migrate-up
migrate-up:
	migrate -database 'postgresql://user:password@localhost:5432/animal-chipization?sslmode=disable' -path db/migrations up

.PHONY: migrate-down
migrate-down:
	migrate -database 'postgresql://user:password@localhost:5432/animal-chipization?sslmode=disable' -path db/migrations down $n