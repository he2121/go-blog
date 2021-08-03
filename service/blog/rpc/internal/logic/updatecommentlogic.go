package logic

import (
	"context"

	"blog/blog"
	"blog/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentLogic) UpdateComment(in *blog.UpdateCommentReq) (*blog.UpdateCommentResp, error) {
	// todo: add your logic here and delete this line

	return &blog.UpdateCommentResp{}, nil
}
