package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"

	"github.com/he2121/go-blog/service/blog/rpc/internal/config"
	"github.com/he2121/go-blog/service/blog/rpc/model"
)

type ServiceContext struct {
	Config       config.Config
	BlogModel    model.BlogModel
	CommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		BlogModel:    model.NewBlogModel(conn),
		CommentModel: model.NewCommentModel(conn),
	}
}
