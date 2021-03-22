package logic

import (
	"Chatin/global"
	"Chatin/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) MSG {
	account, haveAccount := c.GetPostForm("user")
	if !haveAccount {
		return ERROR("no account", 500)
	}
	password, havePassword := c.GetPostForm("password")
	if !havePassword {
		return ERROR("no password", 500)
	}
	user := new(model.User)
	var ct int64
	global.DB.Where("account = ? AND password = ?", account, password).Count(&ct)
	if ct==0{
		return ERROR("no exist", 500)
	}
	global.DB.Where("account = ? AND password = ?", account, password).First(&user)
	return OK(user, "登录成功")
}
