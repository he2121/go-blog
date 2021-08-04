package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/go-xorm/builder"
	"github.com/he2121/go-blog/common/sql-helper"
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
		Update(data Blog) error
		Delete(id int64) error
		MGetBlog(ids []int64) ([]*Blog, error)
		GetBlogList(where WhereBlog, option *sql_helper.Option) ([]*Blog, error)
		Count(where WhereBlog) (int32, error)
	}

	defaultBlogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Blog struct {
		Content   string    `db:"content"`    // 博客的内容
		Status    int64     `db:"status"`     // 博客状态 1:所有人可见 2. 仅自己可见 3. 删除
		FolderID  int64     `db:"folder_id"`  // 博客所属类别/文件夹，0 则无类别
		Extra     string    `db:"extra"`      // 一些额外的json数据
		CreatedAt time.Time `db:"created_at"` // 创建时间
		ID        int64     `db:"id"`         // id
		UserID    int64     `db:"user_id"`    // 博客的作者
		Title     string    `db:"title"`      // 博客标题
		IsFolder  int64     `db:"is_folder"`  // 0: 正常博客，1：博客类别/文件夹
		UpdatedAt time.Time `db:"updated_at"` // 修改时间
	}

	WhereBlog struct {
		IDs          []int64    `db:"id" operator:"in"`
		UserIDs      []int64    `db:"user_id" operator:"in"`
		Title        *string    `operator:"like"`
		IsFolder     *bool      `db:"is_folder"`
		Status       *int32     `db:"status"`
		FolderID     *int64     `db:"folder_id"`
		CreatedAtGTE *time.Time `db:"created_at" operator:">="`
	}
)

func NewBlogModel(conn sqlx.SqlConn) BlogModel {
	return &defaultBlogModel{
		conn:  conn,
		table: "`blog`",
	}
}

func (m *defaultBlogModel) Insert(data Blog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, blogRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Content, data.Status, data.FolderID, data.Extra, data.UserID, data.Title, data.IsFolder)
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

func (m *defaultBlogModel) MGetBlog(ids []int64) (blogs []*Blog, err error) {
	in := builder.In("id", ids)
	sqlStr, args, err := builder.MySQL().Select(blogRows).From(m.table).Where(in).ToSQL()
	if err != nil {
		return nil, err
	}
	blogs = []*Blog{}
	if err = m.conn.QueryRows(&blogs, sqlStr, args...); err != nil {
		return nil, err
	}
	return blogs, err
}

func (m *defaultBlogModel) GetBlogList(where WhereBlog, option *sql_helper.Option) (blogs []*Blog, err error) {
	conds, err := sql_helper.WrapWhere(where)
	if err != nil {
		return
	}
	b := builder.MySQL().Select(blogRows).From(m.table).Limit(option.Limit, option.Offset)
	for _, cond := range conds {
		b.Where(cond)
	}
	sqlStr, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}
	blogs = []*Blog{}
	if err = m.conn.QueryRows(&blogs, sqlStr, args...); err != nil {
		return nil, err
	}
	return blogs, err
}

func (m *defaultBlogModel) Count(where WhereBlog) (int32, error) {
	conds, err := sql_helper.WrapWhere(where)
	b := builder.MySQL().Select("count(*)").From(m.table)
	for _, cond := range conds {
		b.Where(cond)
	}
	sqlStr, args, err := b.ToSQL()
	if err != nil {
		return 0, err
	}
	var count int32
	if err := m.conn.QueryRow(&count, sqlStr, args...); err != nil {
		return 0, err
	}
	return count, nil
}

func (m *defaultBlogModel) Update(data Blog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, blogRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Content, data.Status, data.FolderID, data.Extra, data.UserID, data.Title, data.IsFolder, data.ID)
	return err
}

func (m *defaultBlogModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
