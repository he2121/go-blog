package svc

import (
	"github.com/he2121/go-blog/service/blog/rpc/blogservice"
	"github.com/tal-tech/go-zero/zrpc"

	"github.com/he2121/go-blog/service/blog/api/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	BlogRpc blogservice.BlogService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		BlogRpc: blogservice.NewBlogService(zrpc.MustNewClient(c.BlogRpc)),
	}
}
