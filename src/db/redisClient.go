package NFT

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.ClusterClient

func InitRedis() {
	Redis = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"172.20.0.1:6379", // Replace with the IP and port of your Redis cluster nodes
			"172.20.0.2:6380",
			"172.20.0.3:6381",
		},
		Password: "123",
	})

	ctx := context.Background()
	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
