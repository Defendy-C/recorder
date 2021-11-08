package main

import (
	"flag"
	"fmt"

	"recorder/service/file_sys/cmd/rpc/filesys"
	"recorder/service/file_sys/cmd/rpc/internal/config"
	"recorder/service/file_sys/cmd/rpc/internal/server"
	"recorder/service/file_sys/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/filesys.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewFileSysServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		filesys.RegisterFileSysServer(grpcServer, srv)

		switch c.Mode {
		case service.DevMode, service.TestMode:
			reflection.Register(grpcServer)
		default:
		}

	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
