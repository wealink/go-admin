package routers

import (
	"gin-example/apis"
	_ "gin-example/docs"
	"gin-example/middleware/jwt"
	"gin-example/middleware/permission"
	"gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(util.Cors())
	r.Use(gin.Recovery())
	gin.SetMode("debug")
	g := r.Group("")
	NoAuth(g)
	Auth(g)
	return r
}

func NoAuth(g *gin.RouterGroup) {
	g.POST("/login", apis.Login)
	g.GET("/logout", apis.Logout)
	g.GET("/getcode", apis.GetCode)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func Auth(g *gin.RouterGroup) {
	apiv1 := g.Group("/api/v1")
	//token校验
	apiv1.Use(jwt.JWT())
	//不需要控制接口权限
	//用户相关
	apiv1.GET("/info", apis.Info)
	//角色菜单关联
	apiv1.GET("/treemenus",apis.GetTreeMenus)
	apiv1.GET("/treerolemenus/:id",apis.GetTreeRoleMenus)
	apiv1.GET("/rolemenus",apis.GetRoleMenus)
	//接口权限校验
	apiv1.Use(permission.AuthCheckRole())
	//用户
	apiv1.GET("/users", apis.GetUsers)
	apiv1.GET("/user/:id", apis.GetUser)
	apiv1.POST("/users", apis.AddUser)
	apiv1.PUT("/users/:id", apis.EditUser)
	apiv1.PUT("/user/:id",apis.ResetUserPwd)
	apiv1.DELETE("/users/:id", apis.DeleteUser)
	//角色
	apiv1.GET("/roles", apis.GetRoles)
	apiv1.GET("/role/:id", apis.GetRole)
	apiv1.POST("/roles", apis.AddRole)
	apiv1.PUT("/roles/:id", apis.EditRole)
	apiv1.DELETE("/roles/:id", apis.DeleteRole)
	//菜单
	apiv1.GET("/menus", apis.GetMenus)
	apiv1.GET("/menu/:id", apis.GetMenu)
	apiv1.POST("/menus", apis.AddMenu)
	apiv1.PUT("/menus/:id", apis.EditMenu)
	apiv1.DELETE("/menus/:id", apis.DeleteMenu)

}
