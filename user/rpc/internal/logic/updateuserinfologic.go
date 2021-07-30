package logic

import (
	"context"
	"time"

	"github.com/jinzhu/copier"
	"github.com/tidwall/sjson"
	"golang.org/x/crypto/bcrypt"

	"github.com/he2121/go-blog/user/rpc/internal/svc"
	"github.com/he2121/go-blog/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改指定用户ID 的基本信息
func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user.UpdateUserInfoReq) (*user.UpdateUserInfoResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(in.ID)
	if err != nil {
		return nil, err
	}
	// 先 copy 一下，有些需要特殊处理
	if err = copier.CopyWithOption(userInfo, in, copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}
	// 这里用merge json 好点， 偷个懒
	if in.UserExtra != nil && len(in.UserExtra.AvatarUrl) != 0 {
		extra, err := sjson.Set(userInfo.Extra, "avatar_url", in.UserExtra.AvatarUrl)
		if err != nil {
			return nil, err
		}
		userInfo.Extra = extra
	}
	if in.BirthDate != 0 {
		userInfo.BirthDate = time.Unix(in.BirthDate, 0)
	}
	if len(in.PassWord) != 0 {
		password, err := bcrypt.GenerateFromPassword([]byte(in.PassWord), 10)
		if err != nil {
			return nil, err
		}
		userInfo.Password = string(password)
	}
	// 这里的逻辑很恶心，待测试zero值，是否更新，go-zero 的 dal层不咋友好
	if err := l.svcCtx.UserModel.Update(*userInfo); err != nil {
		return nil, err
	}
	return &user.UpdateUserInfoResp{}, nil
}
