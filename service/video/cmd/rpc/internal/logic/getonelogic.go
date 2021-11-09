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
	if validate.ValuesHasZero(in.Id) {
		return nil, model.ErrValidate
	}

	v, err := l.svcCtx.VideoModel.FindOne(in.Id)
	if err != nil {
		return nil, err
	}

	return &video.GetOneResp{
		UserId:    v.UserId,
		Name:      v.Title,
		FileId:    v.FileId,
		CreatedAt: validate.DateToString(&v.CreatedAt),
		Desc:      v.Description,
	}, nil
}
