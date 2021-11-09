package logic

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"recorder/service/file_sys/cmd/rpc/filesys"
	"recorder/service/file_sys/cmd/rpc/internal/svc"
	"recorder/service/file_sys/model"
	"strconv"

	"github.com/tal-tech/go-zero/core/logx"
)

type StoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreLogic {
	return &StoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StoreLogic) Store(stream filesys.FileSys_StoreServer) error {
	// 第一次返回id, chunk
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	// id 校验
	obj, err := l.svcCtx.FileSysModel.FindOne(req.Id)
	if err != nil {
		return err
	}
	res := obj.Path
	// 文件检查
	filename := obj.Path + "\\" + strconv.Itoa(int(req.Chunk))
	_, err = os.Stat(filename)
	if err == nil {
		return model.ErrFileExists(filename)
	}

	if ok := os.IsNotExist(err); !ok {
		return err
	}

	var isFinished bool
	fs, err := ioutil.ReadDir(obj.Path)
	if int64(len(fs) + 1) == obj.TotalChunk {
		isFinished = true
	}
	// 打开文件
	fmt.Println(os.Getwd())
	f, err := os.OpenFile(storageSpace + filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	// 关闭文件流
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			l.Logger.Errorf("file close failed: %v\n", err)
		}
	}(f)
	// 数据接收
	buf := bufio.NewWriter(f)
	for req, err = stream.Recv();err == nil; req, err = stream.Recv() {
		fmt.Println(8)
		_, err = buf.Write(req.Data)
		if err != nil {
			fmt.Println("ww---------------")
			break
		}

		err = buf.Flush()
		if err != nil {
			break
		}
	}
	e := stream.SendAndClose(&filesys.StoreResp{
		Path: res,
		IsFinished: isFinished,
	})
	if e != nil {
		l.Logger.Errorf("rpc stream close failed: %v\n", err)
	}

	// 如果err非空, 则在这里进行处理
	switch err {
	case io.EOF:
	case nil:
	default:
		_ = os.Remove(filename)
		return err
	}

	// 如 e 非空则在这里处理
	return e
}
