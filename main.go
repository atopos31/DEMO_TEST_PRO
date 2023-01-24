package main

import (
	"demo/Databace"
	"demo/Viper"
	"demo/routing"
)


func main(){
	Viper.Config()//加载配置文件
	Databace.Mysql_Run()//启动Mysql数据库，创建新的表
	Databace.Run_RedisInit()//启动Ridis数据库(连接池方式)
	Rout.Run_Rout()//加载路由
}