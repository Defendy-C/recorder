package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"recorder/service/video/cmd/rpc/internal/config"
	"recorder/service/video/model"
)

type ServiceContext struct {
	Config config.Config
	VideoModel model.RealVideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		VideoModel: model.NewRealVideoModel(sqlx.NewMysql(c.Mysql.Datasource), c.Cache),
	}
}
