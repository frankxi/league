package console

import (
	"github.com/gin-gonic/gin"
	"../models"
	"../process"
 	log "github.com/sirupsen/logrus"
	"github.com/frankxi/league/comm/e"
	"github.com/frankxi/league/utils"
)

// 注册API菜单
// {"menu_id":0,"parentId":0,"name":"","url":"","perms":"","type":0,"icon":"","orderNum":0}
// 所有字段
// menu_id,parent_id,name,url,perms,type,icon,order_num
func AddConsoleSysMenu(base *gin.RouterGroup) {
	api := base.Group("/sysMenu")
	api.POST("/list", listSysMenu)
	api.POST("/insert", insertSysMenu)
	api.POST("/update", updateSysMenu)
	api.POST("/delete", deleteSysMenu)
}
// 分页获取菜单
func listSysMenu(c *gin.Context) {
	var p models.SysMenuQuery
	if err := c.Bind(&p); err != nil {
		log.Error("listSysMenu=", p)
		utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("listSysMenu=", p)
	page, err := process.PageSysMenu(p)
	if err != nil {
		log.Error("listSysMenu.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c,page)
}

// 新增菜单
func insertSysMenu(c *gin.Context) {
	var p models.SysMenu
	if err := c.Bind(&p); err != nil {
		log.Error("insertSysMenu=", p)
		utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("insertSysMenu=", p)
	err := process.InsertSysMenu(&p)
	if err != nil {
		log.Error("insertSysMenu.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c,nil)
}

// 更新菜单
func updateSysMenu(c *gin.Context) {
	var p models.SysMenu
	if err := c.Bind(&p); err != nil {
		log.Error("updateSysMenu=", p)
		utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("updateSysMenu=", p)
	err := process.UpdateSysMenu(&p)
	if err != nil {
		log.Error("updateSysMenu.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c,nil)
}

// 删除菜单
func deleteSysMenu(c *gin.Context) {
	var p models.SysMenu
	if err := c.Bind(&p); err != nil {
		log.Error("deleteSysMenu=", p)
		utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("deleteSysMenu=", p)
	err := process.DeleteSysMenu(&p)
	if err != nil {
		log.Error("deleteSysMenu.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c,nil)
}
