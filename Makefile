BINARY_NAME=app

build:
	go build -o ${BINARY_NAME} main.go

build_all_platforms:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=window go build -o ${BINARY_NAME}-windows main.go

run_migrations:
	./${BINARY_NAME} -program=migrations

run_api:
	./${BINARY_NAME} -program=http-api

build_run_migrations: build run_migrations

build_run_api: build run_api

clean:
	go clean
	rm ${BINARY_NAME}
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all

goose_up:
	(cd migrations; goose postgres "user=postgres dbname=wiz_freight password=password sslmode=disable" up)

goose_down:
	(cd migrations; goose postgres "user=postgres dbname=wiz_freight password=password sslmode=disable" down)

proto_build:
	find ./proto/app/v1 -name "*.proto" | \
	xargs -I '{}' protoc --go_out=./proto/gen --go_opt=paths=source_relative \
	--go-grpc_out=./proto/gen --go-grpc_opt=paths=source_relative --proto_path=./proto  "{}"