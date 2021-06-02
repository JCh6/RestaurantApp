export PORT=7080
export GRAPHQL=http://localhost:8080/graphql
export ALTER_DIR=http://localhost:8080/alter
export SET_SCHEMA_DIR=http://localhost:8080/admin/schema
export BASE_ENDPOINT=https://kqxty15mpg.execute-api.us-east-1.amazonaws.com
go test -v ./...
go run cmd/app/main.go