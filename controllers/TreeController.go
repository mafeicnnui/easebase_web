package controllers

import (
	"github.com/beego/beego/v2/client/orm"

	"fmt"
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
	var parents []orm.Params
	var children []orm.Params
	var nodes []orm.Params
	_, err := o.Raw("select id,name,url,icon from t_xtqx where status='1' and parent_id='0' and id!=0").Values(&parents)
	for _, parent := range parents {
		st := fmt.Sprintf("select id,name, url,icon from t_xtqx where parent_id='%s'", parent["id"])
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
