package logic

import (
	"context"
	"recorder/service/video/model"
	"recorder/util/validate"
	"time"

	"recorder/service/video/cmd/rpc/internal/svc"
	"recorder/service/video/cmd/rpc/video"

	"github.com/tal-tech/go-zero/core/logx"
)

type UploadedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadedLogic {
	return &UploadedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadedLogic) Uploaded(in *video.UploadedReq) (*video.UploadedResp, error) {
	if validate.ValuesHasZero(in.UserId, in.VideoName) {
		return nil, model.ErrValidate
	}

	v, err := l.svcCtx.VideoModel.FindOneByUserIdName(in.UserId, in.VideoName)
	if err != nil {
		return nil, err
	}

	if !v.FinishedAt.IsZero() {
		return &video.UploadedResp{
			Ok: false,
		}, nil
	}

	v.FinishedAt = time.Now()
	err = l.svcCtx.VideoModel.Update(*v)
	if err != nil {
		return nil, err
	}

	return &video.UploadedResp{
		Ok: true,
	}, nil
}
