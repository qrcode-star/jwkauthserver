package user

import (
	"github.com/gin-gonic/gin"
	"jwkauthserver/app/service/user"
	"net/http"
)

func SignUP(c *gin.Context){
	username:=c.PostForm("username")
	passwd:=c.PostForm("password")
	passwd2:=c.PostForm("password2")
	//var data map[string]string
	data:=map[string]string{
		"passport":username,
		"password":passwd,
		"password2":passwd2,
	}
	if signupstat:=user.SignUp(data); signupstat == nil {

		c.JSON(http.StatusOK, "注册成功")
	} else {
		c.JSON(http.StatusBadRequest, signupstat.Error())
	}

}

func SignIN(c *gin.Context){
	username:=c.PostForm("username")
	passwd:=c.PostForm("password")
	if signinstat:=user.SignIn(username,passwd); signinstat==nil {
		usertoken,err := user.Usertoken(username)
		if err == nil {
			c.JSON(http.StatusOK, usertoken)
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, signinstat.Error())
	}
}