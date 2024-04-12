migration-up:
	migrate -path ./migrations/postgres -database 'postgres://person:1234@0.0.0.0:5432/rentingcars?sslmode=disable' up


migration-down:
	migrate -path ./migrations/postgres -database 'postgres://person:1234@0.0.0.0:5432/rentingcars?sslmode=disable' down