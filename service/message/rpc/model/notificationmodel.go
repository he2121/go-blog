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

	"github.com/he2121/go-blog/service/message/rpc/message"
)

var (
	notificationFieldNames          = builderx.RawFieldNames(&Notification{})
	notificationRows                = strings.Join(notificationFieldNames, ",")
	notificationRowsExpectAutoSet   = strings.Join(stringx.Remove(notificationFieldNames, "`id`", "`created_at`", "`updated_at`"), ",")
	notificationRowsWithPlaceHolder = strings.Join(stringx.Remove(notificationFieldNames, "`id`", "`created_at`", "`updated_at`"), "=?,") + "=?"
)

type (
	NotificationModel interface {
		Insert(data Notification) (sql.Result, error)
		FindOne(id int64) (*Notification, error)
		Update(data Notification) error
		Delete(id int64) error
		MGetNotification(ids []int64) ([]*Notification, error)
		GetNotificationList(where WhereNotification, option *sql_helper.Option) ([]*Notification, error)
		Count(where WhereNotification) (int32, error)
	}

	defaultNotificationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Notification struct {
		EntityID   int64     `db:"entity_id"`   // 触发事件的实体 ID
		TriggerID  int64     `db:"trigger_id"`  // 触发事件的用户 ID
		Status     int64     `db:"status"`      // 1：未读 2：已读
		Extra      string    `db:"extra"`       // 一些额外的json数据
		ID         int64     `db:"id"`          // id
		UserID     int64     `db:"user_id"`     // 通知用户 ID
		EventType  int64     `db:"event_type"`  // 触发该通知的事件类型：1：点赞 2: 评论 3: 发帖 4：关注
		EntityType int64     `db:"entity_type"` // 触发事件的实体类型
		CreatedAt  time.Time `db:"created_at"`  // 创建时间
		UpdatedAt  time.Time `db:"updated_at"`  // 修改时间
	}

	WhereNotification struct {
		UserIDs   []int64            `db:"user_id" operator:"in"` //
		EventType *message.EventType `db:"event_type"`            // 事件类型
	}
)

func NewNotificationModel(conn sqlx.SqlConn) NotificationModel {
	return &defaultNotificationModel{
		conn:  conn,
		table: "`notification`",
	}
}

func (m *defaultNotificationModel) Insert(data Notification) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, notificationRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.EntityID, data.TriggerID, data.Status, data.Extra, data.UserID, data.EventType, data.EntityType)
	return ret, err
}

func (m *defaultNotificationModel) FindOne(id int64) (*Notification, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", notificationRows, m.table)
	var resp Notification
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

func (m *defaultNotificationModel) Update(data Notification) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, notificationRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.EntityID, data.TriggerID, data.Status, data.Extra, data.UserID, data.EventType, data.EntityType, data.ID)
	return err
}

func (m *defaultNotificationModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultNotificationModel) MGetNotification(ids []int64) ([]*Notification, error) {
	in := builder.In("id", ids)
	sqlStr, args, err := builder.MySQL().Select(notificationRows).From(m.table).Where(in).ToSQL()
	if err != nil {
		return nil, err
	}
	var Notifications []*Notification
	if err = m.conn.QueryRows(&Notifications, sqlStr, args...); err != nil {
		return nil, err
	}
	return Notifications, nil
}

func (m *defaultNotificationModel) GetNotificationList(where WhereNotification, option *sql_helper.Option) (Notifications []*Notification, err error) {
	conds, err := sql_helper.WrapWhere(where)
	if err != nil {
		return
	}
	b := builder.MySQL().Select(notificationRows).From(m.table).Limit(option.Limit, option.Offset)
	for _, cond := range conds {
		b.Where(cond)
	}
	sqlStr, args, err := b.ToSQL()
	if err != nil {
		return nil, err
	}
	Notifications = []*Notification{}
	if err = m.conn.QueryRows(&Notifications, sqlStr, args...); err != nil {
		return nil, err
	}
	return Notifications, err
}

func (m *defaultNotificationModel) Count(where WhereNotification) (int32, error) {
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
