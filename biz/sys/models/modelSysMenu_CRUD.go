package models

import (
	"github.com/jinzhu/gorm"
	"github.com/frankxi/league/biz"
	"errors"
)

// 菜单类
type SysMenu struct {
	MenuId   int64  `json:"menu_id" gorm:"primary_key"` // id
	ParentId int    `json:"parentId" `                  // 父菜单ID，一级菜单为0
	Name     string `json:"name" `                      // 菜单名称
	Url      string `json:"url" `                       // 菜单URL
	Perms    string `json:"perms" `                     // 授权(多个用逗号分隔，如：user:list,user:create)
	Type     int    `json:"type" `                      // 类型   0：目录   1：菜单   2：按钮
	Icon     string `json:"icon" `                      // 菜单图标
	OrderNum int    `json:"orderNum" `                  // 排序
}

//menu_id,parent_id,name,url,perms,type,icon,order_num

// 分页查询条件
type SysMenuQuery struct {
	SysMenu
	biz.Page
}

// 分页查询结果
type SysMenuPage struct {
	List []SysMenu `json:"list"`
	biz.Page
}

// 主键查询(注意gorm.ErrRecordNotFound)
func (m *SysMenu) GetById(db *gorm.DB, id int) (SysMenu, error) {
	var one SysMenu
	return one, db.Model(m).First(&one, id).Error
}

// 条件查询
func (m *SysMenu) QueryList(db *gorm.DB, filter map[string]interface{}) ([]SysMenu, error) {
	var result []SysMenu
	return result, db.Model(*m).Where(filter).Find(&result).Error
}

// 分页查询
func (m *SysMenuQuery) QueryPage(db *gorm.DB) (SysMenuPage, error) {
	var page SysMenuPage
	page.Page = m.Page
	db = db.Model(m.SysMenu).Where(m.SysMenu)
	err := db.Count(&page.Total).Error
	if err != nil {
		return page, err
	}
	err = db.Offset((page.PageNo - 1) * page.PageSize).Limit(page.PageSize).Find(&page.List).Error
	return page, err
}

// 新增
func (m *SysMenu) Insert(db *gorm.DB) error {
	return db.Model(m).Create(m).Error
}

// 更新(注意不能自动更新值为''或0)
func (m *SysMenu) Update(db *gorm.DB) error {
	if m.MenuId == 0 {
		return errors.New("没有查询条件！")
	}
	return db.Model(m).Update(m).Error
}

// 更新
func (m *SysMenu) UpdateWhere(db *gorm.DB, filter map[string]interface{}, fields map[string]interface{}) error {
	// 检查查询条件
	if len(filter) == 0 {
		return errors.New("没有查询条件！")
	}
	return db.Model(m).Where(filter).Update(fields).Error
}
