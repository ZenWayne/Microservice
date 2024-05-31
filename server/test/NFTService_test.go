package NFT

import (
	"context"
	"flag"
	"log"
	"math/big"
	BAYC "server/Contract"
	"server/NFT"
	"server/conf"
	pb "server/proto/NFT"
	"server/resources"

	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	configFile = flag.String("config", "./config.toml", "path to config file")
)

func Init() {
	flag.Parse()
	//fmt.Println("Init() :", dir)
	conf.InitConfig(configFile)
	resources.InitGeth()
	NFT.Initvar()
}

func TestAddCollection(t *testing.T) {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	defer conn.Close()
	c := pb.NewNFTServiceClient(conn)

	ctx := context.Background()

	r, err := c.AddCollection(ctx, &pb.AddCollectionRequest{
		Addr: "0x123",
	})

	if err != nil {
		log.Fatalf("could not AddCollection: %v", err)
	}

	log.Printf("AddCollection: %v", r)
}

func TestGetTransaction(t *testing.T) {
	Init()
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewNFTServiceClient(conn)
	ctx := context.Background()

	addr_str := "0xF17d0e18a82106A7F625Cc09394b2E8fA6931418"

	resp, err := client.GetTransaction(ctx, &pb.GetTransactionRequest{
		Addr: addr_str,
	})
	if err != nil {
		t.Fatalf("failed to get transaction: %v", err)
	}

	transactionsMap := make(map[string]([]*pb.Transaction))
	for _, tx := range resp.Transactions {
		transactionsMap[tx.Txhash] = append(transactionsMap[tx.Txhash], tx)
	}

	contractAddr := common.HexToAddress(NFT.BAYC_ADDR)
	contract, err := BAYC.NewContract(contractAddr, resources.EthClient)
	if err != nil {
		log.Fatal("bind contract failed ", err)
	}
	end := NFT.BLOCK_END
	filter_opts := &bind.FilterOpts{
		Start:   NFT.BLOCK_START,
		End:     &end,
		Context: ctx,
	}
	if NFT.BLOCK_END == 0 {
		filter_opts.End = nil
	}

	transferIter, err := contract.FilterTransfer(filter_opts,
		[]common.Address{},
		[]common.Address{},
		[]*big.Int{})

	if err != nil {
		log.Print(err.Error())
		t.Fatalf("failed to get transaction from go-ethereum: %v", err)
		return
	}
	t.Logf("hashes for loop")
	hashes := map[string]chan int{}
	tokenIdChan := make(chan *big.Int, 1)

	//check if tokenId is in order, each hash has its own go routine to read the token id from the channel
	for txhash := range transactionsMap {
		hashes[txhash] = make(chan int, 1)
		go func(_hash string, _hash_chan chan int) {
			i := 0
			for i < len(transactionsMap[_hash]) {
				<-_hash_chan
				tokenIdChan <- new(big.Int).SetBytes(transactionsMap[_hash][i].TokenId)
				i++
			}
		}(txhash, hashes[txhash])
	}

	i := 0

	for transferIter.Next() {
		event := transferIter.Event
		if i < len(resp.Transactions) {
			hashes[event.Raw.TxHash.Hex()] <- 0
			tokenId := <-tokenIdChan
			if event.TokenId.Cmp(tokenId) == 0 {
				t.Logf("block: %v tx:%v token %v pass",
					event.Raw.BlockNumber, event.Raw.TxHash, tokenId)
				i++
			} else {
				t.Errorf("Event from ethClient and gRPC response do not match at index %d", i)
			}

		}
	}
	if i != len(resp.Transactions) {
		t.Errorf("Too More events from gRPC response(%v) than ethClient(%v)", len(resp.Transactions), i)
	}
}
