package NFT

import (
	"Microservice/ent"
	pb "Microservice/proto/NFT"
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	BAYC "Microservice/Contract"
)

type Service struct {
	pb.UnimplementedNFTServiceServer
}

func Init() {
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := Mysql.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func (s *Service) AddCollection(ctx context.Context, in *pb.AddCollectionRequest) (*pb.AddCollectionResponse, error) {

	contractAddr := common.HexToAddress("0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D")
	contract, err := BAYC.NewContract(contractAddr, EthClient)
	if err != nil {
		log.Fatal("bind contract failed ", err)
	}
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(19957537),
		ToBlock:   big.NewInt(19961089),
		Addresses: []common.Address{
			contractAddr,
		},
	}
	logs, err := EthClient.FilterLogs(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	task1, err := Mysql.Transaction.Create().
		SetFrom("0x111").
		SetTo("0x222").Save(ctx)


	builders := make([]*ent.TransactionCreate, 0, len(logs))
	for i, vLog := range logs {
		event, err := contract.ParseTransfer(vLog)
		if err != nil {
			log.Fatal("Parse event error ", err)
			continue
		}
		fmt.Println("fetched event: ", event)
		dbrow :=Mysql.Transaction.Create().
		SetFrom(strings.ToLower(string(event.From.Bytes())))
		

		builders[i] = dbrow
	}

	if err != nil {
		log.Fatalf("failed creating a Transaction: %v", err)
	}

	fmt.Printf("%q => %q", task1.From, task1.To)

	return &pb.AddCollectionResponse{
		Status:  200,
		Message: "Collection added successfully",
	}, nil
}
