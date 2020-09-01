package apis

import (
	"crypto/md5"
	"fmt"
	"gin-example/models"
	"gin-example/pkg/captcha"
	"gin-example/pkg/e"
	"gin-example/pkg/jwt"
	"gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

type Auth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Uuid     string `json:"uuid" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

// @Summary 获取登录token
// @Tags 用户
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"token创建成功"}"
// @Router /login [GET]
func Login(c *gin.Context) {
	var (
		code int
		msg  string
		user models.User
		auth Auth
		err  error
	)
	err = c.BindJSON(&auth)
	data := make(map[string]interface{})
	if err == nil {
		has := md5.Sum([]byte(auth.Password))
		isExist := user.Login(auth.Username, fmt.Sprintf("%x", has))
		if isExist {
			rolename := user.GetRoleNameByUserName(auth.Username)
			rs := captcha.Verify(auth.Uuid, auth.Code)
			if rs == true {
				token, err := jwt.GenerateToken(auth.Username, auth.Password, rolename)
				if err != nil {
					code = e.Code_500
					msg = "token创建失败"
					data = nil
				} else {
					data["token"] = token
					code = e.Code_200
					msg = "登录成功"
				}
			} else {
				code = e.Code_500
				msg = "验证码错误"
				data = nil
			}

		} else {
			code = e.Code_404
			msg = e.Msg_404
			data = nil
		}
	} else {
		code = e.Code_400
		msg = e.Msg_400
		data = nil
	}
	//登录日志
	LoginLogToDB(c, code, msg, auth.Username)
	//response
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Logout(c *gin.Context) {
	code := e.Code_200
	msg := "登出成功"
	//登出日志
	LoginLogToDB(c, code, msg, util.GetUserName(c))
	//response
	c.JSON(code, gin.H{
		"code": code,
		"data": "",
		"msg":  msg,
	})
}

// 添加登录登出记录
func LoginLogToDB(c *gin.Context, code int, msg string, username string) {
	var loginlog models.Loginlog
	loginlog.Username = username
	if code == e.Code_200 {
		loginlog.Status = "0"
	} else {
		loginlog.Status = "1"
	}
	ua := user_agent.New(c.Request.UserAgent())
	loginlog.Ipaddr = c.ClientIP()
	loginlog.Loginlocation = util.GetLocation(loginlog.Ipaddr)
	browserName, browserVersion := ua.Browser()
	loginlog.Browser = browserName + " " + browserVersion
	loginlog.Os = ua.OS()
	loginlog.Logintime = util.GetCurrentTime()
	loginlog.Msg = msg
	loginlog.AddLoginLog()
}

func Info(c *gin.Context) {
	var user models.User
	data := make(map[string]interface{})
	claims, _ := c.Get("claims")
	v := claims.(*jwt.Claims)
	user.Username = v.Username
	rs := user.GetUser()
	data["username"] = rs.Username
	data["phone"] = rs.Phone
	data["avatar"] = rs.Avatar
	data["create_on"] = rs.Created_On
	var roles []string
	roles = append(roles, v.Rolename)
	data["roles"] = roles
	c.JSON(e.Code_200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "info",
	})
}
