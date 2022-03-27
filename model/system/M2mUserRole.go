package system

import "github.com/google/uuid"

type M2mUserRole struct {
	SysRoleId int       `gorm:"column:sys_role_id"`
	SysUserId uuid.UUID `gorm:"column:sys_user_id"`
}
