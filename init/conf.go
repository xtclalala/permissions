package init

import (
	"fmt"
	"github.com/spf13/viper"
	"permissions/utils"
)

func InitConfig() *viper.Viper {
	temp := viper.New()
	temp.SetConfigFile(utils.ConfigPath)
	if err := temp.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件")
	}
	env := temp.GetString("env")

	v := viper.New()
	v.SetConfigName(utils.ConfigName + utils.Point + env)
	v.SetConfigType(utils.ConfigType)
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("")
	}






	return v
}
