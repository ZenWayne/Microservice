#!/bin/bash
cd ./server
go clean -testcache
go test ./test -run TestAddCollection -v
cd -