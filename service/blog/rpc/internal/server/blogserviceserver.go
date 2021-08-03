// Code generated by goctl. DO NOT EDIT!
// Source: blog.proto

package server

import (
	"context"

	"blog/blog"
	"blog/internal/logic"
	"blog/internal/svc"
)

type BlogServiceServer struct {
	svcCtx *svc.ServiceContext
}

func NewBlogServiceServer(svcCtx *svc.ServiceContext) *BlogServiceServer {
	return &BlogServiceServer{
		svcCtx: svcCtx,
	}
}

//  blog 增删改查接口
func (s *BlogServiceServer) GetBlogList(ctx context.Context, in *blog.GetBlogListReq) (*blog.GetBlogListResp, error) {
	l := logic.NewGetBlogListLogic(ctx, s.svcCtx)
	return l.GetBlogList(in)
}

func (s *BlogServiceServer) CreateBlog(ctx context.Context, in *blog.CreateBlogReq) (*blog.CreateBlogResp, error) {
	l := logic.NewCreateBlogLogic(ctx, s.svcCtx)
	return l.CreateBlog(in)
}

func (s *BlogServiceServer) UpdateBlog(ctx context.Context, in *blog.UpdateBlogReq) (*blog.UpdateBlogResp, error) {
	l := logic.NewUpdateBlogLogic(ctx, s.svcCtx)
	return l.UpdateBlog(in)
}

//  comment 增删改查接口
func (s *BlogServiceServer) GetCommentList(ctx context.Context, in *blog.GetCommentListReq) (*blog.GetCommentListResp, error) {
	l := logic.NewGetCommentListLogic(ctx, s.svcCtx)
	return l.GetCommentList(in)
}

func (s *BlogServiceServer) UpdateComment(ctx context.Context, in *blog.UpdateCommentReq) (*blog.UpdateCommentResp, error) {
	l := logic.NewUpdateCommentLogic(ctx, s.svcCtx)
	return l.UpdateComment(in)
}

func (s *BlogServiceServer) CreateComment(ctx context.Context, in *blog.CreateCommentReq) (*blog.CreateCommentResp, error) {
	l := logic.NewCreateCommentLogic(ctx, s.svcCtx)
	return l.CreateComment(in)
}
