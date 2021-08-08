package pack

import (
	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"

	"github.com/he2121/go-blog/service/tag/rpc/model"

	"github.com/he2121/go-blog/service/tag/rpc/tag"
)

func MakeTagDtos(pos []*model.Tag) (ids []int64, dtos map[int64]*tag.Tag) {
	dtos = make(map[int64]*tag.Tag)
	for _, po := range pos {
		ids = append(ids, po.ID)
		dto := MakeTagDto(po)
		dtos[po.ID] = dto
	}
	return ids, dtos
}

func MakeTagDto(po *model.Tag) *tag.Tag {
	dto := &tag.Tag{}
	err := copier.Copy(dto, po)
	logx.Error(err, "MakeTagDto err")
	dto.CreatedAt = po.CreatedAt.Unix()
	dto.UpdatedAt = po.UpdatedAt.Unix()
	return dto
}
