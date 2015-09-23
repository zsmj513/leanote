.PHONY: default install prepare

default: install

prepare:
	go get ./...

install:
	go install
