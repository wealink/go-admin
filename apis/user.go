package apis

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

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
	util.Response(c, e.Code_200, data, e.GetSuccess)

}

// 通过ID获取用户信息
func GetUser(c *gin.Context) {
	var user models.User
	user.Id = com.StrTo(c.Param("id")).MustInt()
	util.Response(c, e.Code_200, user.GetUser(), e.GetSuccess)
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
		user.Password = util.Md5Pwd(user.Password)
		user.AddUser()
		util.Response(c, e.Code_200, "", e.CreatedSuccess)
	} else {
		util.Response(c, e.Code_400, err, e.Msg_400)
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
			util.Response(c, e.Code_200, "", e.UpdatedSuccess)
		} else {
			util.Response(c, e.Code_404, "", e.Msg_404)
		}
	} else {
		util.Response(c, e.Code_400, "", e.Msg_400)
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
		util.Response(c, e.Code_200, "", e.DeletedSuccess)
	} else {
		util.Response(c, e.Code_404, "", e.Msg_404)
	}
}

//重置用户密码
func ResetUserPwd(c *gin.Context) {
	var user models.User
	//user.Id = com.StrTo(c.Param("id")).MustInt()
	err := c.BindJSON(&user)
	if err == nil {
		if user.ExistUserByID() {
			user.Password = util.Md5Pwd(user.Password)
			user.ResetUserPwd()
			util.Response(c, e.Code_200, "", e.UpdatedSuccess)
		} else {
			util.Response(c, e.Code_404, "", e.Msg_404)
		}
	} else {
		util.Response(c, e.Code_400, "", e.Msg_400)
	}
}
