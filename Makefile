.PHONY: build run

BUILD_DIR = build
BIN_FILE = main
MAIN_FILE = cmd/main.go

build:
	CGO_ENABLED=1 GOOS=linux go build -o $(BUILD_DIR)/$(BIN_FILE)  $(MAIN_FILE)

run:
	$(BUILD_DIR)/$(BIN_FILE)
	
post:
	curl -H 'Content-Type: application/json' \
	-d '{"name": "$(d)", "color": "vermelho"}' \
	-X POST \
	-v \
	localhost:8080/category 

put:
	curl -H 'Content-Type: application/json' \
	-d '{"name":"alterada", "color": "vermelho"}' \
	-X PUT \
	-v \
	localhost:8080/category/1 

delete:
	curl -H 'Content-Type: application/json' \
	-X DELETE \
	-v \
	localhost:8080/category/1 

get:
	curl -H 'Content-Type: application/json' \
	-X GET \
	-v \
	localhost:8080/category/6

getall:
	curl \
	-X GET \
	-v \
	localhost:8080/category
