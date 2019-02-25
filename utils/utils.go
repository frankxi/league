package utils

import (
	"github.com/frankxi/league/comm/e"
	"github.com/gin-gonic/gin"
)

type Ret struct {
	Code int
	Desc string
	Data interface{}
}

//gin response method

func OK(c *gin.Context, Data interface{}) {
	ret := new(Ret)
	ret.Code = e.OK
	ret.Desc = ""
	ret.Data = Data
	c.JSON(ret.Code, ret)
}

func Err(c *gin.Context, errCode int) {
	ret := new(Ret)
	ret.Code = errCode
	ret.Desc = e.CodeDescMap[errCode]
	ret.Data = ""
	c.JSON(errCode, ret)
}
