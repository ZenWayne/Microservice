package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	testCase "Microservice/test/NFT"
)

var (
	addr = flag.String("addr", "localhost:8080", "address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	defer conn.Close()

	ctx := context.Background()

	testCase.Test(ctx, conn)
}
