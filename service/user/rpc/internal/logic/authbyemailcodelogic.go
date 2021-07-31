package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/tal-tech/go-zero/core/logx"

	"user/internal/helper"
	"user/internal/svc"
	"user/model"
	"user/user"
)

type AuthByEmailCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthByEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthByEmailCodeLogic {
	return &AuthByEmailCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  邮箱登陆认证 (无则注册)
func (l *AuthByEmailCodeLogic) AuthByEmailCode(in *user.AuthByEmailCodeReq) (*user.AuthByEmailCodeResp, error) {
	code, err := l.svcCtx.Redis.Get(fmt.Sprintf("email:code:%s", in.Email))
	if err != nil {
		l.Error(err)
		return nil, err
	}
	if code != in.Code {
		return nil, errors.New("验证码错误")
	}
	// 不存在则插入一条数据
	userInfo, err := l.svcCtx.UserModel.FindOneByEmail(in.Email)
	if err != model.ErrNotFound && err != nil {
		l.Error(err)
		return nil, err
	} else {
		userInfo = &model.User{Email: in.Email, Status: 1, Gender: 1}
		result, err := l.svcCtx.UserModel.Insert(*userInfo)
		if err != nil {
			return nil, err
		}
		userInfo.ID, err = result.LastInsertId()
	}

	jwtToken, err := helper.GetJwtToken(l.svcCtx, userInfo.ID)
	if err != nil {
		l.Error(err)
		return nil, err
	}
	return &user.AuthByEmailCodeResp{AccessToken: jwtToken}, nil
}
