package Databace

import (
	"demo/Global"
	"demo/Viper"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql_Run() {
	Viper.Config()                //加载配置文件，并将读取的数据保存在Qwe.Config中
	dsn := Global.Config.Mysql.Dsn() //获取格式化后的Mysql连接数据
	fmt.Println("当前连接Mysql数据库信息为:", Global.Config.Mysql.Dsn())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Mysql数据库连接失败：", err)
	} else {
		fmt.Println("Mysql数据库连接成功")
	}

	// 自动迁移,初始化数据库表
	Aerr := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Global.User{})

	if Aerr != nil {
		fmt.Printf("自动迁移失败：%v", err)
	} else {
		fmt.Println("自动迁移并建立表单成功！")
	}
	Global.DB = db
}

