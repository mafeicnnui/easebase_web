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

func GetUserByUserID(userId string) (orm.Params, error) {
	o := orm.NewOrm()
	var user []orm.Params
	st := fmt.Sprintf("SELECT * from t_user where id=%s", userId)
	_, err := o.Raw(st).Values(&user)
	fmt.Println("GetUserByUserID=", user)
	if err != nil {
		return nil, err
	} else {
		return user[0], err
	}
}

func GetUserByUserName(loginName string) (orm.Params, error) {
	o := orm.NewOrm()
	var user []orm.Params
	st := fmt.Sprintf("SELECT * from t_user where login_name='%s'", loginName)
	_, err := o.Raw(st).Values(&user)
	fmt.Println("GetUserByUserName=", user)
	if err != nil {
		return nil, err
	} else {
		return user[0], err
	}
}
