DB_URL := "postgres://postgres:1234@localhost:5432/postdb?sslmode=disable"

CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen-proto.sh
  
run:
	go run cmd/main.go

migration_up:
	migrate -path migrations/ -database "postgresql://postgres:123@localhost:5432/attractiondb?sslmode=disable" -verbose up

migration_down:
	migrate -path migrations/ -database "postgresql://postgres:123@localhost:5432/attractiondb?sslmode=disable" -verbose down

migration_fix:
	migrate -path migrations/ -database "postgresql://postgres:123@localhost:5432/attractiondb?sslmode=disable" force VERSION
