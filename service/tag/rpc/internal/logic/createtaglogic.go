package logic

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/he2121/go-blog/service/tag/rpc/internal/svc"
	"github.com/he2121/go-blog/service/tag/rpc/model"
	"github.com/he2121/go-blog/service/tag/rpc/tag"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTagLogic {
	return &CreateTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTagLogic) CreateTag(in *tag.CreateTagReq) (*tag.CreateTagResp, error) {
	po := model.Tag{}
	if err := copier.Copy(&po, in); err != nil {
		return nil, err
	}
	if po.Count == 0 {
		po.Count = 1
	}
	if _, err := l.svcCtx.TagModel.Insert(po); err != nil {
		return nil, err
	}
	return &tag.CreateTagResp{}, nil
}
