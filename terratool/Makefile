###############################################################################
# Name        : Makefile                                                      #
# Author      : Abel Gancsos                                                  #
# Version     : v. 1.0.0.0                                                    #
# Description : Helps build the utility.      .                               #
###############################################################################
.DEFAULT_GOAL := build
IMAGE_NAME    := $(shell basename "$$PWD" )

build:
	if [ ! -d "./dist" ]; then mkdir dist; fi
	GO111MODULE=off GOPATH="$(realpath .)/src/classes" go build -o dist/$(IMAGE_NAME) src/main.go

clean:
	rm -fr dist/**

