package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"recorder/service/gateway/cmd/api/internal/config"
	"recorder/service/user/cmd/rpc/userclient"
)

type ServiceContext struct {
	Config     config.Config
	UserClient userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		UserClient: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Config:     c,
	}
}
