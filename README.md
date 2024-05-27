# Microservice
Microservice example powered by Kubernetes API Gateway

## tech stack

grpc-web => Envoy(gateway) => grpc server(gRPC ,gorm, go-ethereum)

Additional:
Istio Service mesh

## NFT Microservice

add contract address
switch contract

contract specific
transform get notices immediately even in pending
get token holder's NFT and transacition of it

NFT specific
get NFT transacition detailed info
get NFT json


## gRPC
1. install protocol compiler
pacman install protobuf
2. install plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
3. set up environment
export PATH="$PATH:$(go env GOPATH)/bin"




## go-ethereum


