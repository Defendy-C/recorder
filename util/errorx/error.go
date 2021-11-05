package errorx

import "github.com/tal-tech/go-zero/core/jsonx"

type BaseError struct {
	Code int
	Msg string
}

type HttpError struct {
	BaseError
	HttpCode int
}

func New(code int, msg string) *BaseError {
	return &BaseError{
		Code: code,
		Msg: msg,
	}
}

func (e *BaseError) Error() string {
	return e.Msg
}

func (e *BaseError) NewHttpError(httpCode int) *HttpError {
	return &HttpError{
		HttpCode: httpCode,
		BaseError: *e,
	}
}

func (e *HttpError) JSON() (int, interface{}) {
	j, err := jsonx.Marshal(e.BaseError)
	if err != nil {
		return e.HttpCode, e.BaseError
	}

	return e.HttpCode, string(j)
}