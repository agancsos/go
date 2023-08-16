###############################################################################
# Name        : Makefile                                                      #
# Author      : Abel Gancsos                                                  #
# Version     : v. 1.0.0.0                                                    #
# Description : Helps build the utility.                                      #
###############################################################################
.DEFAULT_GOAL := build
IMAGE_NAME    := $(shell basename "$$PWD" )

build:
	which go
	if [ ! -d "./dist" ]; then mkdir ./dist; fi 
	GOPATH=$(PWD)/plugins go build -o ./dist/plugintester ./src/main.go

clean:
	if [ -d "./dist" ]; then rm -fr ./dist/**; fi

container:
	which docker
	docker build --platform linux/x86_64 -t $(IMAGE_NAME):latest .

plugin:
	which go
	GOPATH=$(PWD)/plugins/ go build -buildmode=plugin -o ./dist/plugins/$(name).so ./plugins/$(name)/*.go
