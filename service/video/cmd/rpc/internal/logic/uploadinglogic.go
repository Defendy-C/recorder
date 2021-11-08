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

type UploadingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadingLogic {
	return &UploadingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadingLogic) Uploading(in *video.UploadingReq) (*video.UploadingResp, error) {
	// validate
	if validate.ValuesHasZero(in.UserId, in.VideoName, in.Path) {
		return nil, model.ErrValidate
	}

	// insert data
	res, err := l.svcCtx.VideoModel.Insert(model.Video{
		Name: in.VideoName,
		UserId: in.UserId,
		Path: in.Path,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	// get latest id
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &video.UploadingResp{
		Id: id,
		Ok: true,
	}, nil
}
