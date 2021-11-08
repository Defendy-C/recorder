package logic

import (
	"context"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"recorder/service/file_sys/cmd/rpc/filesys"
	"recorder/service/file_sys/cmd/rpc/internal/svc"
	"recorder/util/uuid"
	"recorder/util/validate"

	"github.com/tal-tech/go-zero/core/logx"
)

type StorePartlyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStorePartlyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StorePartlyLogic {
	return &StorePartlyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StorePartlyLogic) StorePartly(in *filesys.StorePartlyReq) (*filesys.StorePartlyResp, error) {
	if validate.ValuesHasZero(in.UserId) {
		return nil, ErrValidate
	}

	t := validate.StringToDate(in.CreatedAt)
	if t.IsZero() {
		return nil, ErrValidate
	}

	if in.Filename == "" {
		in.Filename = uuid.UniqueFilename(64)
	}

	path := storageSpace + strings.Join([]string{strconv.Itoa(int(in.UserId)), t.Format("2006-01-02"), in.Filename}, "\\")
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var isFinished bool
	filename := path + "\\" + strconv.Itoa(int(in.Chunk))
	count := len(fs)
	isFinished = count == int(in.Chunks) - 1
	_, err = os.Stat(filename)
	if err == nil {
		return nil, ErrFileExists(filename)
	}

	f, err := os.OpenFile(path, os.O_CREATE | os.O_WRONLY, 0777)
	if err != nil {
		return nil, err
	}

	_, err = f.Write(in.Data)
	if err != nil {
		return nil, err
	}

	return &filesys.StorePartlyResp{
		Path: filename,
		IsFinished: isFinished,
	}, nil
}