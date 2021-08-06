package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	tagFieldNames          = builderx.RawFieldNames(&Tag{})
	tagRows                = strings.Join(tagFieldNames, ",")
	tagRowsExpectAutoSet   = strings.Join(stringx.Remove(tagFieldNames, "`id`", "`created_at`", "`updated_at`"), ",")
	tagRowsWithPlaceHolder = strings.Join(stringx.Remove(tagFieldNames, "`id`", "`created_at`", "`update_at`"), "=?,") + "=?"
)

type (
	TagModel interface {
		Insert(data Tag) (sql.Result, error)
		FindOne(id int64) (*Tag, error)
		Update(data Tag) error
		Delete(id int64) error
	}

	defaultTagModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Tag struct {
		Content    string    `db:"content"`     // 标签内容
		Count      int64     `db:"count"`       // 标签认同数量
		CreatedAt  time.Time `db:"created_at"`  // 创建时间
		UpdatedAt  time.Time `db:"updated_at"`  // 修改时间
		ID         int64     `db:"id"`          // id
		EntityType int64     `db:"entity_type"` // 标签的所属实体: 1: blog 2: comment 3: user
		EntityID   int64     `db:"entity_id"`   // 标签所属实体ID
	}

	WhereTag struct {
		IDs        []int64 `db:"id" operator:"in"`
		EntityIDs  []int64 `db:"entity_id" operator:"in"` // 实体ID
		EntityType *int32  `db:"entity_type"`
		Content    *string `db:"content"` // 内容的全量匹配
	}
)

func NewTagModel(conn sqlx.SqlConn) TagModel {
	return &defaultTagModel{
		conn:  conn,
		table: "`tag`",
	}
}

func (m *defaultTagModel) Insert(data Tag) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, tagRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Content, data.Count, data.EntityType, data.EntityID)
	return ret, err
}

func (m *defaultTagModel) FindOne(id int64) (*Tag, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tagRows, m.table)
	var resp Tag
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTagModel) Update(data Tag) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tagRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Content, data.Count, data.EntityType, data.EntityID, data.ID)
	return err
}

func (m *defaultTagModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
