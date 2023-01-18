package main

import (
	"fmt"
	"github.com/astaxie/beego/utils"
	"strings"
)

// SendEmail
//
//	@Description: 邮箱验证码
//	@param toEmail 请求验证码的邮箱
//	@param msg	验证码

func SendEmail(toEmail, msg string) {
	username := "*" // 发送者的邮箱地址
	password := "*" // 授权密码

	//QQ邮箱：SMTP服务器地址 smtp.qq.com (端口号：587)
	//雅虎邮箱：SMTP服务器地址 smtp.yahoo.com (端口号：587)
	//163邮箱：SMTP服务器地址 smtp.163.com (端口号：25)
	//126邮箱：SMTP服务器地址 smtp.126.com (端口号：25)
	//新浪邮箱：SMTP服务器地址 smtp.sina.com (端口号：25)
	host := "smtp.qq.com" // 邮件协议
	port := "587"         // 端口号

	emailConfig := fmt.Sprintf(`{"username":"%s","password":"%s","host":"%s","port":%s}`, username, password, host, port)
	//fmt.Println("emailConfig", emailConfig)
	emailConn := utils.NewEMail(emailConfig) // beego下的
	emailConn.From = strings.TrimSpace(username)
	emailConn.To = []string{strings.TrimSpace(toEmail)}
	emailConn.Subject = "注册验证码"
	//注意这里我们发送给用户的是激活请求地址  传入秒数
	emailConn.Text = fmt.Sprintf("您的注册验证码为 [ %s ] ,有效期至[ %s ]。", msg, GetCurrentTime(60))
	err := emailConn.Send()
	if err != nil {
		fmt.Println(err)
	}

}
