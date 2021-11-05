package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"recorder/service/user/cmd/rpc/internal/config"
	"recorder/service/user/model"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUsersModel(sqlx.NewMysql(c.Mysql.Datasource), c.Cache),
	}
}
