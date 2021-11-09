package logic

import (
	"context"
	"recorder/service/video/model"
	"recorder/util/validate"

	"recorder/service/video/cmd/rpc/internal/svc"
	"recorder/service/video/cmd/rpc/video"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListLogic {
	return &GetVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoListLogic) GetVideoList(in *video.GetVideoListReq) (*video.GetVideoListResp, error) {
	// validate
	createdAt := validate.StringToDate(in.CreatedAt)
	if validate.ValuesHasZero(createdAt, in.UserId) {
		return nil, model.ErrValidate
	}

	opt := model.NewListOption(in.Page, in.PageSize)
	// get data
	vs, err := l.svcCtx.VideoModel.FindVideosByUserIdCreatedAt(in.UserId, createdAt, opt)
	if err != nil {
		return nil, err
	}

	// pack
	list := make([]*video.GetVideoListRespVideoItem, len(vs))
	for i, v := range vs {
		list[i] = &video.GetVideoListRespVideoItem{
			Id:        v.Id,
			Name:      v.Title,
			CreatedAt: validate.DateToString(&v.CreatedAt),
			Desc:      v.Description,
		}
	}

	return &video.GetVideoListResp{
		List: list,
		TotalPages: opt.TotalPages,
		TotalCount: opt.TotalCount,
	}, nil
}
