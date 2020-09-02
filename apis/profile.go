package apis

import (
	"fmt"
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/jwt"
	"gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Pass struct {
	Oldpassword string `json:"oldpassword"`
	Newpassword string `json:"newpassword"`
}

//个人中心 更新密码
func Pwd(c *gin.Context) {
	var user models.User
	var pass Pass
	data := make(map[string]interface{})
	claims, _ := c.Get("claims")
	v := claims.(*jwt.Claims)
	user.Username = v.Username
	err := c.BindJSON(&pass)
	if err == nil {
		rs := user.Pwd(util.Md5Pwd(pass.Oldpassword), util.Md5Pwd(pass.Newpassword))
		if rs == true {
			util.Response(c, e.Code_200, data, e.UpdatedSuccess)
		} else {
			util.Response(c, e.Code_500, "", e.Msg_500)
		}
	} else {
		util.Response(c, e.Code_400, "", e.Msg_400)
	}
}

//个人中心修改头像
func UploadAvatar(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	guid := uuid.New().String()
	filPath := "static/" + guid + ".jpg"
	for _, file := range files {
		// 上传文件至指定目录
		err := c.SaveUploadedFile(file, filPath)
		if err != nil {
			fmt.Println(err)
		}
	}
	var user models.User
	claims, _ := c.Get("claims")
	v := claims.(*jwt.Claims)
	user.Username = v.Username
	user.Avatar = "/" + filPath
	user.UploadAvatar()
	util.Response(c, e.Code_200, user.Avatar, e.UpdatedSuccess)
}
