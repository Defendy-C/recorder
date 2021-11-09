package logic

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"recorder/service/file_sys/cmd/rpc/filesys"
	"recorder/service/gateway/httperror"
	"recorder/service/video/cmd/rpc/video"
	"recorder/util/validate"

	"recorder/service/gateway/cmd/api/internal/svc"
	"recorder/service/gateway/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) DownloadLogic {
	return DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadLogic) Download(req types.DownloadReq, w io.Writer) error {
	if validate.ValuesHasZero(req.Id, req.Chunk) {
		return httperror.ErrVideoDownloadFailed
	}

	v, err := l.svcCtx.VideoClient.GetOne(l.ctx, &video.GetOneReq{
		Id: int64(req.Id),
	})
	if err != nil {
		l.Logger.Infof("file videoId-chunk %d-%d download failed: %v\n", req.Id, req.Chunk, err)
		return httperror.ErrVideoDownloadFailed
	}

	client, err := l.svcCtx.FileSysClient.GetFile(l.ctx, &filesys.GetFileReq{
		Id:    v.FileId,
		Chunk: int64(req.Chunk),
	})
	if err != nil {
		l.Logger.Infof("file videoId-chunk %d-%d download failed: %v\n", req.Id, req.Chunk, err)
		return httperror.ErrVideoDownloadFailed
	}

	for resp, err := client.Recv(); err == nil; resp, err = client.Recv() {
		_, err = io.Copy(w, bytes.NewReader(resp.File))
		if err != nil {
			l.Logger.Infof("file videoId-chunk %d-%d download failed: %v\n", req.Id, req.Chunk, err)
			return httperror.ErrVideoDownloadFailed
		}

		// 分块发送, 避免占用内存
		w.(http.Flusher).Flush()
	}

	switch err {
	case io.EOF:
	default:
		l.Logger.Infof("file videoId-chunk %d-%d download failed: %v\n", req.Id, req.Chunk, err)
		return httperror.ErrVideoDownloadFailed
	}

	return nil
}
