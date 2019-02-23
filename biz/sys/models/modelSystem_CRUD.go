package models

import (
		"github.com/jinzhu/gorm"
		"github.com/frankxi/league/biz"
		"errors"

) 
// 系统类
type System struct {
	Id	int64	`json:"id" gorm:"primary_key"`	// id
	Name	string	`json:"name" `	// 系统姓名
	Code	string	`json:"code" gorm:"default:''"`	// 系统编码（暂定6位数）
	Desc	string	`json:"desc" `	// 系统说明
	Logo	string	`json:"logo" gorm:"default:''"`	// 系统logo
	IsEnabled	int	`json:"isEnabled" gorm:"default:'1'"`	// 是否有效，默认 1 已启用，-1 未启用。
	CreatedBy	string	`json:"createdBy" gorm:"default:'0'"`	// 创建人USER_ID
	CreatedTm	string	`json:"createdTm" gorm:"default:'CURRENT_TIMESTAMP(6)'"`	// 创建时间，默认是 CURRENT_TIMESTAMP(6)
	UpdatedBy	string	`json:"updatedBy" gorm:"default:'0'"`	// 修改人 USER_ID
	UpdatedTm	string	`json:"updatedTm" gorm:"default:'CURRENT_TIMESTAMP(6)'"`	// 修改时间，修改时 CURRENT_TIMESTAMP(6)
	Version	int	`json:"version" gorm:"default:'0'"`	// 版本号乐观锁
	IsDeleted	int	`json:"isDeleted" gorm:"default:'1'"`	// 记录逻辑状态，是否被删除，1 未删除， -1 已删除。
}
//id,name,code,desc,logo,is_enabled,created_by,created_tm,updated_by,updated_tm,version,is_deleted

// 分页查询条件
type SystemQuery struct {
	System
	biz.Page
}

// 分页查询结果
type SystemPage struct {
	List []System `json:"list"`
	biz.Page
}

// 主键查询(注意gorm.ErrRecordNotFound)
func (m *System) GetById(db *gorm.DB, id int) (System, error) {
	var one System
	return one, db.Model(m).First(&one, id).Error
}

// 条件查询
func (m *System) QueryList(db *gorm.DB, filter map[string]interface{}) ([]System, error) {
	var result []System
	return result, db.Model(*m).Where(filter).Find(&result).Error
}

// 分页查询
func (m *SystemQuery) QueryPage(db *gorm.DB) (SystemPage, error) {
	var page SystemPage
	page.Page = m.Page
	db = db.Model(m.System).Where(m.System)
	err := db.Count(&page.Total).Error
	if err != nil {
		return page, err
	}
	err = db.Offset((page.PageNo - 1) * page.PageSize).Limit(page.PageSize).Find(&page.List).Error
	return page, err
}

// 新增
func (m *System) Insert(db *gorm.DB) error {
	return db.Model(m).Create(m).Error
}

// 更新(注意不能自动更新值为''或0)
func (m *System) Update(db *gorm.DB) error {
	if m.Id == 0 {
		return errors.New("没有查询条件！")
	}
	return db.Model(m).Update(m).Error
}

// 更新
func (m *System) UpdateWhere(db *gorm.DB, filter map[string]interface{}, fields map[string]interface{}) error {
	// 检查查询条件
	if len(filter) == 0 {
		return errors.New("没有查询条件！")
	}
	return db.Model(m).Where(filter).Update(fields).Error
}
