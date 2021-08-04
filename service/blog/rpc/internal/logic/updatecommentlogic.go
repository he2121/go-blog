package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"

	"github.com/he2121/go-blog/service/blog/rpc/blog"
	"github.com/he2121/go-blog/service/blog/rpc/internal/svc"
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
	po, err := l.svcCtx.CommentModel.FindOne(in.ID)
	if err != nil {
		return nil, err
	}
	if err := copier.CopyWithOption(po, in, copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}
	if err := l.svcCtx.CommentModel.Update(*po); err != nil {
		return nil, err
	}
	return &blog.UpdateCommentResp{}, nil
}
