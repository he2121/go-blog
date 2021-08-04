package main

import (
	"flag"
	"fmt"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"

	"github.com/he2121/go-blog/service/blog/rpc/blog"
	"github.com/he2121/go-blog/service/blog/rpc/internal/config"
	"github.com/he2121/go-blog/service/blog/rpc/internal/server"
	"github.com/he2121/go-blog/service/blog/rpc/internal/svc"
)

var configFile = flag.String("f", "etc/blog.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewBlogServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		blog.RegisterBlogServiceServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
