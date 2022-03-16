#!/bin/bash
###############################################################################
# Name        : tests.sh                                                      #
# Author      : Abel Gancsos                                                  #
# Version     : v. 1.0.0.0                                                    #
# Description : Runs Unit Tests for project.                                  #
###############################################################################
BASE_PATH=$(dirname $0)/..
(
	cd $BASE_PATH && GO111MODULE=off GOPATH=$HOME go test ./src/...
)
printf "\e[32mDone!\e[m\n"

