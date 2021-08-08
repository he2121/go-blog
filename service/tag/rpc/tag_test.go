package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/core/conf"

	"github.com/he2121/go-blog/service/tag/rpc/internal/config"
	"github.com/he2121/go-blog/service/tag/rpc/internal/server"
	"github.com/he2121/go-blog/service/tag/rpc/internal/svc"
	"github.com/he2121/go-blog/service/tag/rpc/tag"
)

var srv *server.TagServiceServer

func init() {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv = server.NewTagServiceServer(ctx)
}

func TestCreateTag(t *testing.T) {
	resp, err := srv.CreateTag(context.TODO(), &tag.CreateTagReq{
		EntityType: tag.EntityType_EntityTypeBlog,
		EntityID:   1,
		Content:    "tag1",
	})
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestGetTagList(t *testing.T) {
	resp, err := srv.GetTagList(context.TODO(), &tag.GetTagListReq{
		IDs:        []int64{1},
		EntityIDs:  nil,
		EntityType: nil,
		Content:    nil,
		OrderBy:    nil,
		Offset:     0,
		Limit:      10,
		NeedCount:  nil,
	})
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestUpdateTag(t *testing.T) {
	a := int32(2)
	resp, err := srv.UpdateTag(context.TODO(), &tag.UpdateTagReq{
		ID:    1,
		Count: &a,
	})
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestViews(t *testing.T) {
	// userID 为 3 的用户查看了 博客 ID 1
	_, err := srv.UpdateBasicTag(context.TODO(), &tag.UpdateBasicTagReq{
		UserID:     3,
		ActionType: tag.ActionType_ActionTypeView,
		EntityType: tag.EntityType_EntityTypeBlog,
		EntityID:   1,
	})
	assert.Nil(t, err)

	resp, err := srv.MGetBasicTag(context.TODO(), &tag.MGetBasicTagReq{
		EntityType: tag.EntityType_EntityTypeBlog,
		EntityIDs:  []int64{1},
		ActionType: []tag.ActionType{tag.ActionType_ActionTypeView},
	})
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestLikes(t *testing.T) {
	_, err := srv.UpdateBasicTag(context.TODO(), &tag.UpdateBasicTagReq{
		UserID:     1,
		ActionType: tag.ActionType_ActionTypeLike, // tag.ActionType_ActionTypeUnLike
		EntityType: tag.EntityType_EntityTypeBlog,
		EntityID:   1,
	})
	assert.Nil(t, err)

	resp, err := srv.MGetBasicTag(context.TODO(), &tag.MGetBasicTagReq{
		EntityType: tag.EntityType_EntityTypeBlog,
		EntityIDs:  []int64{1},
		ActionType: []tag.ActionType{tag.ActionType_ActionTypeLike},
	})
	assert.Nil(t, err)
	fmt.Println(resp)
}

func TestFollowing(t *testing.T) {
	_, err := srv.UpdateBasicTag(context.TODO(), &tag.UpdateBasicTagReq{
		UserID:     2,
		ActionType: tag.ActionType_ActionTypeFollowing,
		EntityType: tag.EntityType_EntityTypeUser,
		EntityID:   1,
	})
	assert.Nil(t, err)

	resp, err := srv.MGetBasicTag(context.TODO(), &tag.MGetBasicTagReq{
		EntityType: tag.EntityType_EntityTypeUser,
		EntityIDs:  []int64{1},
		ActionType: []tag.ActionType{tag.ActionType_ActionTypeFollowing},
	})
	assert.Nil(t, err)
	fmt.Println(resp)
}
