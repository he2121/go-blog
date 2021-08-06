package logic

import (
	"context"

	"github.com/he2121/go-blog/service/tag/rpc/internal/svc"
	"github.com/he2121/go-blog/service/tag/rpc/tag"

	"github.com/tal-tech/go-zero/core/logx"
)

type MGetBasicTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMGetBasicTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MGetBasicTagLogic {
	return &MGetBasicTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MGetBasicTagLogic) MGetBasicTag(in *tag.MGetBasicTagReq) (*tag.MGetBasicTagResp, error) {
	// todo: add your logic here and delete this line

	return &tag.MGetBasicTagResp{}, nil
}
