package api

import (
	"github.com/gin-gonic/gin"

	"github.com/frankxi/league/comm/e"
	"github.com/frankxi/league/utils"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}
// @Summary Print
// @Tags username  password
// @Security Bearer
// @Produce  json
// @Param name path string true "username"
// @Resource username
// @Router /api/v1/auth?name={username}&password={password} [get]
// @Success 200 {object} utils.ReturnJSON
func GetAuth(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	//TODO check user login info

	token, err := utils.GenerateToken(username, password)
	if err != nil {
		utils.Err(c, e.ErrNoUserIdCode)
		return
	}
	utils.OK(c, map[string]string{
		"token": token,
	})
}
