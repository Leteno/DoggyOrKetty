#!/bin/bash

workspace=`pwd`
[[ ${GOPATH} =~ ${workspace} ]] || {
    export GOPATH=${workspace}:${GOPATH}
}
go run src/main.go