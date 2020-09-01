package util

import (
	"github.com/gin-gonic/gin"
	"gin-example/pkg/jwt"
)

func GetUserName(c *gin.Context) string {
	claims, _ := c.Get("claims")
	v := claims.(*jwt.Claims)
	if v.Username != ""{
		return v.Username
	}
	return ""
}
