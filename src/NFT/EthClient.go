package NFT

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var EthClient *ethclient.Client

func InitGeth() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/")
	if err != nil {
		log.Fatal(err)
	}
	EthClient = client
}
