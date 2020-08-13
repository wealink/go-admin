package models

import (
	"gin-example/global/orm"
	"github.com/jinzhu/gorm"
	"time"
)

type Menu struct {
	Base
	Name   string `json:"name"`
	Type   string `json:"type"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

func (Menu) TableName() string {
	return "go_menu"
}

//通过菜单id查询菜单是否存在
func (menu *Menu) ExistMenuByID(id int) bool {
	orm.Db.Table(menu.TableName()).Select("id").Where("id = ?", id).First(&menu)

	if menu.Id > 0 {
		return true
	}

	return false
}

//获取所有菜单
func (menu *Menu) GetMenus(pageNum int, pageSize int) (menus []Menu) {
	orm.Db.Table(menu.TableName()).Where("").Offset(pageNum).Limit(pageSize).Find(&menus)
	return
}

//获取所有菜单总和
func (menu *Menu) GetMenuTotal() (count int) {
	orm.Db.Table(menu.TableName()).Model(&Menu{}).Where("").Count(&count)
	return
}

//添加菜单
func (menu *Menu) AddMenu(data map[string]interface{}) bool {
	orm.Db.Table(menu.TableName()).Create(&Menu{
		Name:   data["name"].(string),
		Type:   data["type"].(string),
		Path:   data["path"].(string),
		Method: data["method"].(string),
	})
	return true
}

//编辑菜单
func (menu *Menu) EditMenu(id int, data map[string]interface{}) bool {
	orm.Db.Table(menu.TableName()).Model(&Menu{}).Where("id = ?", id).Update(data)
	return true
}

//删除用户
func (menu *Menu) DeleteMenu(id int) {
	orm.Db.Table(menu.TableName()).Where("id = ?", id).Delete(Menu{})
}

//创建时间和修改时间更新
func (menu *Menu) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}
func (menu *Menu) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
