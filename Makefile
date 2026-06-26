run:
	go run cmd/main.go

build:
	go build -o bin/app cmd/main.go

test:
	go test ./...

fmt:
	go fmt ./...

clean:
	rm -rf bin