package Email

import (
	"demo/Global"
	"demo/Tool"
	"fmt"
	"time"
	"gopkg.in/gomail.v2"
)

func Send_Email(email string) string {
	Vkey,_:=Global.RDS.Do("get",email+"Vkey")
	fmt.Println("vkey:",Vkey)
	if Vkey!=nil{
		fmt.Println("请求过于频繁")
		return "请求过于频繁"
	} 

	vcode := Tool.KeyRand(6)
	now := time.Now()
	t := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	html := fmt.Sprintf(`<div>
		<div>
			尊敬的%s，您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您于 %s 提交的邮箱验证，本次验证码为 %s，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>	
	</div>`, email, t, vcode)

	m := gomail.NewMessage()
	config := gomail.NewDialer(
		Global.Config.Email.Host,
		Global.Config.Email.Port,
		Global.Config.Email.Username,
		Global.Config.Email.Password,
	)
	m.SetAddressHeader("From",Global.Config.Email.Username,"hackerxiao")
	m.SetHeader("To",email)
	m.SetHeader("Subject", "[我的验证码]邮箱验证")
	m.SetBody("text/html", html) 
	err2 :=config.DialAndSend(m)
	if err2!=nil {
		fmt.Println("验证码发送失败！",err2)
		return "验证码发送失败！"
	} else{
		fmt.Printf("地址%s的注册验证码发送成功",email)
		//设置发送频率和保存验证码
		Global.RDS.Do("set",email+"Vkey",1,"EX",60)
		Global.RDS.Do("set",email,vcode,"EX",500)
		return "验证码发送成功"
	}
}
