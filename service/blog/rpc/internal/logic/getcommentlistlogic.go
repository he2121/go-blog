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
	if in.Limit == 0 {
		in.Limit = 999
	}
	whereComment := &model.WhereComment{}
	if err := copier.Copy(whereComment, in); err != nil {
		return nil, err
	}
	option := sql_helper.Option{Offset: int(in.Offset), Limit: int(in.Limit + 1)}
	pos, err := l.svcCtx.CommentModel.GetCommentList(*whereComment, &option)
	if err != nil {
		return nil, err
	}
	res := &blog.GetCommentListResp{}
	if in.NeedCount != nil && *in.NeedCount {
		res.TotalCount, err = l.svcCtx.CommentModel.Count(*whereComment)
		if err != nil {
			return nil, err
		}
	}
	if len(pos) > int(in.Limit) {
		pos = pos[0:in.Limit]
		res.HasMore = true
		return nil, err
	}
	res.IDs, res.Comments = pack.MakeCommentDtos(pos)
	return res, nil
}
