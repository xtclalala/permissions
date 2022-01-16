package conf

type Config struct {
	App *App `mapstructure:"app" yaml:"app"`
	Db  *Db  `mapstructure:"db"  yaml:"db"`
}

type App struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port string `mapstructure:"port" yaml:"port"`
}

type Db struct {
	DbType string `mapstructure:"type" yaml:"type"`
	Host   string `mapstructure:"host" yaml:"host"`
	Port   string `mapstructure:"port" yaml:"port"`
	User   string `mapstructure:"user" yaml:"user"`
	Passwd string `mapstructure:"passwd" yaml:"passwd"`
	DbName string `mapstructure:"dbName" yaml:"dbName"`
}
