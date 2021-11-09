package logic

import (
	"context"
	"io/ioutil"
	"recorder/service/video/model"
	"recorder/util/validate"

	"recorder/service/file_sys/cmd/rpc/filesys"
	"recorder/service/file_sys/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetFileInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileInfoLogic {
	return &GetFileInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFileInfoLogic) GetFileInfo(in *filesys.GetFileInfoReq) (*filesys.GetFileInfoResp, error) {
	if validate.ValuesHasZero(in.Id) {
		return nil, model.ErrValidate
	}

	fs, err := l.svcCtx.FileSysModel.FindOne(in.Id)
	if err != nil {
		return nil, err
	}

	var isFinished bool
	if !fs.FinishedAt.IsZero() {
		isFinished = true
	}

	realPath := storageSpace + fs.Path
	fis, err := ioutil.ReadDir(realPath)
	if err != nil {
		return nil, err
	}

	return &filesys.GetFileInfoResp{
		CreatedAt: validate.DateToString(&fs.CreatedAt),
		IsFinished: isFinished,
		TotalChunks: fs.TotalChunk,
		CurrentChunks: int64(len(fis)),
	}, nil
}
