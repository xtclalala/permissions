package main

import (
	"permissions/global"
	"permissions/init"
)

func main() {
	global.Config = init.InitConfig()

}
