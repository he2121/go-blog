package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/stores/sqlx"

	"user/internal/config"
	"user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	Redis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
		Redis:     redis.NewRedis(c.CacheRedis[0].Host, c.CacheRedis[0].Type, c.CacheRedis[0].Pass),
	}
}
