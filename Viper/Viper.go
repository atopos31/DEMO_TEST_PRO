package Viper

import (
	"demo/Global"
	"fmt"
	"github.com/spf13/viper"
)

func Config() {
	viper.SetConfigName("Config")        // 配置文件名 (不带扩展格式)
	viper.SetConfigType("toml")          // 如果你的配置文件没有写扩展名，那么这里需要声明你的配置文件属于什么格式
	viper.AddConfigPath("./Config") // 配置文件的路径

	Rerr := viper.ReadInConfig() //找到并读取配置文件
	if Rerr != nil {
		fmt.Printf("读取文件出错：%v\n",Rerr)
		return 
	}
	Merr := viper.Unmarshal(&Global.Config)
	if Merr != nil {
		fmt.Printf("反序列化文件出错：%v\n",Merr)
		return 
	}
}
