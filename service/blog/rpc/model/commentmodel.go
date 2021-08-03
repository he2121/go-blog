package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/go-xorm/builder"
	sql_helper "github.com/he2121/go-blog/common/sql-helper"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	commentFieldNames          = builderx.RawFieldNames(&Comment{})
	commentRows                = strings.Join(commentFieldNames, ",")
	commentRowsExpectAutoSet   = strings.Join(stringx.Remove(commentFieldNames, "`id`", "`created_at`"), ",")
	commentRowsWithPlaceHolder = strings.Join(stringx.Remove(commentFieldNames, "`id`", "`created_at`"), "=?,") + "=?"
)

type (
	CommentModel interface {
		Insert(data Comment) (sql.Result, error)
		FindOne(id int64) (*Comment, error)
		Update(data Comment) error
		Delete(id int64) error
	}

	defaultCommentModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Comment struct {
		Content    string    `db:"content"`      // 评论的内容
		Status     int64     `db:"status"`       // 评论状态 1:正常 2. 修改过 3. 删除
		CreatedAt  time.Time `db:"created_at"`   // 创建时间
		Extra      string    `db:"extra"`        // 一些额外的json数据
		Id         int64     `db:"id"`           // id
		Type       int64     `db:"type"`         // 评论类型：1. 对博客的评论 2. 对评论的评论 3. 对用户的评论
		ToId       int64     `db:"to_id"`        // 此评论所归属的id,若type是博客，此id是评论的博客
		FromUserId int64     `db:"from_user_id"` // 评论发起者
		ToUserId   int64     `db:"to_user_id"`   // 评论所回应的人，若是博客则是写博客的人ID
	}

	WhereComment struct {
		IDs          []int64    `db:"id" operator:"in"`
		UserIDs      []int64    `db:"user_id" operator:"in"`
		Title        *string    `operator:"like"`
		IsFolder     *bool      `db:"is_folder"`
		Status       *int32     `db:"title"`
		FolderID     *int64     `db:"folder_id"`
		CreatedAtGTE *time.Time `db:"created_at" operator:">="`
	}
)

func NewCommentModel(conn sqlx.SqlConn) CommentModel {
	return &defaultCommentModel{
		conn:  conn,
		table: "`comment`",
	}
}

func (m *defaultCommentModel) Insert(data Comment) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, commentRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Content, data.Status, data.CreatedAt, data.Extra, data.Type, data.ToId, data.FromUserId, data.ToUserId)
	return ret, err
}

func (m *defaultCommentModel) FindOne(id int64) (*Comment, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentRows, m.table)
	var resp Comment
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

func (m *defaultCommentModel) MGetComment(ids []int64) (comments []*Comment, err error) {
	in := builder.In("id", ids)
	sqlStr, args, err := builder.Select(commentRows).From(m.table).Where(in).ToSQL()
	if err != nil {
		return nil, err
	}
	comments = []*Comment{}
	if err = m.conn.QueryRows(&comments, sqlStr, args); err != nil {
		return nil, err
	}
	return comments, err
}

func (m *defaultCommentModel) GetCommentList(where WhereComment, option *sql_helper.Option) (comments []*Comment, err error) {
	conds, err := sql_helper.WrapWhere(where)
	if err != nil {
		return
	}
	b := builder.Select(commentRows).From(m.table).Limit(option.Offset, option.Limit)
	for _, cond := range conds {
		b.Where(cond)
	}
	sqlStr, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}
	comments = []*Comment{}
	if err = m.conn.QueryRows(&comments, sqlStr, args); err != nil {
		return nil, err
	}
	return comments, err
}

func (m *defaultCommentModel) Update(data Comment) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, commentRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Content, data.Status, data.CreatedAt, data.Extra, data.Type, data.ToId, data.FromUserId, data.ToUserId, data.Id)
	return err
}

func (m *defaultCommentModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
