package logic

import (
	"context"
	"recorder/service/video/model"
	"recorder/util/validate"

	"recorder/service/video/cmd/rpc/internal/svc"
	"recorder/service/video/cmd/rpc/video"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneLogic {
	return &GetOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOneLogic) GetOne(in *video.GetOneReq) (*video.GetOneResp, error) {
	if validate.ValuesHasZero(in.UserId, in.Name) {
		return nil, model.ErrValidate
	}

	v, err := l.svcCtx.VideoModel.FindOneByUserIdName(in.UserId, in.Name)
	if err != nil {
		return nil, err
	}

	finishedAt := ""
	if !v.FinishedAt.IsZero() {
		finishedAt = v.FinishedAt.Format("2006-01-02 15:04:05")
	}

	return &video.GetOneResp{
		Path:       v.Path,
		CreatedAt:  v.CreatedAt.Format("2006-01-02 15:04:05"),
		FinishedAt: finishedAt,
	}, nil
}
