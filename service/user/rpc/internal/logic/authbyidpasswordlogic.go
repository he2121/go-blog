package logic

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/tal-tech/go-zero/core/logx"

	"user/internal/helper"
	"user/internal/svc"
	"user/user"
)

type AuthByIDPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthByIDPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthByIDPasswordLogic {
	return &AuthByIDPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  ID + 密码认证
func (l *AuthByIDPasswordLogic) AuthByIDPassword(in *user.AuthByIDPasswordReq) (*user.AuthByIDPasswordResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(in.ID)
	if err != nil {
		return nil, err
	}
	if len(userInfo.Password) == 0 {
		return nil, errors.New("邮箱登陆没有设置密码无法使用密码登陆")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(in.Password)); err != nil {
		return nil, errors.New("密码错误")
	}
	token, err := helper.GetJwtToken(l.svcCtx, userInfo.ID)
	if err != nil {
		return nil, err
	}
	return &user.AuthByIDPasswordResp{AccessToken: token}, nil
}
