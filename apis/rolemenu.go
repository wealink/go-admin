package apis

import (
	"gin-example/models"
	"gin-example/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取菜单树
func GetTreeRoleMenus(c *gin.Context) {
	var menu models.Menu
	var rolemenu models.RoleMenu
	rolemenu.Roleid = com.StrTo(c.Param("id")).MustInt()
	data := make(map[string]interface{})
	data["menus"]= menu.GetTreeMenus()
	data["checkedKeys"] = rolemenu.GetRoleMenu()
	c.JSON(e.Code_200, gin.H{
		"code": e.Code_200,
		"msg":     e.GetSuccess,
		"data":    data,
	})
}

//获取菜单角色列表
func GetRoleMenus(c *gin.Context)  {
	var rolemenu models.RoleMenu
	//data := make(map[string]interface{})
	//pageIndex, pageSize := util.GetPage(c)
	//data["list"] = menu.GetMenus()
	//data["total"] = menu.GetMenuTotal()
	c.JSON(e.Code_200, gin.H{
		"code": e.Code_200,
		"msg":     e.GetSuccess,
		"data":    rolemenu.GetRoleMenus(),
	})
}
