package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Mysql  MySQLSetting  `mapstructure:"mysql"`
	Redis  RedisSetting  `mapstructure:"redis"`
	Logger LoggerSetting `mapstructure:"logger"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdlConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTIme int    `mapstructure:"connMaxLifeTime"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"logLevel"`
	FileLogName string `mapstructure:"fileLogName"`
	MaxBackups  int    `mapstructure:"maxBackups"`
	MaxAge      int    `mapstructure:"maxAge"`
	MaxSize     int    `mapstructure:"maxSize"`
	Compress    bool   `mapstructure:"compress"`
}
