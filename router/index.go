package router

import (
	"permissions/router/system"
)

type RouterGroup struct {
	System system.SystemRouterGroup
}

var AppRouter = new(RouterGroup)
