package NFT

import (
	"server/conf"
	"server/ent"
	"server/resources"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
)

var (
	BAYC_ADDR   string
	BLOCK_START uint64
	BLOCK_END   uint64
	Mysql       *ent.Client
	Redis       redis.UniversalClient
	EthClient   *ethclient.Client
)

func Initvar() {
	BAYC_ADDR = conf.Config.Eth.BAYC_ADDR
	BLOCK_START = conf.Config.Eth.BLOCK_START
	BLOCK_END = conf.Config.Eth.BLOCK_END

	Mysql = resources.Mysql
	Redis = resources.Redis
	EthClient = resources.EthClient
}
