#/bin/bash
cd "$(dirname "$0")"

proto_cate=$1

if [ -d "$proto_cate" ]; then
  # Directory exists
  protoc --go_out="./" --go_opt=paths=source_relative \
         --go-grpc_out="./" --go-grpc_opt=paths=source_relative \
         "$1/$1.proto"
fi
