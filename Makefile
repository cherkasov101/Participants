FILES = ./cmd/web

all: clean build
	./main

build:
	go build -o main ./cmd/web

clean:
	rm -f main