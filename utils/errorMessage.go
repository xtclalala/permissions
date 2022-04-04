package utils

var (
	SUCCESS = 200
	ERROR   = 500

	// Controller Error
	ParamsResolveFault = 1001

	// User Error
	UsernameAndPasswdError = 2001
	FindOrgError           = 2102
	FindRoleError          = 2103
	FindPermissionError    = 2104
	FindMenuError          = 2105
	FindUserError          = 2106

	UpdateRoleError       = 2201
	UpdateRoleMenusError  = 2202
	UpdateRolePerError    = 2203
	UpdatePermissionError = 2204
	UpdateOrgBaseError    = 2205
	UpdateMenuBaseError   = 2206

	CreateRoleError         = 2301
	CreatePermissionError   = 2302
	CreateOrganizationError = 2303
	CreateMenuError         = 2304

	DeleteUserError         = 2401
	DeleteRoleError         = 2402
	DeletePermissionError   = 2403
	DeleteOrganizationError = 2404
	DeleteMenuError         = 2405

	// Token Error
	TokenExpired      = 3001
	TokenNotValid     = 3002
	TokenInvalid      = 3003
	TokenMalformed    = 3004
	ClaimParseFailed  = 3005
	TokenCreateFailed = 3006
	NOTOKEN           = 3007
)

var codeMsg = map[int]string{
	SUCCESS: "操作成功",
	ERROR:   "操作失败",

	ParamsResolveFault: "sys.",

	UsernameAndPasswdError: "sys.user.usernameAndPasswdError",
	FindOrgError:           "sys.user.findOrgError",
	FindRoleError:          "sys.user.findRoleError",
	FindPermissionError:    "sys.user.findPermissionError",
	FindMenuError:          "sys.user.findMenuError",
	FindUserError:          "sys.user.findUserError",

	UpdateRoleError:       "sys.user.updateRoleError",
	UpdateRoleMenusError:  "sys.user.updateRoleMenusError",
	UpdateRolePerError:    "sys.user.updateRolePerError",
	UpdatePermissionError: "sys.user.updatePermissionError",
	UpdateOrgBaseError:    "sys.user.updateOrgBaseError",
	UpdateMenuBaseError:   "sys.user.updateMenuBaseError",

	CreateRoleError:         "sys.user.createRoleError",
	CreatePermissionError:   "sys.user.createPermissionError",
	CreateOrganizationError: "sys.user.createOrganizationError",
	CreateMenuError:         "sys.user.createMenuError",

	DeleteUserError:         "sys.user.deleteUserError",
	DeleteRoleError:         "sys.user.deleteRoleError",
	DeletePermissionError:   "sys.user.deletePermissionError",
	DeleteOrganizationError: "sys.user.deleteOrganizationError",
	DeleteMenuError:         "sys.user.deleteMenuError",

	NOTOKEN:           "sys.token.noToken",
	TokenExpired:      "sys.token.expired",
	TokenInvalid:      "sys.token.invalid",
	TokenMalformed:    "sys.token.malformed",
	TokenNotValid:     "sys.token.notValid",
	ClaimParseFailed:  "sys.claim.parseError",
	TokenCreateFailed: "sys.token.createFailed",
}

func GetErrorMessage(code int) string {
	return codeMsg[code]
}
