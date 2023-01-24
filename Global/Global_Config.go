package Global

import (
	"demo/Structs"

	"github.com/golang-jwt/jwt"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)


var Config Structs_config.Config //全局配置信息
var DB *gorm.DB  //全局Mysql数据库接口，适用于各个请求访问数据库数据
var RDS redis.Conn //全局Ridis数据库接口

type Claims struct {
	Email string // 用户名
	Password string //用户密码
	jwt.StandardClaims 
}
