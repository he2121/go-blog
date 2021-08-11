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

type GetMessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageListLogic {
	return &GetMessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageListLogic) GetMessageList(in *message.GetMessageListReq) (*message.GetMessageListResp, error) {
	if in.Limit == 0 {
		in.Limit = 999
	}
	whereMessage := &model.WhereMessage{}
	if err := copier.Copy(whereMessage, in); err != nil {
		return nil, err
	}
	option := sql_helper.Option{Offset: int(in.Offset), Limit: int(in.Limit + 1)}
	pos, err := l.svcCtx.MessageModel.GetMessageList(*whereMessage, &option)
	if err != nil {
		return nil, err
	}
	res := &message.GetMessageListResp{}
	if in.NeedCount != nil && *in.NeedCount {
		res.TotalCount, err = l.svcCtx.MessageModel.Count(*whereMessage)
		if err != nil {
			return nil, err
		}
	}
	if len(pos) > int(in.Limit) {
		pos = pos[0:in.Limit]
		res.HasMore = true
		return nil, err
	}
	res.IDs, res.Messages = pack.MakeMessageDtos(pos)
	return res, nil
}
