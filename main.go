package main

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod download

import (
	"gorm.io/gorm"
	"permissions/global"
	"permissions/initServe"
)

func main() {
	global.Viper = initServe.InitConfig()
	db := initServe.InitDb()
	if db != nil {
		initServe.InitTables(db)

		defer func(db *gorm.DB) {
			sqlDb, err := db.DB()
			if err != nil {
				sqlDb.Close()
			}
		}(db)
	}

}
