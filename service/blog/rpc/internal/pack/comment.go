package pack

import (
	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"

	"github.com/he2121/go-blog/service/blog/rpc/blog"
	"github.com/he2121/go-blog/service/blog/rpc/model"
)

func MakeCommentDtos(pos []*model.Comment) (ids []int64, dtos map[int64]*blog.Comment) {
	dtos = make(map[int64]*blog.Comment)
	for _, po := range pos {
		ids = append(ids, po.ID)
		dto := MakeCommentDto(po)
		dtos[po.ID] = dto
	}
	return ids, dtos
}

func MakeCommentDto(po *model.Comment) *blog.Comment {
	dto := &blog.Comment{}
	err := copier.Copy(dto, po)
	logx.Error(err, "MakeCommentDto err")
	dto.Extra = &blog.CommentExtra{}
	dto.CreatedAt = po.CreatedAt.Unix()
	return dto
}
