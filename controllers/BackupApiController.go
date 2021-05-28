package controllers

import (
	"easebase_web/utils"
	"encoding/json"
	"fmt"
	//"github.com/beego/beego/v2/client/orm"
	//"net/url"
)

type BackupApiController struct {
	BaseController
}

// backup api interface
func (c *BackupApiController) Get() {
	flag := c.Ctx.Input.Param(":flag")
	tag := c.Ctx.Input.Param(":tag")
	if flag == "config" {
		cfg, err := utils.GetBackupConfig(tag)
		if err == nil {
			c.SuccessJson("BackupApiController->Get", &cfg)
		} else {
			c.ErrorJson("BackupApiController->Get", 500, err.Error(), nil)
		}
	} else {
		c.ErrorJson("BackupApiController->Get", 500, "非法参数", nil)
	}
}

// backup api interface
func (c *BackupApiController) Post() {
	//read variable parameter
	flag := c.Ctx.Input.Param(":flag")
	tag := c.Ctx.Input.Param(":tag")
	if flag == "config" {
		cfg, err := utils.GetBackupConfig(tag)
		if err == nil {
			c.SuccessJson("BackupApiController->Get", &cfg)
		} else {
			c.ErrorJson("BackupApiController->Get", 500, err.Error(), nil)
		}
	} else if flag == "log" {
		par := map[string]string{}
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &par)
		fmt.Println("log par=", par)
		if err != nil {
			c.ErrorJson("BackupApiController->log", 500, err.Error(), nil)
		}
		res := utils.SaveBackupTotal(par)
		if res.Code == 0 {
			c.SuccessJson("BackupApiController->log", res.Msg)
		} else {
			c.ErrorJson("BackupApiController->log", 500, res.Msg, nil)
		}

	} else if flag == "detail" {
		par := map[string]string{}
		fmt.Println("detail par=", par)
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &par)
		if err != nil {
			c.ErrorJson("BackupApiController->detail", 500, err.Error(), nil)
		}
		res := utils.SaveBackupDetail(par)
		if res.Code == 0 {
			c.SuccessJson("BackupApiController->detail", res.Msg)
		} else {
			c.ErrorJson("BackupApiController->detail", 500, res.Msg, nil)
		}
	} else if flag == "push" {

	} else if flag == "run" {

	} else if flag == "stop" {

	} else {
		c.ErrorJson("BackupApiController->Get", 500, "非法参数", nil)
	}

}
