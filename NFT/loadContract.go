package NFT

import (
	MysqlClient "Microservice/mysql"
	pb "Microservice/proto/NFT"
	"context"

	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	From    string `gorm:"type:CHAR(42);"`
	To      string `gorm:"type:CHAR(42);"`
	TokenId uint64 `gorm:"type:BIGINT(20);"`
}

type Service struct {
	pb.UnimplementedNFTServiceServer
}

func Init() {
	MysqlClient.DB.AutoMigrate(&Transfer{})
}

func (s *Service) AddCollection(ctx context.Context, in *pb.AddCollectionRequest) (*pb.AddCollectionResponse, error) {

	return &pb.AddCollectionResponse{
		Status:  200,
		Message: "Collection added successfully",
	}, nil
}
