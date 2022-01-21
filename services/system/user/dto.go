package user

import "permissions/services/common/page/dto"

type SearchUser struct {
	dto.BasePage
	User
}

type User struct {
	Username  string `json:"username"`
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}
