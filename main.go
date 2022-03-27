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

	// 可以分开创建数据 再绑定关联 用Save
	// 一张有数据 另一张没数据 可以在创建的同时关联 用Create
	//p := system.SysPermission{
	//	BaseID:model.BaseID{
	//		2,
	//	},
	//	Name: "cecece2",
	//	Sort: 2,
	//}
	//p2 := system.SysPermission{
	//	BaseID: model.BaseID{
	//		ID: 2,
	//	},
	//	Name: "qaz",
	//	Sort: 2,
	//}
	//var menu system.SysMenu
	//err := global.Db.Where(&system.M2mRolePermission{SysPermissionId: 2}).Delete(&system.M2mRolePermission{}).Error
	//err := global.Db.(&p).Error
	//fmt.Println(err)
}
