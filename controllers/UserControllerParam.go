package controllers

import (
	"easebase_web/models"
	"github.com/beego/beego/v2/client/orm"
	"strconv"

	"fmt"
)

//功能：用户信息表
type UserControllerByParId struct {
	BaseController
}

//delete user
func (c *UserControllerByParId) Delete() {
	id := c.Ctx.Input.Param(":id")
	int32, err := strconv.Atoi(id)
	if err != nil {
		c.ErrorJson("UserControllerByParId->Delete", 500, err.Error(), nil)
	}

	o := orm.NewOrm()
	user := models.TUser{Id: int32}
	_, err = o.Delete(&user)
	if err != nil {
		c.ErrorJson("UserControllerByParId->Delete", 500, err.Error(), nil)
	} else {
		c.SuccessJson("UserControllerByParId->Delete", nil)
	}
}

//query user
func (c *UserControllerByParId) Get() {
	id := c.Ctx.Input.Param(":id")
	int32, err := strconv.Atoi(id)
	if err != nil {
		c.ErrorJson("UserControllerByParId->Get", 500, err.Error(), nil)
	}
	o := orm.NewOrm()
	u := models.TUser{Id: int32}
	err = o.Read(&u)
	if err != nil {
		c.ErrorJson("UserControllerByParId->Get", 500, err.Error(), nil)

	} else {
		c.SuccessJson("UserControllerByParId->Get", &u)
	}
}

type UserControllerByParName struct {
	BaseController
}

func (c *UserControllerByParName) Get() {
	name := c.Ctx.Input.Param(":name")
	var users []*models.TUser
	orm := orm.NewOrm()
	qs := orm.QueryTable("user")
	n, err := qs.Filter("name__contains", name).All(&users) // 过滤器
	if err == nil && n >= 0 {
		c.SuccessJson("UserControllerByParName->Get", &users)
	} else {
		c.ErrorJson("UserControllerByParName->Get", 500, err.Error(), nil)
	}
}

type UserControllerByParName2 struct {
	BaseController
}

func (c *UserControllerByParName2) Get() {
	name := c.GetString("name")
	var users []*models.TUser
	orm := orm.NewOrm()
	qs := orm.QueryTable("user")
	n, err := qs.Filter("name__contains", name).All(&users) // 过滤器
	if err == nil && n > 0 {
		fmt.Println(&users)
		c.SuccessJson("UserControllerByParName->Get", &users)
	} else {
		c.ErrorJson("UserControllerByParName->Get", 500, err.Error(), nil)
	}
}

//功能：用户角色表

type UserRoleControllerByParId struct {
	BaseController
}

//query userRole
func (c *UserRoleControllerByParId) Get() {
	id := c.Ctx.Input.Param(":userId")
	userId, err := strconv.Atoi(id)
	var roles []*models.TUserRole
	orm := orm.NewOrm()
	qs := orm.QueryTable("t_user_role")
	_, err = qs.Filter("user_id", userId).All(&roles) // 过滤器
	if err != nil {
		c.ErrorJson("UserControllerByParId->Get", 500, err.Error(), nil)

	} else {
		c.SuccessJson("UserControllerByParId->Get", &roles)
	}
}

//delete userRole
func (c *UserRoleControllerByParId) Delete() {
	id := c.Ctx.Input.Param(":userId")
	userId, err := strconv.Atoi(id)
	o := orm.NewOrm()
	user := models.TUserRole{UserId: userId}
	_, err = o.Delete(&user)
	if err != nil {
		c.ErrorJson("UserRoleControllerByParId->Delete", 500, err.Error(), nil)
	} else {
		c.SuccessJson("UserRoleControllerByParId->Delete", nil)
	}
}
