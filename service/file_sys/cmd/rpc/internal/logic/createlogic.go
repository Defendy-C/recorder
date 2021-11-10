package logic

import (
	"context"
	"os"
	"recorder/service/file_sys/cmd/rpc/filesys"
	"recorder/service/file_sys/cmd/rpc/internal/svc"
	"recorder/service/file_sys/model"
	"recorder/util/uuid"
	"recorder/util/validate"
	"strconv"
	"strings"

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

func (l *CreateLogic) Create(in *filesys.CreateReq) (*filesys.CreateResp, error) {
	createdAt := validate.StringToDate(in.CreatedAt)
	if createdAt.IsZero() || validate.ValuesHasZero(in.UserId, in.TotalChunks) {
		return nil, model.ErrValidate
	}

	path := strings.Join([]string{strconv.Itoa(int(in.UserId)), createdAt.Format("2006-01-02"), uuid.UniqueFilename(64)}, "\\")
	_, err := l.svcCtx.FileSysModel.FindOneByPath(path)
	switch err {
	case model.ErrNotFound:
	case nil:
		return nil, model.ErrDirExists(path)
	default:
		return nil, err
	}

	// 创建目录
	err = os.MkdirAll(storageSpace + path, 0777)
	if err != nil {
		return nil, err
	}

	res, err := l.svcCtx.FileSysModel.Insert(model.FileSys{
		Path: path,
		CreatedAt: createdAt,
		TotalChunk: in.TotalChunks,
	})
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &filesys.CreateResp{
		Id: id,
	}, nil
}
