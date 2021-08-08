package logic

import (
	"context"

	sql_helper "github.com/he2121/go-blog/common/sql-helper"
	"github.com/jinzhu/copier"

	"github.com/he2121/go-blog/service/tag/rpc/internal/pack"
	"github.com/he2121/go-blog/service/tag/rpc/internal/svc"
	"github.com/he2121/go-blog/service/tag/rpc/model"
	"github.com/he2121/go-blog/service/tag/rpc/tag"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetTagListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagListLogic {
	return &GetTagListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  自定义TAG 增删改查接口
func (l *GetTagListLogic) GetTagList(in *tag.GetTagListReq) (*tag.GetTagListResp, error) {
	whereTag := model.WhereTag{}
	if err := copier.CopyWithOption(&whereTag, in, copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}
	option := sql_helper.Option{
		Offset:  int(in.Offset),
		Limit:   int(in.Limit),
		OrderBy: in.OrderBy,
	}
	pos, err := l.svcCtx.TagModel.GetTagList(whereTag, &option)
	if err != nil {
		return nil, err
	}
	res := &tag.GetTagListResp{}

	if in.NeedCount != nil && *in.NeedCount {
		res.TotalCount, err = l.svcCtx.TagModel.Count(whereTag)
		if err != nil {
			return nil, err
		}
	}
	if len(pos) > int(in.Limit) {
		pos = pos[0:in.Limit]
		res.HasMore = true
		return nil, err
	}
	res.IDs, res.Tags = pack.MakeTagDtos(pos)
	return res, nil
}
