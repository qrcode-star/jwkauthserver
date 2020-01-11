package main

import (
	"github.com/gin-gonic/gin"
	"jwkauthserver/app/api/user"
)

func main() {
	r := gin.Default()
	rsa := r.Group("/user")
	rsa.POST("/signup",user.SignUP)
	rsa.POST("/signin",user.SignIN)
	r.Run(":80")
}