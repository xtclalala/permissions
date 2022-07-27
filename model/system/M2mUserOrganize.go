package system

import "github.com/google/uuid"

type M2mUserOrganize struct {
	SysOrganizeId int       `gorm:"column:sys_organize_id;index"`
	SysUserId     uuid.UUID `gorm:"column:sys_user_id;index"`
}

func (m M2mUserOrganize) TableName() string {
	return "sys_m2m_user_organize"
}
