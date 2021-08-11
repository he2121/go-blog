package pack

import (
	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"

	"github.com/he2121/go-blog/service/message/rpc/message"
	"github.com/he2121/go-blog/service/message/rpc/model"
)

func MakeMessageDtos(pos []*model.Message) (ids []int64, dtos map[int64]*message.Message) {
	dtos = make(map[int64]*message.Message)
	for _, po := range pos {
		ids = append(ids, po.ID)
		dto := MakeMessageDto(po)
		dtos[po.ID] = dto
	}
	return ids, dtos
}

func MakeMessageDto(po *model.Message) *message.Message {
	dto := &message.Message{}
	if err := copier.Copy(dto, po); err != nil {
		logx.Error(err, "MakeMessageDto err")
	}
	dto.CreatedAt = po.CreatedAt.Unix()
	dto.UpdatedAt = po.UpdatedAt.Unix()
	return dto
}
