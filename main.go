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
	//var role system.SysRole
	//global.Db.Where("id = ?", 2).First(&role)
	//fmt.Println(role)
	//var user system.SysUser
	//global.Db.Where("id = ?", "a8d57864-142d-458b-8c1e-70961ed5ff48").Find(&user)
	//user.SysRoles = []*system.SysRole{&role}
	//fmt.Println(user)
	//err := global.Db.Save(&user).Error
	//fmt.Println(err)
}
