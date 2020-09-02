package util

import (
	"crypto/md5"
	"fmt"
	"gin-example/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func GetUserName(c *gin.Context) string {
	claims, _ := c.Get("claims")
	v := claims.(*jwt.Claims)
	if v.Username != "" {
		return v.Username
	}
	return ""
}

func Md5Pwd(password string) string {
	has := md5.Sum([]byte(password))
	return fmt.Sprintf("%x", has)
}
