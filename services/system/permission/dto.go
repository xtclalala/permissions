package permission

import (
	"permissions/services/common/page/dto"
)

type SearchPermission struct {
	dto.BasePage
	Permission
}

type Permission struct {
	Name      string `json:"name"`
	SysMenuId uint   `json:"menuId"`
}
