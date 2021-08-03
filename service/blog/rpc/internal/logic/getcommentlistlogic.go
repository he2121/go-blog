package logic

import (
	"context"

	"blog/blog"
	"blog/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  comment 增删改查接口
func (l *GetCommentListLogic) GetCommentList(in *blog.GetCommentListReq) (*blog.GetCommentListResp, error) {
	// todo: add your logic here and delete this line

	return &blog.GetCommentListResp{}, nil
}
