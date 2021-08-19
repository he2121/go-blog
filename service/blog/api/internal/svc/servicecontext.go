package svc

import (
	"github.com/he2121/go-blog/service/blog/rpc/blog"
	"github.com/tal-tech/go-zero/zrpc"

	"github.com/he2121/go-blog/service/blog/api/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	BlogRpc blog.BlogServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		BlogRpc: blog.NewBlogServiceClient(zrpc.MustNewClient(c.BlogRpc)),
	}
}
