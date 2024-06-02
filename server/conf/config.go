package conf

import (
	"log"

	"github.com/BurntSushi/toml"
)

type ServerConf struct {
	Server Server
	MySQL  Database
	Redis  Redis
	Eth    Ethereum
}

type Server struct {
	IP   string
	PORT int
}

type Database struct {
	IP   string
	Port int
	User string
	Pass string
	DB   string
}

type Redis struct {
	IsCluster bool
	Server    []string
	Pass      string
}

type Ethereum struct {
	RPC_URL     string
	API_KEY     string
	BAYC_ADDR   string
	BLOCK_START uint64
	BLOCK_END   uint64
}

const (
	Test       string = "test"
	Production string = "production"
)

type tomlconfig struct {
	Env         string
	Eth_Network string
	Server      map[string]Server
	MySQL       map[string]Database
	Redis       map[string]Redis
	Eth         map[string]Ethereum `toml:"ethereum"`
}

var Config *ServerConf

func InitConfig(path *string) {
	tomlConf := &tomlconfig{}
	if _, err := toml.DecodeFile(*path, &tomlConf); err != nil {
		panic(err)
	}
	//log.Printf("tomlConf: %v", *tomlConf)
	Config = &ServerConf{
		MySQL:  tomlConf.MySQL[tomlConf.Env],
		Server: tomlConf.Server[tomlConf.Env],
		Redis:  tomlConf.Redis[tomlConf.Env],
		Eth:    tomlConf.Eth[tomlConf.Eth_Network],
	}
	log.Printf("Config: %v", *Config)
}
