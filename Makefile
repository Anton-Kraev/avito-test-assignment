.PHONY:
.SILENT:

app:
	go run ./cmd/app/main.go

apigen:
	oapi-codegen -generate types -o "openapi/api/types.gen.go" \
 				 -package "api" "openapi/schema/schema.yaml"
	oapi-codegen -generate server -o "openapi/api/api.gen.go" \
				 -package "api" "openapi/schema/schema.yaml"