package Databace

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"demo/Global"
	"time"
)

func RedisPollInit() *redis.Pool { //设置连接池并创建连接
	return &redis.Pool{
		MaxIdle:     Global.Config.Rdis.MaxIdle,   //最大空闲数
		MaxActive:   Global.Config.Rdis.MaxActive, //最大连接数，0不设上
		Wait:        Global.Config.Rdis.Wait,
		IdleTimeout: time.Duration(Global.Config.Rdis.IdleTimeout) * time.Second, //空闲等待时间

		Dial: func() (redis.Conn, error) { //创建链接
			fmt.Println(Global.Config.Rdis)
			Base := redis.DialDatabase(Global.Config.Rdis.Db)                                                     //连接到指定的数据库号码
			c, err := redis.Dial(Global.Config.Rdis.Method, Global.Config.Mysql.Host+":"+Global.Config.Rdis.Port, Base) //redis IP地址
			if err != nil {
				fmt.Println(err)
				return nil, err //连接失败返回空，而不是错误信息
			}
			fmt.Println("Ridis数据库初始化成功！")
			return c, err //返回一个结构体，包含redis.Conn和error类型的两个量
		},
	}

}

func Run_RedisInit() { //初始化Ridis数据库连接
	Global.RDS = RedisPollInit().Get()//从结构体redis.Pool中取出redis.Conn
}

func RedisClose() {
	_ = Global.RDS.Close() //关闭数据库连接，返回错误信息，一般不会出错
}
