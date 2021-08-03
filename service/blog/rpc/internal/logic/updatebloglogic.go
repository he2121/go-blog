package logic

import (
	"context"

	"blog/blog"
	"blog/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateBlogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBlogLogic {
	return &UpdateBlogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBlogLogic) UpdateBlog(in *blog.UpdateBlogReq) (*blog.UpdateBlogResp, error) {
	// todo: add your logic here and delete this line

	return &blog.UpdateBlogResp{}, nil
}
