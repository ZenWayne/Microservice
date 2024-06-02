package main

import (
	"fmt"
	"log"
	"net"
	"server/NFT"
	"server/conf"
	pb "server/proto/NFT"
	"server/resources"

	"google.golang.org/grpc"
)

func main() {

	//RedisClient.Init()
	resources.InitMysql()
	resources.InitRedis()
	resources.InitGeth()
	NFT.Initvar()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.Config.Server.IP, conf.Config.Server.PORT))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterNFTServiceServer(s, &NFT.Service{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
