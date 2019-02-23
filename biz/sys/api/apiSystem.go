package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/frankxi/league/biz/sys/models"
	"github.com/frankxi/league/biz/sys/process"
	"github.com/frankxi/league/comm/e"
	"github.com/frankxi/league/utils"
)

// 新增系统
func insertSystem(c *gin.Context) {
	var p models.System
	if err := c.Bind(&p); err != nil {
		log.Error("insertSystem=", p)
		utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("insertSystem=", p)
	err := process.InsertSystem(&p)
	if err != nil {
		log.Error("insertSystem.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c, nil)
}

// 更新系统
func updateSystem(c *gin.Context) {
	var p models.System
	if err := c.Bind(&p); err != nil {
		log.Error("updateSystem=", p)
		utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("updateSystem=", p)
	err := process.UpdateSystem(&p)
	if err != nil {
		log.Error("updateSystem.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c, nil)
}

// 删除系统
func deleteSystem(c *gin.Context) {
	var p models.System
	if err := c.Bind(&p); err != nil {
		log.Error("deleteSystem=", p)
		utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("deleteSystem=", p)
	err := process.DeleteSystem(&p)
	if err != nil {
		log.Error("deleteSystem.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c, nil)
}
