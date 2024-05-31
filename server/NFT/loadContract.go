package NFT

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"server/ent"
	"server/ent/transaction"

	pb "server/proto/NFT"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	BAYC "server/Contract"
)

type Service struct {
	pb.UnimplementedNFTServiceServer
}

func (s *Service) AddCollection(ctx context.Context, in *pb.AddCollectionRequest) (*pb.AddCollectionResponse, error) {
	contractAddr := common.HexToAddress(BAYC_ADDR)
	contract, err := BAYC.NewContract(contractAddr, EthClient)
	if err != nil {
		log.Fatal("bind contract failed ", err)
	}

	end := BLOCK_END
	filter_opts := &bind.FilterOpts{
		Start:   BLOCK_START,
		End:     &end,
		Context: ctx,
	}
	if BLOCK_END == 0 {
		filter_opts.End = nil
	}

	transferIter, err := contract.FilterTransfer(filter_opts, []common.Address{}, []common.Address{}, []*big.Int{})
	if err != nil {
		log.Print(err.Error())
		return &pb.AddCollectionResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	builders := []*ent.TransactionCreate{}
	for transferIter.Next() {
		event := transferIter.Event
		fmt.Println("fetched event: ", event)
		dbrow := Mysql.Transaction.Create().
			SetBlockNumber(event.Raw.BlockNumber).
			SetTxHash(event.Raw.TxHash).
			SetFrom(event.From).
			SetTo(event.To).
			SetTokenId(event.TokenId)

		builders = append(builders, dbrow)
	}
	log.Print("CreateBulk")
	if _, err := Mysql.Transaction.CreateBulk(builders...).Save(ctx); err != nil {
		log.Printf("failed creating a Transaction: %v", err)
	}
	log.Print("CreateBulk finished")
	return &pb.AddCollectionResponse{
		Status:  http.StatusOK,
		Message: "Collection added successfully",
	}, nil
}

func (s *Service) GetTransaction(ctx context.Context, in *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {

	if !common.IsHexAddress(in.Addr) {
		return nil, fmt.Errorf("address Not Valid received addr:%v", in.Addr)
	}

	addr := common.HexToAddress(in.Addr)
	txs, err := Mysql.Transaction.Query().
		Where(
			transaction.Or(
				transaction.FromEQ(addr),
				transaction.ToEQ(addr),
			),
		).
		Order(
			transaction.ByID(),
		).
		All(ctx)

	if err != nil {
		log.Printf("failed querying a Transaction: %v", err)
		return nil, fmt.Errorf("failed querying a Transaction")
	}

	resp := &pb.GetTransactionResponse{}
	for _, tx := range txs {
		resp.Status = http.StatusOK
		transaction := &pb.Transaction{
			Blocknumber: tx.BlockNumber,
			Txhash:      tx.TxHash.Hex(),
			From:        tx.From.Hex(),
			To:          tx.To.Hex(),
			TokenId:     make([]byte, 32),
		}
		tx.TokenId.FillBytes([]byte(transaction.TokenId))
		resp.Transactions = append(resp.Transactions, transaction)
	}

	return resp, nil
}
