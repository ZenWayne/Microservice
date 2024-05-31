#/bin/bash

proto_cate=$1

if [ -d "./proto/$proto_cate" ]; then
  # Directory exists
  protoc --go_out="./server/"  \
         --go-grpc_out="./server/"  \
         "./proto/$1/$1.proto"
fi
