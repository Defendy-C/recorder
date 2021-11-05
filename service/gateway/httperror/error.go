package httperror

import (
	"net/http"
	"recorder/util/errorx"
)

var (
	ErrUserNotExists = errorx.New(10001, "用户不存在").NewHttpError(http.StatusBadRequest)
	ErrLoginFailed   = errorx.New(10002, "登录失败, 请稍后重试").NewHttpError(http.StatusBadRequest)
)