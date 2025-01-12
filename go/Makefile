.PHONY: run build clean

build:
	go build -o bin/cgol -ldflags="-s -w" main.go

run:
	go run main.go --file ../input.txt --amount 10

clean:
	rm -rf bin