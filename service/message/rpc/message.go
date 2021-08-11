package main

import (
	"flag"
	"fmt"

	"github.com/he2121/go-blog/service/message/rpc/internal/config"
	"github.com/he2121/go-blog/service/message/rpc/internal/server"
	"github.com/he2121/go-blog/service/message/rpc/internal/svc"
	"github.com/he2121/go-blog/service/message/rpc/message"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/message.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewMessageServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		message.RegisterMessageServiceServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
