package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

type ReturnMsg struct {
	Name string
	Code int
	Msg  string
	Data interface{}
}

func (c *BaseController) SuccessJson(name string, data interface{}) {
	res := ReturnMsg{
		name, 200, "success", data,
	}
	c.Data["json"] = res
	c.ServeJSON() //对json进行序列化输出
	c.StopRun()
}

func (c *BaseController) ErrorJson(name string, code int, msg string, data interface{}) {
	res := ReturnMsg{
		name, code, msg, data,
	}

	c.Data["json"] = res
	c.ServeJSON() //对json进行序列化输出
	c.StopRun()
}
