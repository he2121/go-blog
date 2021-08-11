package logic

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/he2121/go-blog/service/message/rpc/internal/svc"
	"github.com/he2121/go-blog/service/message/rpc/message"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateNotificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNotificationLogic {
	return &UpdateNotificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateNotificationLogic) UpdateNotification(in *message.UpdateNotificationReq) (*message.UpdateNotificationResp, error) {
	po, err := l.svcCtx.NotificationModel.FindOne(in.ID)
	if err != nil {
		return nil, err
	}
	if err := copier.CopyWithOption(po, in, copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}
	if err := l.svcCtx.NotificationModel.Update(*po); err != nil {
		return nil, err
	}
	return &message.UpdateNotificationResp{}, nil
}
