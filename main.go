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
	//u := system.User{
	//	UserBaseInfo: system.UserBaseInfo{
	//		Username: "12345",
	//		Password: "1214566",
	//		LoginName: "12356633",
	//	},
	//	SysOrganizeIds: []uint{1},
	//	SysRoleIds: []uint{1},
	//}
	//s, err := utils.Validate(&u)
	//fmt.Println(s)
	//fmt.Println(err)
	//initServe.RunWindowServer()

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
