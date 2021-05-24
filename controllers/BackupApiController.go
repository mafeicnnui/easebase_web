package controllers

import (
	"easebase_web/utils"
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
	} else if flag == "log" {

	} else if flag == "detail" {

	} else if flag == "push" {

	} else if flag == "run" {

	} else if flag == "stop" {

	} else {
		c.ErrorJson("BackupApiController->Get", 500, "非常参数", nil)
	}

}
