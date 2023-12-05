package xredis

import (
	"lark/com/pkgs/xlog"
	"lark/pb"

	redsync "github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	redis "github.com/redis/go-redis/v9"
)

var cli *RedisClient

type RedisClient struct {
	client    *redis.Client
	RedisSync *redsync.Redsync
	Prefix    string
}

func NewRedisClient(cfg *pb.Jredis) *RedisClient {
	//先设置单机模式
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address[0],
		Password: cfg.Password,
		DB:       int(cfg.Db),
	})
	//redis 锁
	pool := goredis.NewPool(client)
	redsync := redsync.New(pool)

	cli = &RedisClient{
		client:    client,
		RedisSync: redsync,
		Prefix:    cfg.Prefix,
	}
	return cli
}

func GetRedisClient() *RedisClient {
	if cli == nil {
		xlog.Error("Redis client is not exist, create it first")
		return nil
	}
	return cli
}
