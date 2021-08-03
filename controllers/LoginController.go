package controllers

import (
	"easebase_web/utils"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
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

	//通过用户ID获取用户信息
	var user orm.Params
	user, err := utils.GetUserByUserName(username)
	fmt.Println("user=", user)
	fmt.Println(user)
	fmt.Println("login_name=", user["login_name"])
	fmt.Println("User_ID=", user["id"])

	token, err := utils.GenerateToken(user["id"].(string), user["login_name"].(string), 600)
	if err == nil {
		c.SuccessJson("LoginController->Post", token)
	} else {
		c.ErrorJson("LoginController->Post", 500, err.Error(), nil)
	}
}
