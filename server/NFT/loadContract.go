package NFT

import (
	pb "Microservice/proto/NFT"
	"Microservice/server/ent"
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	BAYC "Microservice/server/Contract"
)

type Service struct {
	pb.UnimplementedNFTServiceServer
}

func (s *Service) AddCollection(ctx context.Context, in *pb.AddCollectionRequest) (*pb.AddCollectionResponse, error) {
	_ctx := context.Background()
	contractAddr := common.HexToAddress("0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D")
	contract, err := BAYC.NewContract(contractAddr, EthClient)
	if err != nil {
		log.Fatal("bind contract failed ", err)
	}

	log.Print("qurey")
	var end uint64 = 19967402
	filter_opts := &bind.FilterOpts{
		Start:   19967401,
		End:     &end,
		Context: _ctx,
	}
	logs, err := contract.FilterTransfer(filter_opts, []common.Address{}, []common.Address{}, []*big.Int{})
	if err != nil {
		log.Print(err.Error())
		return &pb.AddCollectionResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	builders := []*ent.TransactionCreate{}
	for logs.Next() {
		event := logs.Event
		fmt.Println("fetched event: ", event)
		dbrow := Mysql.Transaction.Create().
			SetFrom(event.From).
			SetTo(event.To).
			SetTokenId(event.TokenId)

		builders = append(builders, dbrow)
	}
	log.Print("CreateBulk")
	if _, err := Mysql.Transaction.CreateBulk(builders...).Save(_ctx); err != nil {
		log.Printf("failed creating a Transaction: %v", err)
	}
	log.Print("CreateBulk finished")
	return &pb.AddCollectionResponse{
		Status:  http.StatusOK,
		Message: "Collection added successfully",
	}, nil
}
