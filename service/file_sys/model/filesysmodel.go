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
	fileSysFieldNames          = builderx.RawFieldNames(&FileSys{})
	fileSysRows                = strings.Join(fileSysFieldNames, ",")
	fileSysRowsExpectAutoSet   = strings.Join(stringx.Remove(fileSysFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	fileSysRowsWithPlaceHolder = strings.Join(stringx.Remove(fileSysFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheFileSysIdPrefix   = "cache:fileSys:id:"
	cacheFileSysPathPrefix = "cache:fileSys:path:"
)

type (
	FileSysModel interface {
		Insert(data FileSys) (sql.Result, error)
		FindOne(id int64) (*FileSys, error)
		FindOneByPath(path string) (*FileSys, error)
		Update(data FileSys) error
		Delete(id int64) error
	}

	defaultFileSysModel struct {
		sqlc.CachedConn
		table string
	}

	FileSys struct {
		Id         int64     `db:"id"`
		Path       string    `db:"path"`        // 存放路径
		CreatedAt  time.Time `db:"created_at"`  // 创建时间
		FinishedAt time.Time `db:"finished_at"` // 完成时间
		TotalChunk int64     `db:"total_chunk"` // 文件分段数
	}
)

func NewFileSysModel(conn sqlx.SqlConn, c cache.CacheConf) FileSysModel {
	return &defaultFileSysModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`file_sys`",
	}
}

func (m *defaultFileSysModel) Insert(data FileSys) (sql.Result, error) {
	fileSysPathKey := fmt.Sprintf("%s%v", cacheFileSysPathPrefix, data.Path)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, fileSysRowsExpectAutoSet)
		return conn.Exec(query, data.Path, data.CreatedAt, data.FinishedAt, data.TotalChunk)
	}, fileSysPathKey)
	return ret, err
}

func (m *defaultFileSysModel) FindOne(id int64) (*FileSys, error) {
	fileSysIdKey := fmt.Sprintf("%s%v", cacheFileSysIdPrefix, id)
	var resp FileSys
	err := m.QueryRow(&resp, fileSysIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", fileSysRows, m.table)
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

func (m *defaultFileSysModel) FindOneByPath(path string) (*FileSys, error) {
	fileSysPathKey := fmt.Sprintf("%s%v", cacheFileSysPathPrefix, path)
	var resp FileSys
	err := m.QueryRowIndex(&resp, fileSysPathKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `path` = ? limit 1", fileSysRows, m.table)
		if err := conn.QueryRow(&resp, query, path); err != nil {
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

func (m *defaultFileSysModel) Update(data FileSys) error {
	fileSysIdKey := fmt.Sprintf("%s%v", cacheFileSysIdPrefix, data.Id)
	fileSysPathKey := fmt.Sprintf("%s%v", cacheFileSysPathPrefix, data.Path)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, fileSysRowsWithPlaceHolder)
		return conn.Exec(query, data.Path, data.CreatedAt, data.FinishedAt, data.TotalChunk, data.Id)
	}, fileSysPathKey, fileSysIdKey)
	return err
}

func (m *defaultFileSysModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	fileSysIdKey := fmt.Sprintf("%s%v", cacheFileSysIdPrefix, id)
	fileSysPathKey := fmt.Sprintf("%s%v", cacheFileSysPathPrefix, data.Path)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, fileSysIdKey, fileSysPathKey)
	return err
}

func (m *defaultFileSysModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheFileSysIdPrefix, primary)
}

func (m *defaultFileSysModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", fileSysRows, m.table)
	return conn.QueryRow(v, query, primary)
}
