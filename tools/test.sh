#!/bin/bash
cd ./server
go clean -testcache
go test ./test -config=../config.toml -run TestAddCollection -v
cd -