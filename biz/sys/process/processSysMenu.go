package process

// 菜单业务类
import (
	"github.com/frankxi/league/biz/sys/models"
	"github.com/frankxi/league/comm/dao"
)

// ID查询
func GetSysMenu(id int) (models.SysMenu, error) {
	var q models.SysMenu
	return q.GetById(dao.GetDB(), id)
}

// 列表
func ListSysMenu(query map[string]interface{}) ([]models.SysMenu, error) {
	var q models.SysMenu
	return q.QueryList(dao.GetDB(), query)
}

// 分页
func PageSysMenu(query models.SysMenuQuery) (models.SysMenuPage, error) {
	return query.QueryPage(dao.GetDB())
}

// 新增
func InsertSysMenu(m *models.SysMenu) error {
	return m.Insert(dao.GetDB())
}

// 更新
func UpdateSysMenu(m *models.SysMenu) error {
	return m.Update(dao.GetDB())
}

// 删除
func DeleteSysMenu(m *models.SysMenu) error {
	return nil
}
