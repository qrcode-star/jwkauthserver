package api

import (
	"github.com/gin-gonic/gin"
	"jwkauthserver/app/service/jwk"
	"net/http"
)

func Genjwk(c *gin.Context){
	keyid,err:=service.Insertjwk()
	if err ==nil {
		c.JSON(http.StatusOK, keyid)
	} else {
		c.JSON(http.StatusBadRequest,"GenJwkFail!")
	}
}

func Getpvk(c *gin.Context){
	keyid:=c.PostForm("keyid")
	pvk,err,error:=service.Selectpvkjwk(keyid)
	if err == nil && error == nil {
		c.JSON(http.StatusOK, pvk)
	} else {
		c.JSON(http.StatusBadRequest, "GetPvkFAID!")
	}
}

func Getpuk(c *gin.Context){
	keyid:=c.PostForm("keyid")
	puk,err,error:=service.Selectpukjwk(keyid)
	if err == nil && error == nil {
		c.JSON(http.StatusOK, puk)
	} else {
		c.JSON(http.StatusBadRequest, "GetPukFAID!")
	}
}