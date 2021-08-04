package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"

	"github.com/he2121/go-blog/service/blog/rpc/blog"
	"github.com/he2121/go-blog/service/blog/rpc/internal/svc"
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
	po, err := l.svcCtx.BlogModel.FindOne(in.ID)
	if err != nil {
		return nil, err
	}
	if err := copier.CopyWithOption(po, in, copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}
	if in.Extra != nil {
		// to update something
	}
	if err := l.svcCtx.BlogModel.Update(*po); err != nil {
		return nil, err
	}
	return &blog.UpdateBlogResp{}, nil
}
