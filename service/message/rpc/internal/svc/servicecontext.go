package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"

	"github.com/he2121/go-blog/service/message/rpc/internal/config"
	"github.com/he2121/go-blog/service/message/rpc/model"
)

type ServiceContext struct {
	Config            config.Config
	NotificationModel model.NotificationModel
	MessageModel      model.MessageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:            c,
		NotificationModel: model.NewNotificationModel(conn),
		MessageModel:      model.NewMessageModel(conn),
	}
}
