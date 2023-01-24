package Rout

import (
	"demo/Databace_Operate"
	"demo/Global"
	"demo/Tool"
	"demo/email"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func POSTlogin(c *gin.Context) {
	var userdata Global.User
	err := c.ShouldBindJSON(&userdata) //接受登录信息
	if err != nil {
		fmt.Printf("获取登录信息出错，错误信息：%v", err)
		return
	}
	//获取数据后与数据库对比
	//如果与数据库数据相符，返回ok：true，js会向home发出GET请求
	name,OK :=Databace_O.Login_verify(&userdata)
	if OK {
		token,_:=Tool.Token(userdata)
		c.JSON(http.StatusOK, gin.H{
			"ok":       true,
			"name":     name,
			"token":    token,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":       false,
			"name":		nil,
			"token":    nil, 
		})
	}

	fmt.Println(userdata)
}

//找回密码请求
func POSTforget(c *gin.Context) {
	email := c.PostForm("email") //接受找回邮箱
	//之后会进入发送验证码阶段
	_, err := Databace_O.Register_EV(email)
	if err {
		c.String(http.StatusOK, "该邮箱尚未注册")
		return
	}
	reponse := Email.Send_Email(email)
	c.String(http.StatusOK, reponse)
	fmt.Println(email)
}

// 重置密码请求
func POSTupdate(c *gin.Context) {
	var update Global.Update
	c.ShouldBindJSON(&update)
	OK := Databace_O.Verify_Vcode(update.Email, update.Captcha)
	if OK {
		Databace_O.Change_password(update.Email,update.Password)
		c.String(http.StatusOK,"修改成功")
	} else {
		c.String(http.StatusOK,"验证码错误")
	}
	fmt.Println(update)
}

//发送注册验证码请求
func send(c *gin.Context) {
	email := c.PostForm("email") //接受注册邮箱
	re, err := Databace_O.Register_EV(email)
	if !err {
		c.String(http.StatusOK, re)
		return
	}
	//之后会进入发送验证码阶段
	reponse := Email.Send_Email(email)
	c.String(http.StatusOK, reponse)
	fmt.Println(email)
}

//注册请求
func POSTregister(c *gin.Context) {
	//这里处理注册信息，需要验证redis数据库里面的验证码
	var registerdata Global.Register
	err := c.ShouldBindJSON(&registerdata)
	if err != nil {
		fmt.Printf("获取注册信息出错，错误信息：%v", err)
		return
	}

	re, err3 := Databace_O.Register_UV(registerdata) //检测是否重复注册
	if !err3 {
		c.String(http.StatusOK, re)
		return
	}
	//将验证码与redis数据库中的数据对比，相符，注册成功，将用户信息导入Mysql
	OK := Databace_O.Verify_Vcode(registerdata.Email, registerdata.Captcha)
	fmt.Println(OK)
	if OK {
		fmt.Println("hello")
		reponse := Databace_O.Rreister_Data(registerdata)
		c.String(http.StatusOK, reponse)
		return
	} else {
		c.String(http.StatusOK, "验证码错误！")
		return
	}
}
