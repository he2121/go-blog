package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/core/conf"

	"user/internal/config"
	"user/internal/server"
	"user/internal/svc"
	"user/user"
)

var srv *server.UserServiceServer

func init() {
	//flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv = server.NewUserServiceServer(ctx)
}

func TestSendEmailCode(t *testing.T) {
	resp, err := srv.SendEmailCode(context.TODO(), &user.SendEmailCodeReq{Email: "1070259395@qq.com"})
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestAuthByEmailCode(t *testing.T) {
	resp, err := srv.AuthByEmailCode(context.TODO(), &user.AuthByEmailCodeReq{
		Email: "1070259395@qq.com",
		Code:  "607557",
	})
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestUpdateUserInfo(t *testing.T) {
	resp, err := srv.UpdateUserInfo(context.TODO(), &user.UpdateUserInfoReq{
		ID:        1,
		Phone:     "123",
		Name:      "test",
		Gender:    2,
		BirthDate: 0,
		UserExtra: nil,
		PassWord:  "123456",
	})
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestAuthByIDPassword(t *testing.T) {
	resp, err := srv.AuthByIDPassword(context.TODO(), &user.AuthByIDPasswordReq{
		ID:       1,
		Password: "1234567",
	})
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestMGetUserInfo(t *testing.T) {
	resp, err := srv.MGetUserInfo(context.TODO(), &user.MGetUserInfoReq{IDs: []int64{1}})
	assert.Nil(t, err)
	fmt.Printf("%+v", resp)
}
