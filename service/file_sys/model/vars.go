package model

import (
	"errors"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound
var ErrValidate = errors.New("validate parameters failed")
func ErrFileExists(filename string) error {
	return errors.New(fmt.Sprintf("the file %s has existed", filename))
}
func ErrDirExists(dir string) error {
	return errors.New(fmt.Sprintf("the file %s has existed", dir))
}
