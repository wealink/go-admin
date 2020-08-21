package models

import (
	"gin-example/global/orm"
	"github.com/jinzhu/gorm"
	"time"
)

type Menu struct {
	Base
	Name   string `json:"name"`
	Path   string `json:"path"`
	Title   string `json:"title"`
	Icon string `json:"icon"`
	Componenturl string `json:"componenturl"`
	Parentid int `json:"parentid"`
	Type string `json:"type"`
	Method string `json:"method"`
}

type MenuView struct {
	Menu
	Children []MenuView `json:"children"`
}

type MenuRoles struct {
	Menu
	Roles []string `json:"roles"`
}

type MenuViewRoles struct {
	Menu
	Children []Menu `json:"children"`
	Roles []string `json:"roles"`
}

type MenuLable struct {
	Id       int         `json:"id"`
	Label    string      `json:"label"`
	Children []MenuLable `json:"children"`
}

func (Menu) TableName() string {
	return "go_menu"
}

//通过菜单id查询菜单是否存在
func (menu *Menu) ExistMenuByID() bool {
	orm.Db.Table(menu.TableName()).Select("id").Where("id = ?", menu.Id).First(&menu)

	if menu.Id > 0 {
		return true
	}

	return false
}

//获取所有菜单
func (menu *Menu) GetMenus()(M []MenuView) {
	var C,J []MenuView
	if menu.Type != "" {
		orm.Db.Table(menu.TableName()).Where("type = ?",menu.Type).Find(&M)
	} else {
		orm.Db.Table(menu.TableName()).Where("type = 'M'").Find(&M)
		for i := 0; i < len(M); i++ {
			orm.Db.Table(menu.TableName()).Where("parentid = ? ", M[i].Id).Find(&C)
			for j := 0; j < len(C); j++ {
				orm.Db.Table(menu.TableName()).Where("parentid = ? ", C[j].Id).Find(&J)
				C[j].Children = J
			}
			M[i].Children = C
		}
	}
	return
}

//获取菜单树
func (menu *Menu) GetTreeMenus()(m []MenuLable) {
	menus := menu.GetMenus()
	var tmp MenuLable
	for i := 0; i < len(menus); i++ {
		tmp.Id = menus[i].Id
		tmp.Label = menus[i].Title
		tmp.Children = nil
		for j := 0; j < len(menus[i].Children); j++{
			var tmp1 MenuLable
			tmp1.Id = menus[i].Children[j].Id
			tmp1.Label = menus[i].Children[j].Title
			//tmp.Children = append(tmp.Children,tmp1)
			for k := 0; k < len(menus[i].Children[j].Children); k++{
				var tmp2 MenuLable
				tmp2.Id = menus[i].Children[j].Children[k].Id
				tmp2.Label = menus[i].Children[j].Children[k].Title
				tmp1.Children = append(tmp1.Children,tmp2)
			}
			tmp.Children = append(tmp.Children,tmp1)
		}
		m = append(m,tmp)
	}
	return
}


//通过id获取菜单信息
func (menu *Menu) GetMenu() (menus Menu) {
	table := orm.Db.Table(menu.TableName())
	table.Where("id = ?", menu.Id).Find(&menus)
	return
}

//获取所有菜单总和
func (menu *Menu) GetMenuTotal() (count int) {
	orm.Db.Table(menu.TableName()).Where("componenturl = 'Layout'").Count(&count)
	return
}

//添加菜单
func (menu *Menu) AddMenu() bool {
	// 判断菜单是否冲突
	var count int
	if menu.Type == "J" {
		orm.Db.Table(menu.TableName()).Where("title = ? ", menu.Title).Count(&count)
	} else {
		orm.Db.Table(menu.TableName()).Where("name = ?", menu.Name).Count(&count)
	}
	if count > 0 {
		return false
	}
	if menu.Type == "M" {
		orm.Db.Table(menu.TableName()).Create(&Menu{
			Name:         menu.Name,
			Path:         menu.Path,
			Title:        menu.Title,
			Icon:         menu.Icon,
			Componenturl: menu.Componenturl,
			Type: menu.Type,
		})
		return true
	} else if menu.Type == "C" {
		orm.Db.Table(menu.TableName()).Create(&Menu{
			Name:         menu.Name,
			Path:         menu.Path,
			Title:        menu.Title,
			Icon:         menu.Icon,
			Componenturl: menu.Componenturl,
			Parentid:     menu.Parentid,
			Type: menu.Type,
		})
		return true
	} else {
		orm.Db.Table(menu.TableName()).Create(&Menu{
			Path:         menu.Path,
			Title:        menu.Title,
			Parentid:     menu.Parentid,
			Method: menu.Method,
			Type: menu.Type,
		})
		return true
	}
	return false
}

//编辑菜单
func (menu *Menu) EditMenu() bool {
	orm.Db.Table(menu.TableName()).Model(&Menu{}).Where("id = ?", menu.Id).Update(&menu)
	return true
}

//删除菜单
func (menu *Menu) DeleteMenu() {
	orm.Db.Table(menu.TableName()).Where("id = ?", menu.Id).Delete(Menu{})
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
