package controllers

import (
	"easebase_web/utils"
	"encoding/json"
	//"github.com/beego/beego/v2/client/orm"
	//"net/url"
)

type SyncApiController struct {
	BaseController
}

// sync api interface
func (c *SyncApiController) Get() {
	flag := c.Ctx.Input.Param(":flag")
	tag := c.Ctx.Input.Param(":tag")
	if flag == "config" {
		cfg, err := utils.GetSyncConfig(tag)
		if err == nil {
			c.SuccessJson("SyncApiController->Get", &cfg)
		} else {
			c.ErrorJson("SyncApiController->Get", 500, err.Error(), nil)
		}
	} else {
		c.ErrorJson("SyncApiController->Get", 500, "非法参数", nil)
	}
}

// backup api interface
func (c *SyncApiController) Post() {
	//read variable parameter
	flag := c.Ctx.Input.Param(":flag")
	tag := c.Ctx.Input.Param(":tag")
	if flag == "config" {
		cfg, err := utils.GetSyncConfig(tag)
		if err == nil {
			c.SuccessJson("SyncApiController->Get", &cfg)
		} else {
			c.ErrorJson("SyncApiController->Get", 500, err.Error(), nil)
		}
	} else if flag == "log" {
		par := map[string]string{}
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &par)
		if err != nil {
			c.ErrorJson("SyncApiController->log", 500, err.Error(), nil)
		}
		res := utils.SaveSyncTotal(par)
		if res.Code == 0 {
			c.SuccessJson("SyncApiController->log", res.Msg)
		} else {
			c.ErrorJson("SyncApiController->log", 500, res.Msg, nil)
		}

	} else if flag == "detail" {
		par := map[string]string{}
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &par)
		if err != nil {
			c.ErrorJson("SyncApiController->detail", 500, err.Error(), nil)
		}
		res := utils.SaveSyncDetail(par)
		if res.Code == 0 {
			c.SuccessJson("SyncApiController->detail", res.Msg)
		} else {
			c.ErrorJson("SyncApiController->detail", 500, res.Msg, nil)
		}
	} else if flag == "push" {

	} else if flag == "run" {

	} else if flag == "stop" {

	} else {
		c.ErrorJson("SyncApiController->Get", 500, "非法参数", nil)
	}

}
