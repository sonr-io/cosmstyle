
TASK := $(shell task -t .taskfile.dist.yml)


.PHONY: all build generate release clean

all: build

generate:
	task: templ-gen

release:
	task: nebula:publish

clean:
	rm -rf ./build

build:
	go build -o ./build/nebula ./main.go
	cp ./build/nebula ../../www/nebula
	rm -rf ./build
