SERVERENDPOINT=cmd/server/main.go

run:
	go run $(SERVERENDPOINT)
build:
	go build $(SERVERENDPOINT)

.PHONY: run