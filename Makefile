ll: 
	make -B backend | make -B frontend

# flutter commands

frontend:
	cd frontend ; \
	flutter run -d web-server --web-hostname=0.0.0.0 --web-port=8000


# Golang commands

all:	
	cd backend ; \
	test vet fmt lint run

test:
	cd backend ; \
	go test ./...

vet:
	cd backend ; \
	go vet ./...

fmt:
	cd backend ; \
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l ; \
	test -z $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)

lint:
	cd backend ; \
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

build:
	cd backend ; \
	go build -o ./build/main main.go
    
backend:
	cd backend ; \
	go run main.go

# SQL database

SQL_User_name := postgres
SQL_Password := password
SQL_Host := 127.0.0.1
SQL_Port := 5432
SQL_Database := thefreepress
SQL_SSL_mode := disable
Migration_file_path := db/migration
PostgreSQL_address := postgres://$(SQL_User_name):$(SQL_Password)@$(SQL_Host):$(SQL_Port)/$(SQL_Database)?sslmode=$(SQL_SSL_mode)

migrate-up:	
	cd backend ; \
	migrate -path $(Migration_file_path) -database $(PostgreSQL_address) -verbose up

migrate-down:
	cd backend ; \
	migrate -path $(Migration_file_path) -database $(PostgreSQL_address) -verbose down

