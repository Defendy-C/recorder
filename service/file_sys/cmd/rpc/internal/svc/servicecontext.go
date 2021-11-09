package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"recorder/service/file_sys/cmd/rpc/internal/config"
	"recorder/service/file_sys/model"
)

type ServiceContext struct {
	Config       config.Config
	FileSysModel model.FileSysModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		FileSysModel: model.NewFileSysModel(sqlx.NewMysql(c.Mysql.Datasource), c.Cache),
	}
}
