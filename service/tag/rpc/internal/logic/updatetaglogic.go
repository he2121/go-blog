package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &tag.UpdateTagResp{}, nil
}
