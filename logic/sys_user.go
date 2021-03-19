package logic

import (
	"Chatin/global"
	"Chatin/model"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) MSG {
	name, haveName := c.GetPostForm("name")

	if !haveName {
		return ERROR("no account", 500)
	}
	password, havePassword := c.GetPostForm("password")
	if !havePassword {
		return ERROR("no account", 500)
	}
	user := new(model.User)
	global.DB.Where("account = ? & password = ?", name, password).Find(&user)
	return OK(user, "success")
}
