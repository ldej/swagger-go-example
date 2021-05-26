.PHONY: swagger

run:
	go run .

install_deps:
	go get -u github.com/swaggo/swag/cmd/swag

swagger:
	swag init -g ./server.go -o ./swagger
