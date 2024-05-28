package NFT

import (
	pb "Microservice/proto/NFT"
	"context"
	"log"

	"google.golang.org/grpc"
)

func Test(ctx context.Context, conn *grpc.ClientConn) {
	c := pb.NewNFTServiceClient(conn)

	r, err := c.AddCollection(ctx, &pb.AddCollectionRequest{
		Addr: "0x123",
	})

	if err != nil {
		log.Fatalf("could not AddCollection: %v", err)
	}

	log.Printf("AddCollection: %v", r)
}

