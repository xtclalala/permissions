package request

import (
	"permissions/model/common"
	"permissions/model/system"
)

type SearchMenu struct {
	common.BasePage
	system.SysMenu
}
