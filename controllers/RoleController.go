package controllers

import (
	"easebase_web/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

type RoleController struct {
	BaseController
}

//query role
func (c *RoleController) Get() {
	name := c.GetString("name")
	o := orm.NewOrm()
	var roles []orm.Params
	st := fmt.Sprintf(
		`SELECT  a.id               AS Id,          
                        a.name             AS Name,
						a.creator          AS Creator,
						a.create_date      AS CreateDate,
						a.updater          AS Updater,
						a.last_update_date AS LastUpdateDate,
						CASE a.status WHEN '1' THEN '启用' WHEN '0' THEN '禁用' END AS Status
					FROM t_role a
					 WHERE  instr(a.name,'%s')>0  ORDER BY a.create_date`, name)

	_, err := o.Raw(st).Values(&roles)
	if err == nil {
		c.SuccessJson("RoleController->Get", &roles)
	} else {
		c.ErrorJson("RoleController->Get", 500, err.Error(), nil)
	}
}

//insert role
func (c *RoleController) Put() {
	name := c.GetString("name")
	status := c.GetString("status")
	rolePrivileges := make([]string, 0, 1000)
	c.Ctx.Input.Bind(&rolePrivileges, "role_privileges")

	o := orm.NewOrm()
	role := models.TRole{
		Name:    name,
		Status:  status,
		Creator: "DBA",
		Updater: "DBA",
	}
	RoleId, err := o.Insert(&role)
	if err != nil {
		c.ErrorJson("RoleController->Put", 500, err.Error(), nil)
	} else {
		for _, privilegeId := range rolePrivileges {
			rolePrivilege := models.TRolePrivs{
				RoleId:  int(RoleId),
				PrivId:  privilegeId,
				Creator: "DBA",
				Updater: "DBA",
			}
			_, err := o.Insert(&rolePrivilege)
			if err != nil {
				c.ErrorJson("插入角色权限出错!", 500, err.Error(), nil)
			}
		}
		c.SuccessJson("RoleController->Put", "id="+strconv.FormatInt(RoleId, 10))
	}
}

//update role
func (c *RoleController) Post() {
	roleId, _ := c.GetInt("id")
	name := c.GetString("name")
	status := c.GetString("status")
	rolePrivileges := make([]string, 0, 1000)
	c.Ctx.Input.Bind(&rolePrivileges, "role_privileges")
	fmt.Println("rolePrivileges=", rolePrivileges)

	o := orm.NewOrm()
	role := models.TRole{
		Id:      roleId,
		Name:    name,
		Status:  status,
		Creator: "DBA",
		Updater: "DBA",
	}

	_, err := o.Update(&role)
	if err != nil {
		c.ErrorJson("更新角色信息出错!", 500, err.Error(), nil)
	} else {
		//删除角色权限信息
		rolePrivilege := models.TRolePrivs{RoleId: roleId}
		_, err = o.Delete(&rolePrivilege)
		if err != nil {
			c.ErrorJson("删除角色权限出错!", 500, err.Error(), nil)
		}
		//插入变更后角色权限
		for _, privilegeId := range rolePrivileges {
			rolePrivilege = models.TRolePrivs{
				RoleId: roleId,
				PrivId: privilegeId,
			}
			_, err := o.Insert(&rolePrivilege)
			if err != nil {
				c.ErrorJson("插入角色权限出错!", 500, err.Error(), nil)
			}
		}
		c.SuccessJson("角色信息变更成功!", "")
	}
}

//delete role
func (c *RoleController) Delete() {
	id, _ := c.GetInt("role_id")
	o := orm.NewOrm()
	role := models.TRole{Id: id}
	_, err := o.Delete(&role)
	if err != nil {
		c.ErrorJson("RoleController->Delete", 500, err.Error(), nil)
	} else {
		rolePrivilege := models.TRolePrivs{RoleId: id}
		_, err := o.Delete(&rolePrivilege)
		if err != nil {
			c.ErrorJson("RolePrivilegesController->Delete", 500, err.Error(), nil)
		}
		c.SuccessJson("RolePrivilegesController->Delete", nil)
	}
}

//MenuControllerByParId Controller
type RoleControllerByParId struct {
	BaseController
}

//query role by id
func (c *RoleControllerByParId) Get() {
	id := c.Ctx.Input.Param(":id")
	roleId, _ := strconv.Atoi(id)
	o := orm.NewOrm()
	role := models.TRole{Id: roleId}
	err := o.Read(&role)
	fmt.Println("role=", role)
	if err != nil {
		c.ErrorJson("RoleControllerByParId->Get", 500, err.Error(), nil)
	} else {
		c.SuccessJson("RoleControllerByParId->Get", &role)
	}
}

//功能：角色权限控制器

type RolePrivilegesControllerByParId struct {
	BaseController
}

//query userRole
func (c *RolePrivilegesControllerByParId) Get() {
	roleId := c.Ctx.Input.Param(":roleId")
	o := orm.NewOrm()
	var roleQx []orm.Params
	st := fmt.Sprintf(
		`select cast(a.id as char) as 'value',
                           concat((select name from t_xtqx b where b.id=a.parent_id),'=>',name) AS 'desc' 
					   from t_xtqx a 
                         where a.status='1' AND a.url !='' 
                           and a.id in (select priv_id from t_role_privs WHERE role_id='%s')`, roleId)
	_, err := o.Raw(st).Values(&roleQx)
	if err != nil {
		c.ErrorJson("RolePrivilegesControllerByParId->Get", 500, err.Error(), nil)

	} else {
		c.SuccessJson("RolePrivilegesControllerByParId->Get", &roleQx)
	}
}
