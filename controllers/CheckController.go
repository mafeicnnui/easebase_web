package controllers

import (
	"easebase_web/utils"
	"fmt"
)

type CheckController struct {
	BaseController
}

//query dm by dmm
func (c *CheckController) Post() {
	token := c.GetString("token")
	fmt.Println("token=", token)
	if utils.ValidateToken(token) != nil {
		c.ErrorJson("CheckController->Get", 402, "认证过期!", nil)
	} else {
		c.SuccessJson("CheckController->Get", token)
	}
}
