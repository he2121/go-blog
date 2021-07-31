package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"

	"user/internal/pack"
	"user/internal/svc"
	"user/user"
)

type MGetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MGetUserInfoLogic {
	return &MGetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  批量获取指定用户ID 的基本信息
func (l *MGetUserInfoLogic) MGetUserInfo(in *user.MGetUserInfoReq) (*user.MGetUserInfoResp, error) {
	users, err := l.svcCtx.UserModel.MGetUser(in.IDs)
	if err != nil {
		return nil, err
	}
	resp := &user.MGetUserInfoResp{}
	resp.IDs, resp.Users = pack.MakeUserDtos(users)
	return resp, nil
}
