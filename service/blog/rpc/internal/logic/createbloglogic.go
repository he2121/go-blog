package logic

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/tal-tech/go-zero/core/logx"

	"github.com/he2121/go-blog/service/blog/rpc/blog"
	"github.com/he2121/go-blog/service/blog/rpc/internal/svc"
	"github.com/he2121/go-blog/service/blog/rpc/model"
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
	blogInfo := &model.Blog{}
	if err := copier.Copy(blogInfo, in); err != nil {
		return nil, err
	}
	if _, err := l.svcCtx.BlogModel.Insert(*blogInfo); err != nil {
		return nil, err
	}
	// to do 通知关注此博客的人
	return &blog.CreateBlogResp{}, nil
}
