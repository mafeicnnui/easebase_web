package utils

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

func GetUserRoleIDByUserID(userId int) []string {
	var roleId []string
	o := orm.NewOrm()
	var userRoles []orm.Params
	st := fmt.Sprintf(`select id from t_user_role where user_id='%d'`, userId)
	_, err := o.Raw(st).Values(&userRoles)
	if err != nil {
		fmt.Println("GetUserRoleIDByUserID->error:", err.Error())
	}
	for _, userRole := range userRoles {
		roleId = append(roleId, userRole["id"].(string))
	}
	return roleId
}
