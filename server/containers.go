//go:build containers
// +build containers

package main

import (
	"os"
	"server/conf"
	"strconv"
	"strings"
)

func init() {
	redis_servers := strings.Split(os.Getenv("REDIS_SERVERS"), ",")
	var isCluster bool
	if len(redis_servers) > 1 {
		isCluster = true
	} else {
		isCluster = false
	}
	conf.Config = &conf.ServerConf{
		MySQL: conf.Database{
			IP: os.Getenv("MYSQL_URL"),
			Port: func() int {
				port, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
				return port
			}(),
			User: os.Getenv("MYSQL_USER"),
			Pass: os.Getenv("MYSQL_PASSWORD"),
			DB:   os.Getenv("MYSQL_DATABASE"),
		},

		Redis: conf.Redis{
			IsCluster: isCluster,
			Server:    redis_servers,
			Pass:      os.Getenv("REDIS_PASS"),
		},
		Server: conf.Config.Server,
		Eth:    conf.Config.Eth,
	}
	conf.Config.Eth.API_KEY = os.Getenv("API_KEY")
}
