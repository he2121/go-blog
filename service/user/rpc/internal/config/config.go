package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.ClusterConf
	// jwt-token 的签名密钥与失效时间
	LoginAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
