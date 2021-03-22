package logic

import (
	"Chatin/global"
	"Chatin/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) MSG {
	name, haveName := c.GetPostForm("user")
	fmt.Printf(name)

	if !haveName {
		return ERROR("no account", 500)
	}
	password, havePassword := c.GetPostForm("password")
	fmt.Printf(password)
	if !havePassword {
		return ERROR("no password", 500)
	}
	user := new(model.User)
	global.DB.Where("account = ? & password = ?", name, password).Find(&user)
	return OK(user, "登录成功")
}
