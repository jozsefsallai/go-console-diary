run: build
	./diary help

build: update
	go build -o diary

update:
	go get -u github.com/mitchellh/go-homedir

all: update build run