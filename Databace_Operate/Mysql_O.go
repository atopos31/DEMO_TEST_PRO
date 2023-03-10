package Databace_O

import (
	"demo/Global"
)

func Login_verify(Userdata *Global.User) (string,bool) { //登陆验证
	var V_userdata Global.User
	Global.DB.Where("email=?", Userdata.Email).First(&V_userdata)
	if Userdata.Password == V_userdata.Password {
		return V_userdata.Username,true
	} else {
		return "nil",false
	}
}


func User_data_Queries(Name string) Global.User { //用户信息查找
	var U_userdata Global.User
	Global.DB.Where("username=?", Name).First(&U_userdata)
	return U_userdata
}

func Register_EV (email string) (string,bool){
	var UU Global.User
	err1 :=Global.DB.Where("email = ?",email).First(&UU).Error
	if err1==nil {
		return "邮箱已被注册",false
	}  else {
		return  "验证通过",true
	}
} 


func Register_UV (R_Data Global.Register) (string,bool){
	var UU Global.User
	err1 :=Global.DB.Where("username = ?",R_Data.Username).First(&UU).Error
	if err1==nil {
		return "用户名已存在",false
	}  else {
		return  "验证通过",true
	}
} 

func Rreister_Data(R_Data Global.Register) string {//注册成功，数据导入数据库
	R_put_D := Global.User{
		Username: R_Data.Username,
		Email: R_Data.Email,
		Password: R_Data.Password,
	}
	Global.DB.Create(R_put_D)
	return "注册成功！"
}

func Change_password (email string,new_password string) {
	var UU Global.User
	Global.DB.Model(&UU).Where("email = ?",email).Update("password",new_password)
}

