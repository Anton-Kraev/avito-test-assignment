.PHONY:
.SILENT:

app:
	docker-compose up --build -d app

apigen:
	oapi-codegen -generate types -o "openapi/api/types.gen.go" \
 				 -package "api" "openapi/schema/schema.yaml"
	oapi-codegen -generate server -o "openapi/api/api.gen.go" \
				 -package "api" "openapi/schema/schema.yaml"

postgres:
	docker-compose up --build -d postgres

migrate:
	./scripts/migration.sh

add-migration:
	goose -dir ./migrations -s create $(name) sql
