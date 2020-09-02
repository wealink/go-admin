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
	util.Response(c, e.Code_200, data, e.GetSuccess)
}

//通过id获取角色信息
func GetRole(c *gin.Context) {
	var role models.Role
	role.Id = com.StrTo(c.Param("id")).MustInt()
	util.Response(c, e.Code_200, role.GetRole(), e.GetSuccess)
}

//新增用户
func AddRole(c *gin.Context) {
	var role models.PostRole
	err := c.BindJSON(&role)
	if err == nil {
		rs := role.AddRole()
		if rs == true {
			util.Response(c, e.Code_200, "", e.CreatedSuccess)
		} else {
			util.Response(c, e.Code_500, "", "角色已经存在！")
		}
	} else {
		util.Response(c, e.Code_400, err, e.Msg_400)
	}
}

//编辑角色
func EditRole(c *gin.Context) {
	//var role models.Role
	var role models.PostRole
	role.Id = com.StrTo(c.Param("id")).MustInt()
	err := c.BindJSON(&role)
	if err == nil {
		if role.ExistRoleByID() {
			role.EditRole()
			util.Response(c, e.Code_200, "", e.UpdatedSuccess)
		} else {
			util.Response(c, e.Code_404, "", e.Msg_404)
		}
	} else {
		util.Response(c, e.Code_400, "", e.Msg_400)
	}
}

//删除用户
func DeleteRole(c *gin.Context) {
	var role models.Role
	var rolemenu models.RoleMenu
	role.Id = com.StrTo(c.Param("id")).MustInt()
	if role.ExistRoleByID() {
		rolemenu.DeleteRoleMenu(role.Id)
		role.DeleteRole(role.Id)
		util.Response(c, e.Code_200, "", e.DeletedSuccess)
	} else {
		util.Response(c, e.Code_404, "", e.Msg_404)
	}
}
