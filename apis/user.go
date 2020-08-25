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
	"github.com/unknwon/com"
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
// @Success 200 {string} json "{"captcha":200,"data":{},"msg":"token创建成功"}"
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
		isExist, _ := user.Login(auth.Username, fmt.Sprintf("%x", has))
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
					msg = "token创建成功"
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
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Logout(c *gin.Context) {
	c.JSON(e.Code_200, gin.H{
		"code": e.Code_200,
		"data": "",
		"msg":  "logout",
	})
}

func Info(c *gin.Context) {
	data := make(map[string]interface{})
	claims, _ := c.Get("claims")
	v := claims.(*jwt.Claims)
	data["avatar"] = "/static/tezign.jpg"
	data["username"] = v.Username
	var roles []string
	roles = append(roles, v.Rolename)
	//data["rolename"] = v.Rolename
	data["roles"] = roles
	c.JSON(e.Code_200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "info",
	})
}

// @Summary 获取所有用户信息
// @Tags 用户
// @Produce  json
// @Param token query string true "身份令牌"
// @Success 200 {string} json "{"captcha":200,"data":{},"msg":"获取成功"}"
// @Router /api/v1/users [GET]
func GetUsers(c *gin.Context) {
	var user models.UserView
	user.Username = c.Query("username")
	user.Phone = c.Query("phone")
	user.Status = c.Query("status")
	data := make(map[string]interface{})
	pageIndex, pageSize := util.GetPage(c)
	data["list"] = user.GetUsers(pageIndex, pageSize)
	data["total"] = user.GetUserTotal()
	data["pageIndex"] = pageIndex
	data["pageSize"] = pageSize
	c.JSON(e.Code_200, gin.H{
		"code": e.Code_200,
		"msg":  e.GetSuccess,
		"data":  data,
	})

}

// 通过ID获取用户信息
//通过id获取角色信息
func GetUser(c *gin.Context) {
	var user models.User
	user.Id = com.StrTo(c.Param("id")).MustInt()

	c.JSON(e.Code_200, gin.H{
		"code": e.Code_200,
		"msg":  e.GetSuccess,
		"data": user.GetUser(),
	})
}

// @Summary 添加用户
// @Tags 用户
// @Produce  json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Param role_id query string true "角色ID"
// @Param token query string true "身份令牌"
// @Success 200 {string} json "{"captcha":200,"data":{},"msg":"添加成功"}"
// @Router /api/v1/users [POST]
func AddUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err == nil {
		//password, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		has := md5.Sum([]byte(user.Password))
		user.Password = fmt.Sprintf("%x", has)
		user.AddUser()
		c.JSON(e.Code_200, gin.H{
			"code": e.Code_200,
			"data":    "",
			"msg":     e.CreatedSuccess,
		})
		return
	} else {
		c.JSON(e.Code_400, gin.H{
			"code": e.Code_400,
			"data":    err,
			"msg":  e.Code_400,
		})
		return
	}
}

// @Summary 编辑用户信息
// @Tags 用户
// @Produce  json
// @Param password query string true "密码"
// @Param role_id query string true "角色ID"
// @Param token query string true "身份令牌"
// @Success 200 {string} json "{"captcha":200,"data":{},"msg":"编辑成功"}"
// @Router /api/v1/users/{id} [PUT]
func EditUser(c *gin.Context) {
	var user models.User
	user.Id = com.StrTo(c.Param("id")).MustInt()
	err := c.BindJSON(&user)
	if err == nil {
		if user.ExistUserByID() {
			user.EditUser()
			c.JSON(e.Code_200, gin.H{
				"code": e.Code_200,
				"data":    make(map[string]interface{}),
				"msg":     e.UpdatedSuccess,
			})
			return
		} else {
			c.JSON(e.Code_404, gin.H{
				"code": e.Code_404,
				"data":    "",
				"msg":     e.Msg_404,
			})
			return
		}
	} else {
		c.JSON(e.Code_400, gin.H{
			"code": e.Code_400,
			"data":    "",
			"msg":     e.Code_400,
		})
		return
	}
}

// @Summary 删除用户信息
// @Tags 用户
// @Produce  json
// @Success 200 {string} json "{"captcha":200,"data":{},"msg":"删除成功"}"
// @Router /api/v1/users/{id} [DELETE]
func DeleteUser(c *gin.Context) {
	var user models.User
	user.Id = com.StrTo(c.Param("id")).MustInt()
	if user.ExistUserByID() {
		user.DeleteUser()
		c.JSON(e.Code_200, gin.H{
			"code": e.Code_200,
			"data": make(map[string]interface{}),
			"msg":  e.DeletedSuccess,
		})
		return
	} else {
		c.JSON(e.Code_404, gin.H{
			"code": e.Code_404,
			"data": make(map[string]interface{}),
			"msg":  e.Msg_404,
		})
		return
	}
}

//重置用户密码
func ResetUserPwd(c *gin.Context) {
	var user models.User
	//user.Id = com.StrTo(c.Param("id")).MustInt()
	err := c.BindJSON(&user)
	if err == nil {
		if user.ExistUserByID() {
			has := md5.Sum([]byte(user.Password))
			user.Password = fmt.Sprintf("%x", has)
			user.ResetUserPwd()
			c.JSON(e.Code_200, gin.H{
				"code": e.Code_200,
				"data":    make(map[string]interface{}),
				"msg":     e.UpdatedSuccess,
			})
			return
		} else {
			c.JSON(e.Code_404, gin.H{
				"code": e.Code_404,
				"data":    "",
				"msg":     e.Msg_404,
			})
			return
		}
	} else {
		c.JSON(e.Code_400, gin.H{
			"code": e.Code_400,
			"data":    "",
			"msg":     e.Code_400,
		})
		return
	}
}