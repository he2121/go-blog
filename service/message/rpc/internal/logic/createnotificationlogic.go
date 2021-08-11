package logic

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/he2121/go-blog/service/message/rpc/internal/svc"
	"github.com/he2121/go-blog/service/message/rpc/message"
	"github.com/he2121/go-blog/service/message/rpc/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateNotificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNotificationLogic {
	return &CreateNotificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  系统通知的增删改查
func (l *CreateNotificationLogic) CreateNotification(in *message.CreateNotificationReq) (*message.CreateNotificationResp, error) {
	notificationInfo := &model.Notification{}
	if err := copier.Copy(notificationInfo, in); err != nil {
		return nil, err
	}
	if _, err := l.svcCtx.NotificationModel.Insert(*notificationInfo); err != nil {
		return nil, err
	}
	return &message.CreateNotificationResp{}, nil
}
