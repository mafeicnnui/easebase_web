package controllers

import (
	"easebase_web/utils"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/beego/beego/v2/server/web/context"
)

type TreeController struct {
	BaseController
}

//query tree
func (c *TreeController) Post() {
	o := orm.NewOrm()
	var parents []orm.Params
	var children []orm.Params
	var nodes []orm.Params
	_, err := o.Raw("select id,name,url,icon from t_xtqx where parent_id='0' and id!=0").Values(&parents)
	for _, parent := range parents {
		st := fmt.Sprintf("select id,name,concat(name,'^',SUBSTR(REPLACE(url,'/','_'),2)) as url,icon from t_xtqx where parent_id='%s'", parent["id"])
		//st := fmt.Sprintf("select id,name, url,icon from t_xtqx where parent_id='%s'", parent["id"])
		_, err = o.Raw(st).Values(&children)
		parent["children"] = children
		nodes = append(nodes, parent)
	}
	if err == nil {
		//fmt.Println(&nodes)
		c.SuccessJson("TreeController->Get", &nodes)
	} else {
		c.ErrorJson("TreeController->Get", 500, err.Error(), nil)
	}
}

type TreeController2 struct {
	BaseController
}

//query tree
func (c *TreeController2) Post() {
	o := orm.NewOrm()
	var user utils.User
	var parents []orm.Params
	var children []orm.Params
	var nodes []orm.Params
	//var user []orm.Params
	fmt.Println("header=>Authorization=", c.Ctx.Request.Header["Authorization"])
	token := c.Ctx.Request.Header["Authorization"][0]
	user, _ = utils.ParseToken(token)
	fmt.Println("TreeController2=>userInfo=", user, user.Id, user.Name)
	st := fmt.Sprintf(`select id,name,url,icon 
							   from t_xtqx
							   where status='1' and parent_id='0' and id!=0
								 and  id in (select distinct parent_id FROM t_xtqx 
											 where id IN(select b.priv_id 
														  from t_user_role a ,t_role_privs b,t_role c
														 where a.role_id=b.role_id                                          
														   and a.role_id=c.id AND c.status='1'
														   and a.user_id='%s')) order by id`, user.Id)
	//fmt.Println("parent=",st)
	_, err := o.Raw(st).Values(&parents)
	//if err == nil {
	//	c.SuccessJson("TreeController2->Post->Parent", &parents)
	//} else {
	//	c.ErrorJson("TreeController2->Post->Parent", 500, err.Error(), nil)
	//}
	for _, parent := range parents {
		st := fmt.Sprintf(`select id,name, url,icon from t_xtqx
                                  where parent_id='%s' and status='1'
                                   and  id IN(select b.priv_id 
											   from t_user_role a ,t_role_privs b,t_role c
											   where a.role_id=b.role_id
												 and a.role_id=c.id and c.status='1'
												 and a.user_id='%s') order by id`, parent["id"], user.Id)
		//fmt.Println("children=",st)
		_, err = o.Raw(st).Values(&children)
		//if err == nil {
		//	c.SuccessJson("TreeController2->Post->Children", &children)
		//} else {
		//	c.ErrorJson("TreeController2->Post->Children", 500, err.Error(), nil)
		//}
		parent["children"] = children
		nodes = append(nodes, parent)
	}
	if err == nil {
		//fmt.Println(&nodes)
		c.SuccessJson("TreeController->Get", &nodes)
	} else {
		c.ErrorJson("TreeController->Get", 500, err.Error(), nil)
	}
}
