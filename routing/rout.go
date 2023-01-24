package Rout

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 路由以及页面的初始化
func Run_Rout() {

	r := gin.Default()

	r.Static("./static", "static") //加载静态文件，css js之类的
	//加载三个html页面文件
	r.LoadHTMLFiles("./html/index.html", "./html/home.html", "./html/forget.html")

	r.Use() //中间件

	r.GET("/login", GETlogin)   //登录页面
	r.GET("/forget", GETforget) //忘记密码页面
	r.GET("/home", GEThome)     //主页面

	r.POST("/login", POSTlogin)       //登录请求
	r.POST("/forget", POSTforget)     //重置密码的验证码请求
	r.POST("/update", POSTupdate)     //重置数据并更新数据到数据库
	r.POST("/send-captcha", send)     //验证码请求
	r.POST("/register", POSTregister) //注册请求

	err := r.Run()
	if err != nil {
		fmt.Println("理由加载失败:", err)
		return
	} else {
		fmt.Println("路由加载成功")
	}
}
