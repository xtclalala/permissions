package initServe

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"permissions/global"
	"permissions/model/system"
	"time"
)

func InitDb() *gorm.DB {
	dbConfig := global.System.Db
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Passwd,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName)

	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		SkipInitializeWithVersion: false,
	}
	// 命名策略
	name := schema.NamingStrategy{
		SingularTable: true,
	}

	gormConfig := gorm.Config{
		// gorm日志模式:silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务(提高运行速度)
		SkipDefaultTransaction: true,
		NamingStrategy:         name,
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig)
	if err != nil {
		panic(fmt.Errorf("连接数据库失败:%s", err))
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(10 * time.Second)
	return db
}

func InitTables(db *gorm.DB) {
	// 加 model
	err := db.AutoMigrate(&system.SysUser{}, &system.SysRole{}, &system.SysMenu{}, &system.SysPermission{})
	if err != nil {
		panic(fmt.Errorf("注册表格失败", err))
	}
	// 设置其他引擎
	//db.Set("gorm:table_options","ENGINE=MyIsAm").AutoMigrate()
	fmt.Println("表格注册成功")
}
