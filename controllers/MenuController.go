package controllers

import (
	"easebase_web/models"
	"easebase_web/utils"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

type MenuController struct {
	BaseController
}

//query tree
func (c *MenuController) Get() {
	name := c.GetString("name")
	o := orm.NewOrm()
	var menu []orm.Params
	st := fmt.Sprintf(
		`select 
					  a.id               AS Id,
					  concat(repeat('&nbsp;',(length(id)-2)*6),a.name) AS Name,
					  a.url              AS Url,
                      a.icon             AS Icon,
					  a.creator          AS Creator,
					  a.create_date      AS CreateDate,
					  a.updater          AS Updater,
					  a.last_update_date AS LastUpdateDate,
					  case a.status when '1' THEN '启用' when '0' THEN '禁用' end AS Status
			     from t_xtqx a where  id<>'0' and  instr(a.name,'%s')>0  ORDER BY a.id`, name)
	_, err := o.Raw(st).Values(&menu)
	if err == nil {
		c.SuccessJson("MenuController->Get", &menu)
	} else {
		c.ErrorJson("MenuController->Get", 500, err.Error(), nil)
	}
}

//insert tree
func (c *MenuController) Put() {
	name := c.GetString("name")
	url := c.GetString("url")
	status := c.GetString("status")
	icon := c.GetString("icon")
	parent := c.GetString("parent")
	id := utils.GetMenuId(parent)
	fmt.Println("Id:", id)

	o := orm.NewOrm()
	m := models.TXtqx{
		Id:       string(id),
		Name:     name,
		Url:      url,
		Status:   status,
		Icon:     icon,
		ParentId: parent,
		Creator:  "DBA",
		Updater:  "DBA",
	}
	_, err := o.Insert(&m)
	fmt.Println(id, err)
	if err != nil {
		c.ErrorJson("MenuController->Put", 500, err.Error(), nil)
	} else {
		c.SuccessJson("MenuController->Put", "id="+id)
	}
}

//update tree
func (c *MenuController) Post() {
	id := c.GetString("id")
	name := c.GetString("name")
	url := c.GetString("url")
	icon := c.GetString("icon")
	status := c.GetString("status")
	parent := c.GetString("parent")
	o := orm.NewOrm()
	user := models.TXtqx{
		Id:       id,
		ParentId: parent,
		Name:     name,
		Url:      url,
		Icon:     icon,
		Status:   status,
		Creator:  "DBA",
		Updater:  "DBA",
	}
	_, err := o.Update(&user)
	if err != nil {
		c.ErrorJson("更新菜单出错!", 500, err.Error(), nil)
	} else {
		c.SuccessJson("菜单变更成功!", "")
	}

}

//delete tree
func (c *MenuController) Delete() {
	id := c.GetString("id")
	o := orm.NewOrm()
	qx := models.TXtqx{Id: id}
	_, err := o.Delete(&qx)
	if err != nil {
		c.ErrorJson("MenuController->Delete", 500, err.Error(), nil)
	} else {
		c.SuccessJson("MenuController->Delete", nil)
	}
}

//MenuParentController Controller
type MenuParentController struct {
	BaseController
}

func (c *MenuParentController) Get() {
	o := orm.NewOrm()
	var menu []orm.Params
	st := fmt.Sprintf(`select cast(id as char) as id,name from t_xtqx where status='1' and url ='' order by id`)
	_, err := o.Raw(st).Values(&menu)
	if err == nil {
		c.SuccessJson("MenuParentController->Get", &menu)
	} else {
		c.ErrorJson("MenuParentController->Get", 500, err.Error(), nil)
	}
}

//MenuControllerByParId Controller
type MenuControllerByParId struct {
	BaseController
}

//query menu
func (c *MenuControllerByParId) Get() {
	id := c.Ctx.Input.Param(":id")
	o := orm.NewOrm()
	qx := models.TXtqx{Id: id}
	err := o.Read(&qx)
	if err != nil {
		c.ErrorJson("MenuControllerByParId->Get", 500, err.Error(), nil)

	} else {
		c.SuccessJson("MenuControllerByParId->Get", &qx)
	}
}

//MenuParentController Controller
type MenuQxController struct {
	BaseController
}

func (c *MenuQxController) Get() {
	o := orm.NewOrm()
	var qx []orm.Params
	st := fmt.Sprintf(`select cast(a.id as char) as 'value',
							          concat((select name from t_xtqx b where b.id=a.parent_id),'=>',name) AS 'desc' 
					   from t_xtqx a where a.status='1' AND a.url !=''`)
	_, err := o.Raw(st).Values(&qx)
	if err == nil {
		c.SuccessJson("MenuQxController->Get", &qx)
	} else {
		c.ErrorJson("MenuQxController->Get", 500, err.Error(), nil)
	}
}

//MenuParentController Controller
type MenuRoleQxController struct {
	BaseController
}

func (c *MenuRoleQxController) Get() {
	//roleId := c.GetString("role_id")
	o := orm.NewOrm()
	var qx []orm.Params
	st := fmt.Sprintf(
		`select cast(a.id as char) as 'value',
					   concat((select name from t_xtqx b where b.id=a.parent_id),'=>',name) AS 'desc' 
				   from t_xtqx a 
				   where a.status='1' AND a.url !=''
                   `)
	_, err := o.Raw(st).Values(&qx)
	if err == nil {
		c.SuccessJson("MenuQxController->Get", &qx)
	} else {
		c.ErrorJson("MenuQxController->Get", 500, err.Error(), nil)
	}
}
