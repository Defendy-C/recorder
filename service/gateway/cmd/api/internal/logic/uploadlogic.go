package logic

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"io"
	"recorder/service/file_sys/cmd/rpc/filesysclient"
	"recorder/service/gateway/cmd/api/internal/svc"
	"recorder/service/gateway/cmd/api/internal/types"
	"recorder/service/gateway/httperror"
	"recorder/util/validate"
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

func (l *UploadLogic) Upload(req types.UploadReq, rd io.ReadCloser) (*types.UploadResp, error) {
	// 1. 验证req, 创建file
	if validate.ValuesHasZero(req.Id, req.Chunk) {
		return nil, httperror.ErrVideoUploadFailed
	}

	closeFd := func() {
		err := rd.Close()
		if err != nil {
			l.Logger.Errorf("file close failed: %v", err)
		}
	}

	// 2. 传输数据
	// if failed: rd.close
	client, err := l.svcCtx.FileSysClient.Store(l.ctx)
	if err != nil {
		l.Logger.Errorf("video_id-chunk: %d-%d upload failed: %v\n", req.Id, req.Chunk, err)
		closeFd()
		return nil, httperror.ErrVideoUploadFailed
	}

	err = client.Send(&filesysclient.StoreReq{
		Id: int64(req.Id),
		Chunk: int64(req.Chunk),
	})
	if err != nil {
		closeFd()
		l.Logger.Errorf("video_id-chunk: %d-%d upload failed: %v\n", req.Id, req.Chunk, err)
		return nil, httperror.ErrVideoUploadFailed
	}

	buf := make([]byte, 1 * 1024)
	c := 0
	for n, err := rd.Read(buf);err == nil;n, err = rd.Read(buf) {
		c++
		fmt.Println(c, n)
		if n == 0 {
			break
		}

		err := client.Send(&filesysclient.StoreReq{
			Data: buf[0:n],
		})
		if err == io.EOF {
			break
		}
		if err != nil {
			l.Logger.Errorf("video_id-chunk: %d-%d upload failed: %v\n", req.Id, req.Chunk, err)
			closeFd()
			return nil, httperror.ErrVideoUploadFailed
		}
	}

	switch err {
	case nil:
	case io.EOF:
	default:
		l.Logger.Errorf("video_id-chunk: %d-%d upload failed: %v\n", req.Id, req.Chunk, err)
		return nil, httperror.ErrVideoUploadFailed
	}

	_, err = client.CloseAndRecv()
	if err != nil {
		l.Logger.Errorf("video_id-chunk: %d-%d upload failed: %v\n", req.Id, req.Chunk, err)
	}

	return &types.UploadResp{
		Ok: true,
	}, nil
}
