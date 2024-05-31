package conf

import (
	"github.com/BurntSushi/toml"
)

const (
	Test       string = "test"
	Production string = "production"
)

type tomlconfig struct {
	Env         string
	Eth_Network string
	MySQL       map[string]database
	Redis       map[string]redis
	Server      map[string]server
	Eth         map[string]ethereum `toml:"ethereum"`
}

type serverConf struct {
	MySQL  database
	Server server
	Redis  redis
	Eth    ethereum
}

type server struct {
	IP   string
	PORT int
}

type database struct {
	IP   string
	Port int
	User string
	Pass string
	DB   string
}

type redis struct {
	IsCluster bool
	Server    []string
	Pass      string
}

type ethereum struct {
	RPC_URL     string
	API_KEY     string
	BAYC_ADDR   string
	BLOCK_START uint64
	BLOCK_END   uint64
}

var Config *serverConf

func InitConfig(path *string) {
	tomlConf := &tomlconfig{}
	if _, err := toml.DecodeFile(*path, &tomlConf); err != nil {
		panic(err)
	}
	//log.Printf("tomlConf: %v", *tomlConf)
	Config = &serverConf{
		MySQL:  tomlConf.MySQL[tomlConf.Env],
		Server: tomlConf.Server[tomlConf.Env],
		Redis:  tomlConf.Redis[tomlConf.Env],
		Eth:    tomlConf.Eth[tomlConf.Eth_Network],
	}
	//log.Printf("Config: %v", *Config)
}
