package logic

import (
	"context"
	"recorder/service/file_sys/cmd/rpc/filesys"
	"recorder/service/gateway/httperror"
	"recorder/service/video/cmd/rpc/video"
	"recorder/util/validate"
	"time"

	"recorder/service/gateway/cmd/api/internal/svc"
	"recorder/service/gateway/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UploadConnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadConnLogic(ctx context.Context, svcCtx *svc.ServiceContext) UploadConnLogic {
	return UploadConnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadConnLogic) UploadConn(req types.UploadConnReq) (*types.UploadConnResp, error) {
	if validate.ValuesHasZero(req.Chunks, req.UserId, req.VideoName) {
		return nil, httperror.ErrVideoUploadFailed
	}

	createdAt := time.Now()
	createdAtStr := validate.DateToString(&createdAt)
	fresp, err := l.svcCtx.FileSysClient.Create(l.ctx, &filesys.CreateReq{
		CreatedAt: createdAtStr,
		TotalChunks: int64(req.Chunks),
		UserId: int64(req.UserId),

	})
	if err != nil {
		return nil, httperror.ErrVideoUploadFailed
	}

	vresp, err := l.svcCtx.VideoClient.Create(l.ctx, &video.CreateReq{
		UserId:    int64(req.UserId),
		FileId:    fresp.Id,
		VideoName: req.VideoName,
		Desc:      req.Desc,
		CreatedAt: createdAtStr,
	})
	if err != nil {
		return nil, httperror.ErrVideoUploadFailed
	}

	return &types.UploadConnResp{
		Ok: vresp.Ok,
		Id: int(vresp.Id),
	}, nil
}
