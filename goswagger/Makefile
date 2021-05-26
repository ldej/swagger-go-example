.PHONY: swagger

run:
	go run .

install_deps:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
	go install github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	swagger generate spec -o ./swagger/swagger.json --scan-models

serve:
	swagger serve ./swagger/swagger.json