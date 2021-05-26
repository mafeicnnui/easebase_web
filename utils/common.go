package utils

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

func CheckTabExists(pTab string, pWhere string) int {
	o := orm.NewOrm()
	var res []orm.Params
	st := fmt.Sprintf(`select count(0) as rec from %s  %s`, pTab, pWhere)
	fmt.Println(st)
	_, err1 := o.Raw(st).Values(&res)
	if err1 != nil {
		panic("Func CheckTabExists Error:" + err1.Error())
	}
	val, err2 := strconv.Atoi(res[0]["rec"].(string))
	if err2 != nil {
		panic("Func CheckTabExists Error:" + err2.Error())
	}
	return val
}
