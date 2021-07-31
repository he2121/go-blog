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
	blogFieldNames          = builderx.RawFieldNames(&Blog{})
	blogRows                = strings.Join(blogFieldNames, ",")
	blogRowsExpectAutoSet   = strings.Join(stringx.Remove(blogFieldNames, "`id`", "`created_at`", "`updated_at`"), ",")
	blogRowsWithPlaceHolder = strings.Join(stringx.Remove(blogFieldNames, "`id`", "`created_at`", "`updated_at`"), "=?,") + "=?"
)

type (
	BlogModel interface {
		Insert(data Blog) (sql.Result, error)
		FindOne(id int64) (*Blog, error)
		FindOneByCreatedAt(createdAt time.Time) (*Blog, error)
		Update(data Blog) error
		Delete(id int64) error
	}

	defaultBlogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Blog struct {
		LikeCount int64     `db:"like_count"` // 喜欢的人数
		CreatedAt time.Time `db:"created_at"` // 创建时间
		Id        int64     `db:"id"`         // id
		UserId    int64     `db:"user_id"`    // 博客的作者
		IsFolder  int64     `db:"is_folder"`  // 0: 正常博客，1：博客类别/文件夹
		Extra     string    `db:"extra"`      // 一些额外的json数据
		UpdatedAt time.Time `db:"updated_at"` // 修改时间
		Title     string    `db:"title"`      // 博客标题
		Content   string    `db:"content"`    // 博客的内容
		Status    int64     `db:"status"`     // 博客状态 1:所有人可见 2. 仅自己可见 3. 删除
		FolderId  int64     `db:"folder_id"`  // 博客所属类别/文件夹，0 则无类别
	}
)

func NewBlogModel(conn sqlx.SqlConn) BlogModel {
	return &defaultBlogModel{
		conn:  conn,
		table: "`blog`",
	}
}

func (m *defaultBlogModel) Insert(data Blog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, blogRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.LikeCount, data.CreatedAt, data.UserId, data.IsFolder, data.Extra, data.UpdatedAt, data.Title, data.Content, data.Status, data.FolderId)
	return ret, err
}

func (m *defaultBlogModel) FindOne(id int64) (*Blog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", blogRows, m.table)
	var resp Blog
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

func (m *defaultBlogModel) FindOneByCreatedAt(createdAt time.Time) (*Blog, error) {
	var resp Blog
	query := fmt.Sprintf("select %s from %s where `created_at` = ? limit 1", blogRows, m.table)
	err := m.conn.QueryRow(&resp, query, createdAt)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBlogModel) Update(data Blog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, blogRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.LikeCount, data.CreatedAt, data.UserId, data.IsFolder, data.Extra, data.UpdatedAt, data.Title, data.Content, data.Status, data.FolderId, data.Id)
	return err
}

func (m *defaultBlogModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
