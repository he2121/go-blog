package logic

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/he2121/go-blog/service/tag/rpc/internal/svc"
	"github.com/he2121/go-blog/service/tag/rpc/tag"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTagLogic {
	return &UpdateTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTagLogic) UpdateTag(in *tag.UpdateTagReq) (*tag.UpdateTagResp, error) {
	po, err := l.svcCtx.TagModel.FindOne(in.ID)
	if err != nil {
		return nil, err
	}
	if err := copier.CopyWithOption(po, in, copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}
	if err := l.svcCtx.TagModel.Update(*po); err != nil {
		return nil, err
	}
	return &tag.UpdateTagResp{}, nil
}
