package Rout

import (
	"demo/Databace_Operate"
	"demo/Tool"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GETlogin(c *gin.Context) { //登陆页面渲染
	c.HTML(http.StatusOK, "index.html", nil)
}

func GETforget(c *gin.Context) { //重置密码页面渲染
	c.HTML(http.StatusOK, "forget.html", nil)
}

func GEThome(c *gin.Context) {
	name, _ := c.GetQuery("name")
	token, _ := c.GetQuery("token")

	JT, _, err := Tool.ParseToken(token)

	//这里从数据库找到相应的名字，并传回前端
	if JT.Valid && err == nil {
		userdata := Databace_O.User_data_Queries(name)
		data, _ := json.Marshal(gin.H{ //测试代码数据
			"username": userdata.Username,
			"email":    userdata.Email,
		})
		c.HTML(http.StatusOK, "home.html", string(data)) //将携带数据转换成json格式和html页面再传给前端。
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}
