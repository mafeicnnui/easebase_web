package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

type DmController struct {
	BaseController
}

//query dm by dmm
func (c *DmController) Post() {
	dm := c.GetString("dm")
	o := orm.NewOrm()
	var dmm []orm.Params
	st := fmt.Sprintf("select dmm,dmmc from t_dmmx where dm='%s' order by dmm", dm)
	_, err := o.Raw(st).Values(&dmm)
	if err == nil {
		c.SuccessJson("DmController->Get", &dmm)
	} else {
		c.ErrorJson("DmController->Get", 500, err.Error(), nil)
	}
}
