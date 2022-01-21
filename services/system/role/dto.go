package role

import "permissions/services/common/page/dto"

type SearchRole struct {
	dto.BasePage
	Role
}

type Role struct {
	Name string `json:"name"`
	Pid  uint   `json:"pid"`
}
