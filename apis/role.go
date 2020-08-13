package apis

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type Role struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

//获取所有角色
func GetRoles(c *gin.Context) {
	var role models.Role
	role.Name = c.Query("name")
	role.Status = c.Query("status")
	data := make(map[string]interface{})
	pageIndex, pageSize := util.GetPage(c)
	data["list"] = role.GetRoles(pageIndex, pageSize)
	data["total"] = role.GetRolesTotal()
	data["pageIndex"] = pageIndex
	data["pageSize"] = pageSize
	c.JSON(e.Code_200, gin.H{
		"code": e.Code_200,
		"msg":  e.GetSuccess,
		"data": data,
	})
}

//通过id获取角色信息
func GetRole(c *gin.Context) {
	var role models.Role
	role.Id = com.StrTo(c.Param("id")).MustInt()
	c.JSON(e.Code_200, gin.H{
		"code": e.Code_200,
		"msg":  e.GetSuccess,
		"data": role.GetRole(),
	})
}

//新增用户
func AddRole(c *gin.Context) {
	var role models.Role
	err := c.BindJSON(&role)
	if err == nil {
		rs := role.AddRole()
		if rs == true {
			c.JSON(e.Code_200, gin.H{
				"code": e.Code_200,
				"data": make(map[string]interface{}),
				"msg":  e.CreatedSuccess,
			})
			return
		} else {
			c.JSON(e.Code_500, gin.H{
				"code": e.Code_500,
				"data": nil,
				"msg":  "角色已经存在！！！",
			})
		}
	} else {
		c.JSON(e.Code_400, gin.H{
			"code": e.Code_400,
			"data": err,
			"msg":  e.Msg_400,
		})
		return
	}
}

//编辑角色
func EditRole(c *gin.Context) {
	var role models.Role
	role.Id = com.StrTo(c.Param("id")).MustInt()
	err := c.BindJSON(&role)
	if err == nil {
		if role.ExistRoleByID() {
			role.EditRole()
			c.JSON(e.Code_200, gin.H{
				"code": e.Code_200,
				"data": make(map[string]interface{}),
				"msg":  e.UpdatedSuccess,
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
	} else {
		c.JSON(e.Code_400, gin.H{
			"code": e.Code_400,
			"data": "",
			"msg":  e.Msg_400,
		})
		return
	}
}

//删除用户
func DeleteRole(c *gin.Context) {
	var role models.Role
	role.Id = com.StrTo(c.Param("id")).MustInt()
	if role.ExistRoleByID() {
		role.DeleteRole(role.Id)
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
