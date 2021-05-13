package utils

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

func GetMenuId(parentId string) string {
	o := orm.NewOrm()
	var menu []orm.Params
	st := fmt.Sprintf(`select count(0) as num ,cast(concat('00',max(id)+1) as char) AS id from t_xtqx where parent_id='%s'`, parentId)
	_, err := o.Raw(st).Values(&menu)
	if err != nil {
		fmt.Println("GetMenuId->Get", err.Error())
	} else {
		fmt.Println("GetMenuId:", &menu)
	}
	num := menu[0]["num"]
	id := menu[0]["id"].(string)

	if num == 0 {
		return parentId + "01"
	} else {
		return id
	}
}
