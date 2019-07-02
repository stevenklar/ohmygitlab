build:
	go build -o bin/ohmygitlab

run:
	go run main.go

clean:
	go mod tidy && go get

install:
	go install
