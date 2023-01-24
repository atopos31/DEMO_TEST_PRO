package Global

//用户信息结构体
type User struct{
	Username string 
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//登录信息结构体
type Login struct{
	Email string `json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

//重置密码结构体
type Update struct{
	Email string `json:"email" binding:"required"`
	Captcha string `json:"captcha" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//注册信息结构体
type Register struct{
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
	Captcha string `json:"captcha" binding:"required"`
}