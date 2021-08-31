package controllers

import (
	"easebase_web/utils"
	"encoding/json"
	"fmt"
)

type PublicApiController struct {
	BaseController
}

// backup api interface
func (c *PublicApiController) Post() {
	flag := c.Ctx.Input.Param(":flag")
	if flag == "decrypt" {
		par := map[string]string{}
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &par)
		fmt.Println("par=", par)
		res, err := utils.Decrypt(par["password"], par["key"])
		fmt.Println("PublicApiController=res", res)
		if err == nil {
			c.SuccessJson("PublicApiController->Post->Decrypt2222", &res)
		} else {
			c.ErrorJson("PublicApiController->Post->Decrypt", 500, err.Error(), nil)
		}
	} else if flag == "encrypt" {
		par := map[string]string{}
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &par)
		res, err := utils.Encrypt(par["password"], par["key"])
		if err == nil {
			c.SuccessJson("PublicApiController->Post->Encrypt", &res)
		} else {
			c.ErrorJson("PublicApiController->Post->Encrypt", 500, err.Error(), nil)
		}
	} else {
		c.ErrorJson("PublicApiController->Post", 500, "非法参数", nil)
	}
}

func (c *PublicApiController) Get() {
	flag := c.Ctx.Input.Param(":flag")
	key := c.GetString("key")
	password := c.GetString("password")
	if flag == "decrypt" {
		res, err := utils.Decrypt(password, key)
		if err == nil {
			c.SuccessJson("PublicApiController->Post->Decrypt", &res)
		} else {
			c.ErrorJson("PublicApiController->Post->Decrypt", 500, err.Error(), nil)
		}
	} else if flag == "encrypt" {
		res, err := utils.Encrypt(password, key)
		if err == nil {
			c.SuccessJson("PublicApiController->Post->Encrypt", &res)
		} else {
			c.ErrorJson("PublicApiController->Post->Encrypt", 500, err.Error(), nil)
		}
	} else {
		c.ErrorJson("PublicApiController->Post", 500, "非法参数", nil)
	}
}
