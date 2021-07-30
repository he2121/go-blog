package logic

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/he2121/go-blog/user/rpc/internal/helper"
	"github.com/he2121/go-blog/user/rpc/internal/svc"
	"github.com/he2121/go-blog/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type SendEmailCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailCodeLogic {
	return &SendEmailCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 邮箱登陆认证发送验证吗
func (l *SendEmailCodeLogic) SendEmailCode(in *user.SendEmailCodeReq) (*user.SendEmailCodeResp, error) {
	// 随机生成验证码并发送
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", r.Int31n(1000000))
	if err := helper.ServerEmail.SendCode(in.Email, code); err != nil {
		l.Error(err)
		return nil, err
	}
	// 存入验证码到redis
	if err := l.svcCtx.Redis.Setex(fmt.Sprintf("email:code:%s", in.Email), code, 20*60); err != nil {
		l.Error(err)
		return nil, err
	}
	return &user.SendEmailCodeResp{}, nil
}
