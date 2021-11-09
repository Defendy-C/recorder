package logic

import (
	"bufio"
	"context"
	"io"
	"os"
	"strconv"

	"recorder/service/file_sys/cmd/rpc/filesys"
	"recorder/service/file_sys/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileLogic {
	return &GetFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFileLogic) GetFile(in *filesys.GetFileReq, stream filesys.FileSys_GetFileServer) error {
	one, err := l.svcCtx.FileSysModel.FindOne(in.Id)
	if err != nil {
		return err
	}

	f, err := os.Open(one.Path + strconv.Itoa(int(in.Chunk)))
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			l.Logger.Errorf("file %s close failed: %v\n", one.Path, err)
		}
	}(f)

	r := bufio.NewReader(f)
	buf := make([]byte, 1024 * 1024)
	for _, err = r.Read(buf); err == nil; _, err = r.Read(buf) {
		err = stream.Send(&filesys.GetFileResp{
			File: buf,
		})
		if err != nil {
			return err
		}
	}

	switch err {
	case io.EOF:
	default:
		return err
	}

	return nil
}
