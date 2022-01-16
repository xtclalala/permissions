package initServe

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"permissions/global"
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
	v.SetConfigFile(utils.ConfigName + utils.Rung + env + utils.Point + utils.ConfigType)
	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
		panic(fmt.Errorf("配置文件读取失败: %s \n", env))
	}
	if err := v.Unmarshal(&global.System); err != nil {
		panic(fmt.Errorf("结构化数据失败：%s \n", err))
	}
	v.OnConfigChange(func(event fsnotify.Event) {
		fmt.Println("配置文件数据更改：", event.Name)
		if err := v.Unmarshal(&global.System); err != nil {
		}
	})

	return v
}
