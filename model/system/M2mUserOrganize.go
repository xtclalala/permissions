package system

import "github.com/google/uuid"

type M2mUserOrganize struct {
	SysOrganizeId int       `gorm:"column:sys_organize_id;index"`
	SysUserId     uuid.UUID `gorm:"column:sys_user_id;index"`
}
