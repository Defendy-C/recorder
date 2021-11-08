package logic

import (
	"context"
	"recorder/service/video/model"
	"recorder/util/validate"

	"recorder/service/video/cmd/rpc/internal/svc"
	"recorder/service/video/cmd/rpc/video"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *video.DeleteReq) (*video.DeleteResp, error) {
	if validate.ValuesHasZero(in.Id) {
		return nil, model.ErrValidate
	}

	err := l.svcCtx.VideoModel.Delete(in.Id)
	if err != nil {
		return nil, err
	}

	return &video.DeleteResp{
		Ok: true,
	}, nil
}
