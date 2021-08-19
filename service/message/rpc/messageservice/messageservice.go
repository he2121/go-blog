// Code generated by goctl. DO NOT EDIT!
// Source: message.proto

//go:generate mockgen -destination ./messageservice_mock.go -package messageservice -source $GOFILE

package messageservice

import (
	"context"

	"github.com/he2121/go-blog/service/message/rpc/message"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	CreateNotificationReq   = message.CreateNotificationReq
	GetNotificationListResp = message.GetNotificationListResp
	UpdateNotificationResp  = message.UpdateNotificationResp
	CreateMessageResp       = message.CreateMessageResp
	Notification            = message.Notification
	CreateNotificationResp  = message.CreateNotificationResp
	GetNotificationListReq  = message.GetNotificationListReq
	GetMessageListResp      = message.GetMessageListResp
	UpdateMessageReq        = message.UpdateMessageReq
	UpdateMessageResp       = message.UpdateMessageResp
	UpdateNotificationReq   = message.UpdateNotificationReq
	Message                 = message.Message
	CreateMessageReq        = message.CreateMessageReq
	GetMessageListReq       = message.GetMessageListReq

	MessageService interface {
		//  系统通知的增删改查
		CreateNotification(ctx context.Context, in *CreateNotificationReq) (*CreateNotificationResp, error)
		GetNotificationList(ctx context.Context, in *GetNotificationListReq) (*GetNotificationListResp, error)
		UpdateNotification(ctx context.Context, in *UpdateNotificationReq) (*UpdateNotificationResp, error)
		//  消息的增删改查
		CreateMessage(ctx context.Context, in *CreateMessageReq) (*CreateMessageResp, error)
		GetMessageList(ctx context.Context, in *GetMessageListReq) (*GetMessageListResp, error)
		UpdateMessage(ctx context.Context, in *UpdateMessageReq) (*UpdateMessageResp, error)
	}

	defaultMessageService struct {
		cli zrpc.Client
	}
)

func NewMessageService(cli zrpc.Client) MessageService {
	return &defaultMessageService{
		cli: cli,
	}
}

//  系统通知的增删改查
func (m *defaultMessageService) CreateNotification(ctx context.Context, in *CreateNotificationReq) (*CreateNotificationResp, error) {
	client := message.NewMessageServiceClient(m.cli.Conn())
	return client.CreateNotification(ctx, in)
}

func (m *defaultMessageService) GetNotificationList(ctx context.Context, in *GetNotificationListReq) (*GetNotificationListResp, error) {
	client := message.NewMessageServiceClient(m.cli.Conn())
	return client.GetNotificationList(ctx, in)
}

func (m *defaultMessageService) UpdateNotification(ctx context.Context, in *UpdateNotificationReq) (*UpdateNotificationResp, error) {
	client := message.NewMessageServiceClient(m.cli.Conn())
	return client.UpdateNotification(ctx, in)
}

//  消息的增删改查
func (m *defaultMessageService) CreateMessage(ctx context.Context, in *CreateMessageReq) (*CreateMessageResp, error) {
	client := message.NewMessageServiceClient(m.cli.Conn())
	return client.CreateMessage(ctx, in)
}

func (m *defaultMessageService) GetMessageList(ctx context.Context, in *GetMessageListReq) (*GetMessageListResp, error) {
	client := message.NewMessageServiceClient(m.cli.Conn())
	return client.GetMessageList(ctx, in)
}

func (m *defaultMessageService) UpdateMessage(ctx context.Context, in *UpdateMessageReq) (*UpdateMessageResp, error) {
	client := message.NewMessageServiceClient(m.cli.Conn())
	return client.UpdateMessage(ctx, in)
}