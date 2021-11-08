package httperror

import (
	"net/http"
	"recorder/util/errorx"
)

var (
	ErrUserNotExists = errorx.New(10001, "用户不存在").NewHttpError(http.StatusBadRequest)
	ErrLoginFailed   = errorx.New(10002, "登录失败, 请稍后重试").NewHttpError(http.StatusBadRequest)
	ErrVideoUploadFailed  = errorx.New(20001, "文件上传失败").NewHttpError(http.StatusBadRequest)
	ErrVideoUploadFinished = errorx.New(20002, "文件已经上传完成").NewHttpError(http.StatusBadRequest)
)