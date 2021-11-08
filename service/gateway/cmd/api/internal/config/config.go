package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret  string
		AccessExpired int
	}
	UserRpc    zrpc.RpcClientConf
	VideoRpc   zrpc.RpcClientConf
	FileSysRpc zrpc.RpcClientConf
}
