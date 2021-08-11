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
	messageFieldNames          = builderx.RawFieldNames(&Message{})
	messageRows                = strings.Join(messageFieldNames, ",")
	messageRowsExpectAutoSet   = strings.Join(stringx.Remove(messageFieldNames, "`id`", "`created_at`", "`updated_at`"), ",")
	messageRowsWithPlaceHolder = strings.Join(stringx.Remove(messageFieldNames, "`id`", "`created_at`", "`updated_at`"), "=?,") + "=?"
)

type (
	MessageModel interface {
		Insert(data Message) (sql.Result, error)
		FindOne(id int64) (*Message, error)
		Update(data Message) error
		Delete(id int64) error
		MGetMessage(ids []int64) ([]*Message, error)
		GetMessageList(where WhereMessage, option *sql_helper.Option) ([]*Message, error)
		Count(where WhereMessage) (int32, error)
	}

	defaultMessageModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Message struct {
		FromID      int64          `db:"from_id"`      // 发送者 ID
		SessionID   string         `db:"session_id"`   // 会话 ID， 由发送者接受者ID组合，小 ID：大 ID
		MessageType int64          `db:"message_type"` // 消息类型 1: 普通私信， 2：群聊
		Content     sql.NullString `db:"content"`      // 私信内容
		CreatedAt   time.Time      `db:"created_at"`   // 创建时间
		UpdatedAt   time.Time      `db:"updated_at"`   // 修改时间
		ID          int64          `db:"id"`           // id
		ToID        int64          `db:"to_id"`        // 接受者 ID
		Status      int64          `db:"status"`       // 1：未读 2：已读 3：撤回
		Extra       string         `db:"extra"`        // 一些额外的json数据
	}

	WhereMessage struct {
		SessionID *string `db:"session_id"`
		ToID      *int64  `db:"to_id"`
		Status    *int32  `db:"status"`
		FromID    *int64  `db:"from_id"`
	}
)

func (m *defaultMessageModel) MGetMessage(ids []int64) ([]*Message, error) {
	in := builder.In("id", ids)
	sqlStr, args, err := builder.MySQL().Select(messageRows).From(m.table).Where(in).ToSQL()
	if err != nil {
		return nil, err
	}
	var messages []*Message
	if err = m.conn.QueryRows(&messages, sqlStr, args...); err != nil {
		return nil, err
	}
	return messages, nil
}

func (m *defaultMessageModel) GetMessageList(where WhereMessage, option *sql_helper.Option) (messages []*Message, err error) {
	conds, err := sql_helper.WrapWhere(where)
	if err != nil {
		return
	}
	b := builder.MySQL().Select(messageRows).From(m.table).Limit(option.Limit, option.Offset)
	for _, cond := range conds {
		b.Where(cond)
	}
	sqlStr, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}
	messages = []*Message{}
	if err = m.conn.QueryRows(&messages, sqlStr, args...); err != nil {
		return nil, err
	}
	return messages, err
}

func (m *defaultMessageModel) Count(where WhereMessage) (int32, error) {
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

func NewMessageModel(conn sqlx.SqlConn) MessageModel {
	return &defaultMessageModel{
		conn:  conn,
		table: "`message`",
	}
}

func (m *defaultMessageModel) Insert(data Message) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, messageRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.FromID, data.SessionID, data.MessageType, data.Content, data.ToID, data.Status, data.Extra)
	return ret, err
}

func (m *defaultMessageModel) FindOne(id int64) (*Message, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", messageRows, m.table)
	var resp Message
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

func (m *defaultMessageModel) Update(data Message) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, messageRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.FromID, data.SessionID, data.MessageType, data.Content, data.ToID, data.Status, data.Extra, data.ID)
	return err
}

func (m *defaultMessageModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
