package logic

import (
	"context"

	sql_helper "github.com/he2121/go-blog/common/sql-helper"
	"github.com/jinzhu/copier"

	"github.com/he2121/go-blog/service/message/rpc/internal/pack"
	"github.com/he2121/go-blog/service/message/rpc/internal/svc"
	"github.com/he2121/go-blog/service/message/rpc/message"
	"github.com/he2121/go-blog/service/message/rpc/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetNotificationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNotificationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNotificationListLogic {
	return &GetNotificationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNotificationListLogic) GetNotificationList(in *message.GetNotificationListReq) (*message.GetNotificationListResp, error) {
	if in.Limit == 0 {
		in.Limit = 999
	}
	whereNotification := &model.WhereNotification{}
	if err := copier.Copy(whereNotification, in); err != nil {
		return nil, err
	}
	option := sql_helper.Option{Offset: int(in.Offset), Limit: int(in.Limit + 1)}
	pos, err := l.svcCtx.NotificationModel.GetNotificationList(*whereNotification, &option)
	if err != nil {
		return nil, err
	}
	res := &message.GetNotificationListResp{}
	if in.NeedCount != nil && *in.NeedCount {
		res.TotalCount, err = l.svcCtx.NotificationModel.Count(*whereNotification)
		if err != nil {
			return nil, err
		}
	}
	if len(pos) > int(in.Limit) {
		pos = pos[0:in.Limit]
		res.HasMore = true
		return nil, err
	}
	res.IDs, res.Notifications = pack.MakeNotificationDtos(pos)
	return res, nil
}
