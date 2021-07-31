package pack

import (
	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tidwall/gjson"

	"github.com/he2121/go-blog/service/user/rpc/model"
	"github.com/he2121/go-blog/service/user/rpc/user"
)

func MakeUserDtos(pos []*model.User) (ids []int64, dtos map[int64]*user.User) {
	dtos = make(map[int64]*user.User)
	for _, po := range pos {
		ids = append(ids, po.ID)
		dto := MakeUserDto(po)
		dtos[po.ID] = dto
	}
	return ids, dtos
}

func MakeUserDto(po *model.User) *user.User {
	dto := &user.User{}
	err := copier.Copy(dto, po)
	logx.Error(err, "MakeUserDto err")
	dto.Extra = &user.UserExtra{
		AvatarUrl: gjson.Get(po.Extra, "avatar_url").String(),
	}
	dto.BirthDate = po.BirthDate.Unix()
	dto.CreatedAt = po.CreatedAt.Unix()
	dto.UpdatedAt = po.UpdatedAt.Unix()
	return dto
}
