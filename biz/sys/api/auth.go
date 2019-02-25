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
// @Summary GetAuth
// @Tags username  password
// @Security Bearer
// @Produce  json
// @Param username path string true "username"
// @Param password path string true "password"
// @Resource username password
// @Router /auth?name={username}&password={password} [get]
// @Success 200 {object} utils.Ret
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
