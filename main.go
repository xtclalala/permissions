package main

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod download

import (
	"gorm.io/gorm"
	"permissions/global"
	initServe2 "permissions/initServe"
)

func main() {
	global.Viper = initServe2.InitConfig()
	global.Db = initServe2.InitDb()
	if global.Db != nil {
		initServe2.InitTables(global.Db)

		defer func(db *gorm.DB) {
			sqlDb, err := db.DB()
			if err != nil {
				sqlDb.Close()
			}
		}(global.Db)
	}
	initServe2.RunWindowServer()
}
