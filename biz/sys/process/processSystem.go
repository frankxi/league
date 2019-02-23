package process

// 系统业务类
import (
	"github.com/frankxi/league/biz/sys/models"
	"github.com/frankxi/league/comm/dao"
)

// ID查询
func GetSystem(id int) (models.System, error) {
	var q models.System
	return q.GetById(dao.GetDB(), id)
}

// 列表
func ListSystem(query map[string]interface{}) ([]models.System, error) {
	var q models.System
	return q.QueryList(dao.GetDB(), query)
}

// 分页
func PageSystem(query models.SystemQuery) (models.SystemPage, error) {
	return query.QueryPage(dao.GetDB())
}

// 新增
func InsertSystem(m *models.System) error {
	return m.Insert(dao.GetDB())
}

// 更新
func UpdateSystem(m *models.System) error {
	return m.Update(dao.GetDB())
}

// 删除
func DeleteSystem(m *models.System) error {
	return nil
}
