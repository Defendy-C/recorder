package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	videoFieldNames          = builderx.RawFieldNames(&Video{})
	videoRows                = strings.Join(videoFieldNames, ",")
	videoRowsExpectAutoSet   = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	videoRowsWithPlaceHolder = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheVideoIdPrefix         = "cache:video:id:"
	cacheVideoUserIdNamePrefix = "cache:video:userId:name:"
)

type (
	VideoModel interface {
		Insert(data Video) (sql.Result, error)
		FindOne(id int64) (*Video, error)
		FindOneByUserIdName(userId int64, name string) (*Video, error)
		Update(data Video) error
		Delete(id int64) error
	}

	defaultVideoModel struct {
		sqlc.CachedConn
		table string
	}

	Video struct {
		Id         int64     `db:"id"`
		Name       string    `db:"name"`        // 视频名
		UserId     int64     `db:"user_id"`     // 用户Id
		Path       string    `db:"path"`        // 存放路径
		CreatedAt  time.Time `db:"created_at"`  // 创建时间
		FinishedAt time.Time `db:"finished_at"` // 完成时间
	}
)

func NewVideoModel(conn sqlx.SqlConn, c cache.CacheConf) VideoModel {
	return &defaultVideoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`video`",
	}
}

func (m *defaultVideoModel) Insert(data Video) (sql.Result, error) {
	videoUserIdNameKey := fmt.Sprintf("%s%v:%v", cacheVideoUserIdNamePrefix, data.UserId, data.Name)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, videoRowsExpectAutoSet)
		return conn.Exec(query, data.Name, data.UserId, data.Path, data.CreatedAt, data.FinishedAt)
	}, videoUserIdNameKey)
	return ret, err
}

func (m *defaultVideoModel) FindOne(id int64) (*Video, error) {
	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, id)
	var resp Video
	err := m.QueryRow(&resp, videoIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) FindOneByUserIdName(userId int64, name string) (*Video, error) {
	videoUserIdNameKey := fmt.Sprintf("%s%v:%v", cacheVideoUserIdNamePrefix, userId, name)
	var resp Video
	err := m.QueryRowIndex(&resp, videoUserIdNameKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `name` = ? limit 1", videoRows, m.table)
		if err := conn.QueryRow(&resp, query, userId, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) Update(data Video) error {
	videoUserIdNameKey := fmt.Sprintf("%s%v:%v", cacheVideoUserIdNamePrefix, data.UserId, data.Name)
	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, videoRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.UserId, data.Path, data.CreatedAt, data.FinishedAt, data.Id)
	}, videoIdKey, videoUserIdNameKey)
	return err
}

func (m *defaultVideoModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, id)
	videoUserIdNameKey := fmt.Sprintf("%s%v:%v", cacheVideoUserIdNamePrefix, data.UserId, data.Name)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, videoIdKey, videoUserIdNameKey)
	return err
}

func (m *defaultVideoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheVideoIdPrefix, primary)
}

func (m *defaultVideoModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
	return conn.QueryRow(v, query, primary)
}
