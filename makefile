SERVERENDPOINT=cmd/server/main.go

run:
	go run $(SERVERENDPOINT)
build:
	go build $(SERVERENDPOINT)

docker:
	docker build --tag telehook-anonymous .

.PHONY: run