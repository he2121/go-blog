package logic

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/he2121/go-blog/service/message/rpc/internal/svc"
	"github.com/he2121/go-blog/service/message/rpc/message"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMessageLogic {
	return &UpdateMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMessageLogic) UpdateMessage(in *message.UpdateMessageReq) (*message.UpdateMessageResp, error) {
	po, err := l.svcCtx.MessageModel.FindOne(in.ID)
	if err != nil {
		return nil, err
	}
	if err := copier.CopyWithOption(po, in, copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}
	if err := l.svcCtx.MessageModel.Update(*po); err != nil {
		return nil, err
	}
	return &message.UpdateMessageResp{}, nil
}
