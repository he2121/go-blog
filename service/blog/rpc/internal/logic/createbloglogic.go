package logic

import (
	"context"

	"blog/blog"
	"blog/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateBlogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBlogLogic {
	return &CreateBlogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBlogLogic) CreateBlog(in *blog.CreateBlogReq) (*blog.CreateBlogResp, error) {
	// todo: add your logic here and delete this line

	return &blog.CreateBlogResp{}, nil
}
