package main

import (
	"Microservice/NFT"
	MysqlClient "Microservice/mysql"
	pb "Microservice/proto/NFT"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "port to listen on")
)

func main() {

	//RedisClient.Init()
	MysqlClient.Init()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterNFTServiceServer(s, &NFT.Service{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
