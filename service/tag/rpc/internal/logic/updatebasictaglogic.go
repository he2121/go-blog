package logic

import (
	"context"

	"github.com/he2121/go-blog/service/tag/rpc/internal/svc"
	"github.com/he2121/go-blog/service/tag/rpc/tag"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateBasicTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBasicTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBasicTagLogic {
	return &UpdateBasicTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  基本 tag 的修改与查询
func (l *UpdateBasicTagLogic) UpdateBasicTag(in *tag.UpdateBasicTagReq) (*tag.UpdateBasicTagResp, error) {
	// todo: add your logic here and delete this line

	return &tag.UpdateBasicTagResp{}, nil
}
