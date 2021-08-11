package pack

import (
	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"

	"github.com/he2121/go-blog/service/message/rpc/message"
	"github.com/he2121/go-blog/service/message/rpc/model"
)

func MakeNotificationDtos(pos []*model.Notification) (ids []int64, dtos map[int64]*message.Notification) {
	dtos = make(map[int64]*message.Notification)
	for _, po := range pos {
		ids = append(ids, po.ID)
		dto := MakeNotificationDto(po)
		dtos[po.ID] = dto
	}
	return ids, dtos
}

func MakeNotificationDto(po *model.Notification) *message.Notification {
	dto := &message.Notification{}
	if err := copier.Copy(dto, po); err != nil {
		logx.Error(err, "MakeNotificationDto err")
	}
	dto.CreatedAt = po.CreatedAt.Unix()
	return dto
}
