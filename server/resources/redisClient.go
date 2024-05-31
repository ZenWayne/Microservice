package resources

import (
	"context"
	"server/conf"

	"github.com/redis/go-redis/v9"
)

var Redis redis.UniversalClient

func InitRedis() {
	if conf.Config.Redis.IsCluster {
		Redis = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    conf.Config.Redis.Server,
			Password: conf.Config.Redis.Pass,
		})
	} else {
		Redis = redis.NewClient(&redis.Options{
			Addr:     conf.Config.Redis.Server[0],
			Password: conf.Config.Redis.Pass,
		})
	}

	ctx := context.Background()
	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
