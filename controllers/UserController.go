package controllers

import (
	"easebase_web/models"
	"easebase_web/utils"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"time"
)

type UserController struct {
	BaseController
}

//query user
func (c *UserController) Get() {
	name := c.GetString("name")
	o := orm.NewOrm()
	var users []orm.Params
	st := fmt.Sprintf(
		`SELECT  a.id,
					  a.emp_no           AS EmpNo,
					  a.name             AS Name,
					  a.login_name       AS LoginName,
					  a.phone            AS Phone,
					  c.dmmc             AS Gender,
					  a.email            AS Email,
					  b.dmmc             AS DeptNo,
					  a.expire_date      AS ExpireDate,
					  a.creator          AS Creator,
					  a.create_date      AS CreateDate,
					  a.updater          AS Updater,
					  a.last_update_date AS LastUpdateDate,
					  CASE a.status WHEN '1' THEN '启用' WHEN '0' THEN '禁用' END AS Status
				FROM t_user a,t_dmmx b ,t_dmmx c
				 WHERE b.dm = '01' AND c.dm ='04' AND  a.dept_no = b.dmm 	AND  a.gender = c.dmm
					AND  instr(a.name,'%s')>0  ORDER BY a.create_date`, name)

	_, err := o.Raw(st).Values(&users)
	if err == nil {
		c.SuccessJson("TreeController->Get", &users)
	} else {
		c.ErrorJson("TreeController->Get", 500, err.Error(), nil)
	}

}

//insert user
func (c *UserController) Put() {
	template := "2006-01-02"
	name := c.GetString("name")
	empNo := c.GetString("emp_no")
	gender := c.GetString("gender")
	email := c.GetString("email")
	deptNo := c.GetString("dept_no")
	expireDate := c.GetString("expire_date")
	DExpireDate, _ := time.Parse(template, expireDate)
	fmt.Println("expireDate1:", expireDate)
	fmt.Println("expireDate2:", DExpireDate)
	password := c.GetString("password")
	status := c.GetString("status")
	loginName := c.GetString("login_name")
	phone := c.GetString("phone")
	userRoles := make([]int, 0, 10)
	c.Ctx.Input.Bind(&userRoles, "user_role")
	fmt.Println("userRoles:", userRoles)

	o := orm.NewOrm()
	user := models.TUser{
		Name:       name,
		EmpNo:      empNo,
		Gender:     gender,
		Email:      email,
		Phone:      phone,
		DeptNo:     deptNo,
		ExpireDate: DExpireDate,
		Password:   password,
		Status:     status,
		LoginName:  loginName,
		Creator:    "DBA",
		Updater:    "DBA",
	}
	id, err := o.Insert(&user)
	fmt.Println(id, err)
	if err != nil {
		c.ErrorJson("UserController->Put", 500, err.Error(), nil)
	} else {
		for _, userRole := range userRoles {
			fmt.Println("userRole:", userRole)
			role := models.TUserRole{
				UserId: int(id),
				RoleId: userRole,
			}
			roleId, err := o.Insert(&role)
			fmt.Println("roleId:", roleId, "error:", err)
		}
		c.SuccessJson("UserController->Put", "id="+strconv.FormatInt(id, 10))
	}

}

//update user
func (c *UserController) Post() {
	userId, _ := c.GetInt("id")
	template := "2006-01-02"
	name := c.GetString("name")
	empNo := c.GetString("emp_no")
	gender := c.GetString("gender")
	email := c.GetString("email")
	deptNo := c.GetString("dept_no")
	expireDate := c.GetString("expire_date")
	DExpireDate, _ := time.Parse(template, expireDate)
	fmt.Println("expireDate1:", expireDate)
	fmt.Println("expireDate2:", DExpireDate)
	password := c.GetString("password")
	status := c.GetString("status")
	loginName := c.GetString("login_name")
	phone := c.GetString("phone")
	userRoles := make([]int, 0, 3)
	c.Ctx.Input.Bind(&userRoles, "user_role")
	fmt.Println("userRoles:", userRoles)

	o := orm.NewOrm()
	user := models.TUser{
		Id:         userId,
		Name:       name,
		EmpNo:      empNo,
		Gender:     gender,
		Email:      email,
		Phone:      phone,
		DeptNo:     deptNo,
		ExpireDate: DExpireDate,
		Password:   password,
		Status:     status,
		LoginName:  loginName,
		Creator:    "DBA",
		Updater:    "DBA",
	}
	_, err := o.Update(&user)
	if err != nil {
		c.ErrorJson("更新用户信息出错!", 500, err.Error(), nil)
	} else {
		//删除用户角色信息
		for _, userRoleId := range utils.GetUserRoleIDByUserID(userId) {
			urId, _ := strconv.Atoi(userRoleId)
			rolePrivilege := models.TUserRole{
				Id: urId,
			}
			_, err := o.Delete(&rolePrivilege)
			if err != nil {
				c.ErrorJson("插入角色权限出错!", 500, err.Error(), nil)
			}
		}

		//插入变更后角色
		for _, userRole := range userRoles {
			fmt.Println("userRole:", userRole)
			role := models.TUserRole{
				UserId: int(userId),
				RoleId: userRole,
			}
			_, err := o.Insert(&role)
			if err != nil {
				c.ErrorJson("插入用户角色出错!", 500, err.Error(), nil)
			}
		}
		c.SuccessJson("用户信息变更成功!", "")
	}

}

//delete user
func (c *UserController) Delete() {
	id, _ := c.GetInt("id")
	o := orm.NewOrm()
	user := models.TUser{Id: id}
	_, err := o.Delete(&user)
	if err != nil {
		c.ErrorJson("UserController->Delete", 500, err.Error(), nil)
	} else {
		for _, userRoleId := range utils.GetUserRoleIDByUserID(id) {
			urId, _ := strconv.Atoi(userRoleId)
			rolePrivilege := models.TUserRole{
				Id: urId,
			}
			_, err := o.Delete(&rolePrivilege)
			if err != nil {
				c.ErrorJson("插入角色权限出错!", 500, err.Error(), nil)
			}
		}
		c.SuccessJson("UserController->Delete", nil)
	}
}

type UserControllerManager struct {
	BaseController
}

//query user page
func (c *UserControllerManager) Get() {
	c.TplName = "query_user.html"
}
