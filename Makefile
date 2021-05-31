APP=main

build: clean
	go build -o ./build/${APP} main.go

run:
	go run main.go