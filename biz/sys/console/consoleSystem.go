package console

import (
	"github.com/gin-gonic/gin"
	"../models"
	"../process"
 	log "github.com/sirupsen/logrus"
	"github.com/frankxi/league/comm/e"
	"github.com/frankxi/league/utils"
)

// 注册API系统
// {"id":0,"name":"","code":"","desc":"","logo":"","isEnabled":0,"createdBy":"","createdTm":"","updatedBy":"","updatedTm":"","version":0,"isDeleted":0}
// 所有字段
// id,name,code,desc,logo,is_enabled,created_by,created_tm,updated_by,updated_tm,version,is_deleted
func AddConsoleSystem(base *gin.RouterGroup) {
	api := base.Group("/system")
	api.POST("/list", listSystem)
	api.POST("/insert", insertSystem)
	api.POST("/update", updateSystem)
	api.POST("/delete", deleteSystem)
}
// 分页获取系统
func listSystem(c *gin.Context) {
	var p models.SystemQuery
	if err := c.Bind(&p); err != nil {
		log.Error("listSystem=", p)
		utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("listSystem=", p)
	page, err := process.PageSystem(p)
	if err != nil {
		log.Error("listSystem.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c,page)
}

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
	utils.OK(c,nil)
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
	utils.OK(c,nil)
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
	utils.OK(c,nil)
}
