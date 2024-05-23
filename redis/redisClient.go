package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.ClusterClient

func Init() {
	Rdb = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"172.22.3.132:6379", // Replace with the IP and port of your Redis cluster nodes
			"172.22.3.132:6380",
			"172.22.3.132:6381",
		},
		Password: "123",
	})

	ctx := context.Background()
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
