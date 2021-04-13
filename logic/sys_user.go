package logic

import (
	"Chatin/global"
	"Chatin/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
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
	global.DB.Table("users").Where("account = ? AND password = ?", account, password).Count(&ct)
	if ct==0{
		return ERROR("no exist", 500)
	}
	global.DB.Table("users").Where("account = ? AND password = ?", account, password).First(&user)
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	issuer := "frank"
	claims := global.Claims{
		ID:       user.ID,
		Account: user.Account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	if err!=nil{
		return ERROR("系统异常",500)
	}
	return OK(gin.H{
		"token":token,
		"name": user.Name,
		"config": user.Config,
		"avatar":user.Avatar,
	}, "登录成功")
}
