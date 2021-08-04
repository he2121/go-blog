package pack

import (
	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"

	"github.com/he2121/go-blog/service/blog/rpc/blog"
	"github.com/he2121/go-blog/service/blog/rpc/model"
)

func MakeBlogDtos(pos []*model.Blog) (ids []int64, dtos map[int64]*blog.Blog) {
	dtos = make(map[int64]*blog.Blog)
	for _, po := range pos {
		ids = append(ids, po.ID)
		dto := MakeBlogDto(po)
		dtos[po.ID] = dto
	}
	return ids, dtos
}

func MakeBlogDto(po *model.Blog) *blog.Blog {
	dto := &blog.Blog{}
	err := copier.Copy(dto, po)
	logx.Error(err, "MakeBlogDto err")
	dto.Extra = &blog.BlogExtra{}
	dto.CreatedAt = po.CreatedAt.Unix()
	dto.UpdatedAt = po.UpdatedAt.Unix()
	return dto
}
