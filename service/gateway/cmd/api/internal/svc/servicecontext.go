package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"recorder/service/file_sys/cmd/rpc/filesysclient"
	"recorder/service/gateway/cmd/api/internal/config"
	"recorder/service/user/cmd/rpc/userclient"
	"recorder/service/video/cmd/rpc/videoclient"
)

type ServiceContext struct {
	Config        config.Config
	UserClient    userclient.User
	VideoClient   videoclient.Video
	FileSysClient filesysclient.FileSys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		UserClient:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		VideoClient:   videoclient.NewVideo(zrpc.MustNewClient(c.VideoRpc)),
		FileSysClient: filesysclient.NewFileSys(zrpc.MustNewClient(c.FileSysRpc)),
		Config:        c,
	}
}
