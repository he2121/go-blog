package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/core/conf"

	"github.com/he2121/go-blog/service/blog/rpc/blog"
	"github.com/he2121/go-blog/service/blog/rpc/internal/config"
	"github.com/he2121/go-blog/service/blog/rpc/internal/server"
	"github.com/he2121/go-blog/service/blog/rpc/internal/svc"
)

var srv *server.BlogServiceServer

func init() {
	//flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv = server.NewBlogServiceServer(ctx)
}

func TestCreateBlog(t *testing.T) {
	_, err := srv.CreateBlog(context.TODO(), &blog.CreateBlogReq{
		UserID:   1,
		Title:    "第一篇博客title",
		IsFolder: false,
		Content:  "第一篇博客content",
		Status:   1,
		FolderID: 0,
	})
	assert.Nil(t, err)
}

func TestCreateComment(t *testing.T) {
	_, err := srv.CreateComment(context.TODO(), &blog.CreateCommentReq{
		CommentType: 1,
		Content:     "一楼",
		ToID:        1,
		FromUserID:  1,
		ToUserID:    1,
	})
	assert.Nil(t, err)
}

func TestGetBlogList(t *testing.T) {
	a := int32(1)
	list, err := srv.GetBlogList(context.TODO(), &blog.GetBlogListReq{
		IDs:          []int64{1},
		UserIDs:      nil,
		Title:        nil,
		IsFolder:     nil,
		Status:       &a,
		FolderID:     nil,
		CreatedAtGTE: nil,
		Offset:       0,
		Limit:        10,
		NeedCount:    nil,
	})
	assert.Nil(t, err)
	fmt.Println(list)
}

func TestGetCommentList(t *testing.T) {
	a := int64(1)
	b := true
	resp, err := srv.GetCommentList(context.TODO(), &blog.GetCommentListReq{
		ToID:        &a,
		CommentType: nil,
		FromUserID:  nil,
		ToUserID:    nil,
		Offset:      0,
		Limit:       9,
		NeedCount:   &b,
	})
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestUpdateBlog(t *testing.T) {
	title := "修改后的标题"
	resp, err := srv.UpdateBlog(context.TODO(), &blog.UpdateBlogReq{
		ID:    1,
		Title: &title,
	})
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestUpdateComment(t *testing.T) {
	content := "修改后的 content"
	resp, err := srv.UpdateComment(context.TODO(), &blog.UpdateCommentReq{
		ID:      1,
		Content: &content,
		Status:  nil,
	})
	assert.Nil(t, err)
	fmt.Println(resp)
}
