package resources

import (
	"log"
	"server/conf"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	BAYC_ADDR   string
	BLOCK_START uint64
	BLOCK_END   uint64
)

var EthClient *ethclient.Client

func InitGeth() {
	ethurl := conf.Config.Eth.RPC_URL + conf.Config.Eth.API_KEY
	client, err := ethclient.Dial(ethurl)
	if err != nil {
		panic(err)
	}
	log.Printf("InitGeth url:" + ethurl)
	BAYC_ADDR = conf.Config.Eth.BAYC_ADDR
	BLOCK_START = conf.Config.Eth.BLOCK_START
	BLOCK_END = conf.Config.Eth.BLOCK_END
	EthClient = client
}
