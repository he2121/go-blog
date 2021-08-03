package logic

import (
	"context"

	"blog/blog"
	"blog/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBlogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBlogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlogListLogic {
	return &GetBlogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  blog 增删改查接口
func (l *GetBlogListLogic) GetBlogList(in *blog.GetBlogListReq) (*blog.GetBlogListResp, error) {
	// todo: add your logic here and delete this line
	return &blog.GetBlogListResp{}, nil
}
