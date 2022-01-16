package initServe

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"permissions/global"
	"time"
)

func InitDb() {
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

	sql, _ := db.DB()
	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
	sql.SetConnMaxLifetime(10 * time.Second)

}
