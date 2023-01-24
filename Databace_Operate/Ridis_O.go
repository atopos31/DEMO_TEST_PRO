package Databace_O

import (
	"demo/Global"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func Verify_Vcode(email string, Vcode string) bool {//验证码验证
	Vcode_G, _ := redis.String(Global.RDS.Do("get", email))
	fmt.Println(Vcode)
	fmt.Println(Vcode_G)
	if Vcode == Vcode_G {
		return true
	} else {
		return false
	}
}
