package v1

import (
	"permissions/api/v1/system"
)

type ApiGroup struct {
	SysApiGroup system.SysApiGroup
}

var ApiGroupApp = new(ApiGroup)
