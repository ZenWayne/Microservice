package main

import (
	"flag"
	"server/conf"
)

var (
	configFile = flag.String("config", "config.toml", "path to config file")
)

func init() {
	//beware of the execution order
	flag.Parse()
	conf.InitConfig(configFile)
}
