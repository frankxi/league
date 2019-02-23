package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/frankxi/league/biz/sys/models"
	"github.com/frankxi/league/biz/sys/process"
	"github.com/frankxi/league/comm/e"
	"github.com/frankxi/league/utils"
)

// 新增菜单
func InsertSysMenu(c *gin.Context) {
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
	utils.OK(c, nil)
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
	utils.OK(c, nil)
}

// 删除菜单
func DeleteSysMenu(c *gin.Context) {
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
	utils.OK(c, nil)
}
