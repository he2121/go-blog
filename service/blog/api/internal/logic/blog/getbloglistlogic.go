package logic

import (
	"context"

	"github.com/he2121/go-blog/service/blog/api/internal/svc"
	"github.com/he2121/go-blog/service/blog/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBlogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBlogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBlogListLogic {
	return GetBlogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBlogListLogic) GetBlogList(req types.GetBlogListReq) (*types.GetBlogListResp, error) {
	// todo: add your logic here and delete this line

	return &types.GetBlogListResp{}, nil
}
