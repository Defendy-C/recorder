package logic

import (
	"context"
	"recorder/service/gateway/cmd/api/internal/svc"
	"recorder/service/gateway/cmd/api/internal/types"
	"recorder/service/gateway/httperror"
	"recorder/service/user/cmd/rpc/userclient"
	"recorder/util/auth"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginResp, error) {
	iat := time.Now().Unix()
	// get user
	resp, err := l.svcCtx.UserClient.GetUser(l.ctx, &userclient.GetUserReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
	return nil, httperror.ErrUserNotExists
	}

	// gen token
	token, err := auth.NewJwtToken(map[string]interface{}{
		"id": resp.Id,
		"exp": iat + int64(l.svcCtx.Config.Auth.AccessExpired),
		"iat": iat,
	}, []byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		logx.Errorf("get token failed: %v\n", err)
		return nil, httperror.ErrLoginFailed
	}

	return &types.LoginResp{
		Token: token,
	}, nil
}
