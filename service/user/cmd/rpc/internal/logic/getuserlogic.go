package logic

import (
	"context"
	"recorder/service/user/cmd/rpc/internal/svc"
	"recorder/service/user/cmd/rpc/user"
	"recorder/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserReply, error) {
	// validate
	if in.Username == "" || in.Password == "" {
		return nil, model.ErrValidate
	}

	// get user
	u, err := l.svcCtx.UserModel.FindOneByUsername(in.Username)
	if err != nil {
		logx.Infof("user %s not found: %v\n", in.Username, err)
		return nil, err
	}

	return &user.GetUserReply{
		Id: u.Id,
	}, nil
}