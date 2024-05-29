package NFT

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

var EthClient *ethclient.Client

func InitGeth() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/a84c7f7d0e974150997af5229c347ce1")
	if err != nil {
		panic(err)
	}
	EthClient = client
}
