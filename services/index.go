package services

import "permissions/services/system"

type ServiceGroup struct {
	SysServiceGroup system.SysServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
