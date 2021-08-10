package controllers

import (
	"easebase_web/utils"
)

type LoginController struct {
	BaseController
}

//query dm by dmm
func (c *LoginController) Post() {
	//获取请求参数
	username := c.GetString("username")
	password := c.GetString("password")

	//检测用户是否存在
	if utils.CheckUserExists(username) == 0 {
		c.ErrorJson("LoginController->Post", 500, "用户不存在!", nil)
	}

	//通过用户ID获取用户信息
	user, err := utils.GetUserByUserName(username)
	if err != nil {
		c.ErrorJson("LoginController->Post", 500, "获取用户信息出错!", nil)
	}

	//对已获取的口令进行加密
	password2, err2 := utils.Encrypt(password, user["login_name"].(string))
	if err2 != nil {
		c.ErrorJson("LoginController->Post", 500, "加密口令出错!", nil)
	}

	//检查密码是否正确
	if user["password"] != password2 {
		c.ErrorJson("LoginController->Post", 500, "口令不正确!", nil)
	}

	//检测用户是否禁用
	if user["status"] == "0" {
		c.ErrorJson("LoginController->Post", 500, "用户已禁用，请联系管理员!", nil)
	}

	//检测用户是否过期
	if user["expire_date"].(string) < utils.GetDate() {
		c.ErrorJson("LoginController->Post", 500, "该用户已过期，请联系管理员!", nil)
	}

	//生成token
	token, err := utils.GenerateToken(user["id"].(string), user["login_name"].(string), 600)
	if err == nil {
		c.SuccessJson("LoginController->Post", token)
	} else {
		c.ErrorJson("LoginController->Post", 500, err.Error(), nil)
	}
}
