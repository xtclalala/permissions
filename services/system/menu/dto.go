package menu

import (
	"permissions/services/common/page/dto"
)

type SearchMenu struct {
	dto.BasePage
	Menu
}

type Menu struct {
	Pid       uint   `json:"pid"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
}
