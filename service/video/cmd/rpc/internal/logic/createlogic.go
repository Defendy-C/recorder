package logic

import (
	"context"
	"recorder/service/video/model"
	"recorder/util/validate"

	"recorder/service/video/cmd/rpc/internal/svc"
	"recorder/service/video/cmd/rpc/video"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *video.CreateReq) (*video.CreateResp, error) {
	if validate.ValuesHasZero(in.UserId, in.CreatedAt, in.VideoName, in.FileId) {
		return nil, model.ErrValidate
	}

	r, err := l.svcCtx.VideoModel.Insert(model.Video{
		Title: in.VideoName,
		UserId: in.UserId,
		FileId: in.FileId,
		CreatedAt: validate.StringToDate(in.CreatedAt),
		Description: in.Desc,
	})
	if err != nil {
		return nil, err
	}

	c, err := r.RowsAffected()
	if err != nil {
		return nil, err
	}

	if c != 1 {
		return &video.CreateResp{
			Ok: false,
		}, nil
	}

	id, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &video.CreateResp{
		Id: id,
		Ok: true,
	}, nil
}
