package model

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	userFieldNames          = builderx.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`created_at`", "`updated_at`", "`birth_date`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`created_at`", "`updated_at`"), "=?,") + "=?"

	cacheUserIdPrefix    = "cache:user:id:"
	cacheUserEmailPrefix = "cache:user:email:"
)

type (
	UserModel interface {
		Insert(data User) (sql.Result, error)
		FindOne(id int64) (*User, error)
		FindOneByEmail(email string) (*User, error)
		Update(data User) error
		Delete(id int64) error
		MGetUser(ids []int64) ([]*User, error)
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Phone     string    `db:"phone"`      // 用户手机
		Name      string    `db:"name"`       // 用户名
		Status    int32     `db:"status"`     // 用户状态 1:正常，2:停用
		Gender    int32     `db:"gender"`     // 性别 1:男 2:女
		CreatedAt time.Time `db:"created_at"` // 创建时间
		ID        int64     `db:"id"`         // id
		Email     string    `db:"email"`      // 注册邮箱
		Extra     string    `db:"extra"`      // 一些额外的json数据
		UpdatedAt time.Time `db:"updated_at"` // 修改时间
		Password  string    `db:"password"`   // 加 salt Hash密码
		BirthDate time.Time `db:"birth_date"` // 出生日期
	}
)

func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Insert(data User) (sql.Result, error) {
	userEmailKey := fmt.Sprintf("%s%v", cacheUserEmailPrefix, data.Email)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.Exec(query, data.Phone, data.Name, data.Status, data.Gender, data.Email, data.Extra, data.Password)
	}, userEmailKey)
	return ret, err
}

func (m *defaultUserModel) FindOne(id int64) (*User, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	var resp User
	err := m.QueryRow(&resp, userIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
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

func (m *defaultUserModel) MGetUser(ids []int64) (users []*User, err error) {
	var args strings.Builder
	args.WriteString("(")
	for i, id := range ids {
		args.WriteString(strconv.FormatInt(id, 10))
		if i != len(ids)-1 {
			args.WriteString(",")
		}
	}
	args.WriteString(")")
	query := fmt.Sprintf("select %s from %s where `id` in %s", userRows, m.table, args.String())
	users = []*User{}
	if err = m.QueryRowsNoCache(&users, query); err != nil {
		return nil, err
	}
	return users, err
}

func (m *defaultUserModel) FindOneByEmail(email string) (*User, error) {
	userEmailKey := fmt.Sprintf("%s%v", cacheUserEmailPrefix, email)
	var resp User
	err := m.QueryRowIndex(&resp, userEmailKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", userRows, m.table)
		if err := conn.QueryRow(&resp, query, email); err != nil {
			return nil, err
		}
		return resp.ID, nil
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

func (m *defaultUserModel) Update(data User) error {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.ID)
	userEmailKey := fmt.Sprintf("%s%v", cacheUserEmailPrefix, data.Email)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.Exec(query, data.Phone, data.Name, data.Status, data.Gender, data.Email, data.Extra, data.Password, data.BirthDate, data.ID)
	}, userIdKey, userEmailKey)
	return err
}

func (m *defaultUserModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	userEmailKey := fmt.Sprintf("%s%v", cacheUserEmailPrefix, data.Email)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, userEmailKey, userIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRow(v, query, primary)
}
