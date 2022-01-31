package conf

type Config struct {
	App    *App    `mapstructure:"app" yaml:"app"`
	Db     *Db     `mapstructure:"db"  yaml:"db"`
	Logger *Logger `mapstructure:"logger"  yaml:"logger"`
	Jwt    *Jwt    `mapstructure:"jwt"  yaml:"jwt"`
}

type App struct {
	Port string `mapstructure:"port" yaml:"port"`
	Auth string `mapstructure:"auth" yaml:"auth"`
}

type Db struct {
	DbType string `mapstructure:"type" yaml:"type"`
	Host   string `mapstructure:"host" yaml:"host"`
	Port   string `mapstructure:"port" yaml:"port"`
	User   string `mapstructure:"user" yaml:"user"`
	Passwd string `mapstructure:"passwd" yaml:"passwd"`
	DbName string `mapstructure:"dbName" yaml:"dbName"`
}

type Logger struct {
	FilePath string `mapstructure:"filePath" yaml:"filePath"`
	FileName string `mapstructure:"fileName" yaml:"fileName"`
}

type Jwt struct {
	SignKey    string `mapstructure:"signKey" yaml:"signKey"`
	Timeout    int    `mapstructure:"timeout" yaml:"timeout"`
	Iss        string `mapstructure:"iss" yaml:"iss"`
	BufferTime int    `mapstructure:"bufferTime" yaml:"bufferTime"`
}
