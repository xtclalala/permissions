package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"permissions/conf"
)

var (
	Viper  *viper.Viper
	System conf.Config
	Db     *gorm.DB
)
