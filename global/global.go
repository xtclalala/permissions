package global

import (
	"github.com/spf13/viper"
	"permissions/conf"
)

var (
	Viper  *viper.Viper
	System conf.Config
)
