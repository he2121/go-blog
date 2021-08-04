package logic

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/tal-tech/go-zero/core/logx"

	"github.com/he2121/go-blog/service/blog/rpc/blog"
	"github.com/he2121/go-blog/service/blog/rpc/internal/svc"
	"github.com/he2121/go-blog/service/blog/rpc/model"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentLogic) CreateComment(in *blog.CreateCommentReq) (*blog.CreateCommentResp, error) {
	commentInfo := &model.Comment{}
	if err := copier.Copy(commentInfo, in); err != nil {
		return nil, err
	}
	commentInfo.Status = 1
	if _, err := l.svcCtx.CommentModel.Insert(*commentInfo); err != nil {
		return nil, err
	}
	// to do 通知相应的用户
	return &blog.CreateCommentResp{}, nil
}
