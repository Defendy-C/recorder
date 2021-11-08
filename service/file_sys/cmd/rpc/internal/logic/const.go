package logic

import (
	"errors"
	"fmt"
)

const storageSpace = "../../../../storage"

var ErrValidate = errors.New("validate parameters failed")

func ErrFileExists(filename string) error {
	return errors.New(fmt.Sprintf("the file %s has existed", filename))
}