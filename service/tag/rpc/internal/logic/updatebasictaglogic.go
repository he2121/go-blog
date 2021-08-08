package logic

import (
	"context"
	"fmt"

	"github.com/tal-tech/go-zero/core/stores/redis"

	"github.com/he2121/go-blog/service/tag/rpc/internal/svc"
	"github.com/he2121/go-blog/service/tag/rpc/tag"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateBasicTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBasicTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBasicTagLogic {
	return &UpdateBasicTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  基本 tag 的修改与查询
func (l *UpdateBasicTagLogic) UpdateBasicTag(in *tag.UpdateBasicTagReq) (*tag.UpdateBasicTagResp, error) {
	if err := actionTypeFunc[in.ActionType](l.svcCtx, in); err != nil {
		return nil, err
	}
	return &tag.UpdateBasicTagResp{}, nil
}

type actionFunc func(svc *svc.ServiceContext, in *tag.UpdateBasicTagReq) error

var incrActionFunc actionFunc = func(svc *svc.ServiceContext, in *tag.UpdateBasicTagReq) error {
	key := generatorRedisKey(ActionType2RedisPrefix[in.ActionType], in.EntityType, in.EntityID)
	if _, err := svc.Redis.Incr(key); err != nil {
		return err
	}
	return nil
}

var addSetActionFunc actionFunc = func(svc *svc.ServiceContext, in *tag.UpdateBasicTagReq) error {
	key := generatorRedisKey(ActionType2RedisPrefix[in.ActionType], in.EntityType, in.EntityID)
	if _, err := svc.Redis.Sadd(key, in.UserID); err != nil {
		return err
	}
	return nil
}

var removeSetActionFunc actionFunc = func(svc *svc.ServiceContext, in *tag.UpdateBasicTagReq) error {
	key := generatorRedisKey(ActionType2RedisPrefix[in.ActionType], in.EntityType, in.EntityID)
	if _, err := svc.Redis.Srem(key, in.UserID); err != nil {
		return err
	}
	return nil
}

var actionTypeFunc = map[tag.ActionType]func(svc *svc.ServiceContext, in *tag.UpdateBasicTagReq) error{
	tag.ActionType_ActionTypeView:   incrActionFunc,
	tag.ActionType_ActionTypeLike:   addSetActionFunc,
	tag.ActionType_ActionTypeUnLike: removeSetActionFunc,
	tag.ActionType_ActionTypeHate:   addSetActionFunc,
	tag.ActionType_ActionTypeUnHate: removeSetActionFunc,
	tag.ActionType_ActionTypeFollowing: func(svc *svc.ServiceContext, in *tag.UpdateBasicTagReq) error {
		err := svc.Redis.Pipelined(func(pipeliner redis.Pipeliner) error {
			// 先记 followers
			key := generatorRedisKey(ActionType2RedisPrefix[in.ActionType], in.EntityType, in.EntityID)
			pipeliner.SAdd(key, in.UserID)
			// 在记 following
			key = generatorRedisKey("following", in.EntityType, in.UserID)
			pipeliner.SAdd(key, in.EntityID)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	},
	tag.ActionType_ActionTypeUnFollowing: func(svc *svc.ServiceContext, in *tag.UpdateBasicTagReq) error {
		err := svc.Redis.Pipelined(func(pipeliner redis.Pipeliner) error {
			// 先记 followers
			key := generatorRedisKey(ActionType2RedisPrefix[in.ActionType], in.EntityType, in.EntityID)
			pipeliner.SRem(key, in.UserID)
			// 在记 following
			key = generatorRedisKey("following", in.EntityType, in.UserID)
			pipeliner.SRem(key, in.EntityID)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	},
}

func generatorRedisKey(prefix string, entityType tag.EntityType, id int64) string {
	return fmt.Sprintf("%s:%s:%d", prefix, EntityType2RedisKey[entityType], id)
}

var ActionType2RedisPrefix = map[tag.ActionType]string{
	tag.ActionType_ActionTypeView:        "views",
	tag.ActionType_ActionTypeLike:        "like",
	tag.ActionType_ActionTypeUnLike:      "like",
	tag.ActionType_ActionTypeHate:        "hate",
	tag.ActionType_ActionTypeUnHate:      "hate",
	tag.ActionType_ActionTypeFollowing:   "followers",
	tag.ActionType_ActionTypeUnFollowing: "followers",
}

var EntityType2RedisKey = map[tag.EntityType]string{
	tag.EntityType_EntityTypeBlog:    "blog",
	tag.EntityType_EntityTypeComment: "comment",
	tag.EntityType_EntityTypeUser:    "user",
}
