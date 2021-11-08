package logic

import (
	"context"
	"recorder/service/file_sys/cmd/rpc/filesys"
	"recorder/service/gateway/cmd/api/internal/svc"
	"recorder/service/gateway/cmd/api/internal/types"
	"recorder/service/gateway/httperror"
	"recorder/service/video/cmd/rpc/videoclient"
	"recorder/service/video/model"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) UploadLogic {
	return UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req types.UploadReq, data []byte) (*types.UploadResp, error) {
	// status:  0 还没上传 1 上传中 2 已上传完
	var status int
	var ok bool
	one, err := l.svcCtx.VideoClient.GetOne(l.ctx, &videoclient.GetOneReq{
		UserId: int64(req.UserId),
		Name: req.Name,
	})
	createdAt := one.CreatedAt
	switch err {
	case nil:
		if one.FinishedAt != "" {
			return nil, httperror.ErrVideoUploadFinished
		}
	case model.ErrNotFound:
		createdAt = time.Now().Format("2006-01-02 15:04:05")
	default:
		l.Logger.Infof("file %s upload failed: %v\n", req.Name, err)
		return nil, httperror.ErrVideoUploadFailed

	}

	resp, err := l.svcCtx.FileSysClient.StorePartly(l.ctx, &filesys.StorePartlyReq{
		UserId:    int64(req.UserId),
		Filename:  req.Name,
		Data:      data,
		CreatedAt: createdAt,
		Chunk:     int64(req.Chunk),
		Chunks:    int64(req.Chunks),
	})
	if err != nil {
		l.Logger.Infof("file %s upload failed: %v\n", req.Name, err)
		return nil, httperror.ErrVideoUploadFailed
	}

	status = 1
	if resp.IsFinished {
		status = 2
	}

	switch status {
	case 0:
		resp, err := l.svcCtx.VideoClient.Uploading(l.ctx, &videoclient.UploadingReq{
			UserId: int64(req.UserId),
			VideoName: req.Name,
			Path: resp.Path,
		})
		if err != nil {
			l.Logger.Infof("file %s upload failed: %v\n", req.Name, err)
			return nil, httperror.ErrVideoUploadFailed
		}

		ok = resp.Ok
	case 1:
		ok = true
	case 2:
		resp, err := l.svcCtx.VideoClient.Uploaded(l.ctx, &videoclient.UploadedReq{
			UserId: int64(req.UserId),
			VideoName: req.Name,
		})
		if err != nil {
			l.Logger.Infof("file %s upload failed: %v\n", req.Name, err)
			return nil, httperror.ErrVideoUploadFailed
		}

		ok = resp.Ok
	}

	return &types.UploadResp{
		Ok: ok,
	}, nil
}
