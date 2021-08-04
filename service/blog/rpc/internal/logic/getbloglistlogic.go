package logic

import (
	"context"

	sql_helper "github.com/he2121/go-blog/common/sql-helper"
	"github.com/jinzhu/copier"

	"github.com/tal-tech/go-zero/core/logx"

	"github.com/he2121/go-blog/service/blog/rpc/blog"
	"github.com/he2121/go-blog/service/blog/rpc/internal/pack"
	"github.com/he2121/go-blog/service/blog/rpc/internal/svc"
	"github.com/he2121/go-blog/service/blog/rpc/model"
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
	if in.Limit == 0 {
		in.Limit = 999
	}
	whereBlog := &model.WhereBlog{}
	if err := copier.Copy(whereBlog, in); err != nil {
		return nil, err
	}
	option := sql_helper.Option{Offset: int(in.Offset), Limit: int(in.Limit + 1)}
	pos, err := l.svcCtx.BlogModel.GetBlogList(*whereBlog, &option)
	if err != nil {
		return nil, err
	}
	res := &blog.GetBlogListResp{}

	if in.NeedCount != nil && *in.NeedCount {
		res.TotalCount, err = l.svcCtx.BlogModel.Count(*whereBlog)
		if err != nil {
			return nil, err
		}
	}
	if len(pos) > int(in.Limit) {
		pos = pos[0:in.Limit]
		res.HasMore = true
		return nil, err
	}
	res.IDs, res.Blogs = pack.MakeBlogDtos(pos)
	return res, nil
}
