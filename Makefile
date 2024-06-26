SHELL := /bin/bash

build_and_launch:
	go build && ./sport_planning_go_app
develop:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."
init_prod_db:
	rm -f database/database.db && sqlite3 database/database.db < database/init.sql
test_models:
	go test ./test
lint:
	golangci-lint run