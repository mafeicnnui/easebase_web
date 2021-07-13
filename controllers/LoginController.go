package controllers

import (
	"easebase_web/utils"
	"fmt"
)

type LoginController struct {
	BaseController
}

//query dm by dmm
func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println(username, password)

	/*
	 1.检测用户是否存在
	 2.检查密码是否正确
	 3.检测用户是否过期
	 4.检测用户是否禁用
	*/

	token, err := utils.GenerateToken(60)
	if err == nil {
		c.SuccessJson("DmController->Get", token)
	} else {
		c.ErrorJson("DmController->Get", 500, err.Error(), nil)
	}
}
