package model

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"time"
)

type RealVideoModel interface {
	VideoModel
	FindVideosByUserIdCreatedAt(userId int64, createdAt time.Time, opt *ListOption) ([]*Video, error)
}

type ListOption struct {
	Limit      int64
	Offset     int64
	TotalPages int64
	TotalCount int64
}

func NewListOption(page int64, pageSize int64) *ListOption {
	limit := pageSize
	if pageSize == 0 {
		limit = 1
	}

	offset := (page - 1) * limit
	if page == 0 {
		offset = 0
	}

	return &ListOption{
		Limit: limit,
		Offset: offset,
	}
}

func NewRealVideoModel(conn sqlx.SqlConn, c cache.CacheConf) RealVideoModel {
	return &defaultVideoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`video`",
	}
}

func (m *defaultVideoModel) FindVideosByUserIdCreatedAt(userId int64, createdAt time.Time, opt *ListOption) ([]*Video, error) {
	panic("implement me")
}