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
	global.Db = initServe.InitDb()
	if global.Db != nil {
		initServe.InitTables(global.Db)

		defer func(db *gorm.DB) {
			sqlDb, err := db.DB()
			if err != nil {
				sqlDb.Close()
			}
		}(global.Db)
	}

	// 可以分开创建数据 再绑定关联 用Save
	// 一张有数据 另一张没数据 可以在创建的同时关联 用Create
	//p := system.SysPermission{
	//	BaseID: model.BaseID{
	//		ID: 1,
	//	},
	//	Name: "qaz",
	//	Sort: 1,
	//}
	//p2 := system.SysPermission{
	//	BaseID: model.BaseID{
	//		ID: 2,
	//	},
	//	Name: "qaz",
	//	Sort: 2,
	//}
	//var menu system.SysMenu
	//global.Db.Where("id = ?", 2).Find(&menu)
	//err := global.Db.Model(&menu).Association("SysPermissions").Append([]system.SysPermission{p, p2})
	//fmt.Println(err)
}
