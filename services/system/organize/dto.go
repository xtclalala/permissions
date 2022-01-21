package organize

import "permissions/services/common/page/dto"

type SearchOrganize struct {
	dto.BasePage
	Organize
}

type Organize struct {
	Name string `json:"name" gorm:"size:50;not null"`
	Pid  uint   `json:"pid" gorm:"default:0"`
}
