package logic

import (
	"context"
	"errors"
	"strconv"

	"github.com/he2121/go-blog/service/tag/rpc/internal/svc"
	"github.com/he2121/go-blog/service/tag/rpc/tag"

	"github.com/tal-tech/go-zero/core/logx"
)

type MGetBasicTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMGetBasicTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MGetBasicTagLogic {
	return &MGetBasicTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MGetBasicTagLogic) MGetBasicTag(in *tag.MGetBasicTagReq) (*tag.MGetBasicTagResp, error) {
	if len(in.EntityIDs) == 0 || len(in.ActionType) == 0 {
		return nil, errors.New("参数错误")
	}

	resp := &tag.MGetBasicTagResp{BasicTags: map[int64]*tag.BasicTag{}}
	for _, entityID := range in.EntityIDs {
		basicTag := tag.BasicTag{}
		for _, actionType := range in.ActionType {
			key := generatorRedisKey(ActionType2RedisPrefix[actionType], in.EntityType, entityID)
			// 浏览量单独处理
			if actionType == tag.ActionType_ActionTypeView {
				val, err := l.svcCtx.Redis.Get(key)
				if err != nil {
					return nil, err
				}
				views, err := strconv.ParseInt(val, 10, 64)
				if err != nil {
					return nil, err
				}
				basicTag.Views = views
				continue
			}
			// 其余的使用 redis set 存储
			vals, err := l.svcCtx.Redis.Smembers(key)
			if err != nil {
				return nil, err
			}
			ids, err := stringList2Int64List(vals)
			if err != nil {
				return nil, err
			}
			if actionType == tag.ActionType_ActionTypeLike || actionType == tag.ActionType_ActionTypeUnLike {
				basicTag.LikeIDs = ids
			}
			if actionType == tag.ActionType_ActionTypeHate || actionType == tag.ActionType_ActionTypeUnHate {
				basicTag.HateIDs = ids
			}
			if actionType == tag.ActionType_ActionTypeFollowing || actionType == tag.ActionType_ActionTypeUnFollowing {
				basicTag.FollowerIDs = ids
				// 如果查询对象是 user， 则把 following 也返回
				if in.EntityType == tag.EntityType_EntityTypeUser {
					redisKey := generatorRedisKey("following", in.EntityType, entityID)
					vals, err := l.svcCtx.Redis.Smembers(redisKey)
					if err != nil {
						return nil, err
					}
					ids, err := stringList2Int64List(vals)
					if err != nil {
						return nil, err
					}
					basicTag.FollowingIDs = ids
				}
			}
		}
		resp.EntityIDs = append(resp.EntityIDs, entityID)
		resp.BasicTags[entityID] = &basicTag
	}

	return resp, nil
}

func stringList2Int64List(strs []string) (ids []int64, err error) {
	for _, str := range strs {
		parseInt, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, parseInt)
	}
	return ids, nil
}
