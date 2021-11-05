package model

import (
	"errors"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound
var ErrValidate = errors.New("validate parameters failed")