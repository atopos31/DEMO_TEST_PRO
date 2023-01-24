package Structs_config

type Common struct {
	Orgin string `mapstructure:"orgin"`
	Port  string `mapstructure:"port"`
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Method   string `mapstructure:"method"`
	Database string `mapstructure:"database"`
	Config   string `mapstructure:"config"`
}

type Rdis struct {
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	Db          int    `mapstructure:"db"`
	Password    string `mapstructure:"password"`
	Method      string `mapstructure:"method"`
	MaxIdle     int    `mapstructure:"MaxIdle"`
	MaxActive   int    `mapstructure:"MaxActive"`
	Wait        bool   `mapstructure:"Wait"`
	IdleTimeout int    `mapstructure:"IdleTimeout"`
}

type Email struct {
	Host string		`mapstructure:"host"`
	Port int `mapstructure:"port"`
	Username string  `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Expires  int    `mapstructure:"expires"`  
}

type Jwt struct {
	Expires int `mapstructure:"expires"`  
	Issuer string  `mapstructure:"issuer"`  
	Key string `mapstructure:"key"`  
}

type Config struct {
	Common Common `mapstructure:"common"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Rdis   Rdis   `mapstructure:"rdis"`
	Email  Email `mapstructure:"email"`
	Jwt Jwt `mapstructure:"jwt"`  
}

//将Mysql连接数据格式化
func (Mc *Mysql) Dsn() string {
	return Mc.Username + ":" + Mc.Password + "@" + Mc.Method + "(" + Mc.Host + ":" + Mc.Port + ")" + "/" + Mc.Database + "?" + Mc.Config
}
