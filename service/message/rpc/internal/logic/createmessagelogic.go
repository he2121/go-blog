package logic

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/he2121/go-blog/service/message/rpc/internal/svc"
	"github.com/he2121/go-blog/service/message/rpc/message"
	"github.com/he2121/go-blog/service/message/rpc/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMessageLogic {
	return &CreateMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  消息的增删改查
func (l *CreateMessageLogic) CreateMessage(in *message.CreateMessageReq) (*message.CreateMessageResp, error) {
	messageInfo := &model.Message{}
	if err := copier.Copy(messageInfo, in); err != nil {
		return nil, err
	}
	sessionID := fmt.Sprintf("%d:%d", in.FromID, in.ToID)
	if in.FromID > in.ToID {
		sessionID = fmt.Sprintf("%d:%d", in.ToID, in.FromID)
	}
	messageInfo.SessionID = sessionID

	if _, err := l.svcCtx.MessageModel.Insert(*messageInfo); err != nil {
		return nil, err
	}
	return &message.CreateMessageResp{}, nil
}
