package controllers

import (
	"easebase_web/utils"
	"encoding/json"
	//"github.com/beego/beego/v2/client/orm"
	//"net/url"
)

type PublicApiController struct {
	BaseController
}

// backup api interface
func (c *PublicApiController) Post() {
	flag := c.Ctx.Input.Param(":flag")
	if flag == "decode" {
		par := map[string]string{}
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &par)
		res, err := utils.Decrypt(par["password"], par["key"])
		if err == nil {
			c.SuccessJson("PublicApiController->Get", &res)
		} else {
			c.ErrorJson("PublicApiController->Get", 500, err.Error(), nil)
		}
	} else {
		c.ErrorJson("PublicApiController->Get", 500, "非法参数", nil)
	}
}
