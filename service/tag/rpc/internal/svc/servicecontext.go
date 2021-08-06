package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/stores/sqlx"

	"github.com/he2121/go-blog/service/tag/rpc/internal/config"
	"github.com/he2121/go-blog/service/tag/rpc/model"
)

type ServiceContext struct {
	Config   config.Config
	TagModel model.TagModel
	Redis    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		TagModel: model.NewTagModel(conn),
		Redis:    redis.NewRedis(c.CacheRedis[0].Host, c.CacheRedis[0].Type, c.CacheRedis[0].Pass),
	}
}
