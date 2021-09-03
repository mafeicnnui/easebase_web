package mysql2mysql

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strings"
)

type MySQL struct {
	ip       string
	port     string
	user     string
	password string
	db       string
	alias    string
}

func (m MySQL) InitDB() {
	ds := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", m.user, m.password, m.ip, m.port, m.db)
	orm.RegisterDataBase(m.alias, "mysql", ds)
}

func (m MySQL) Exec(sql string) {
	o := orm.NewOrmUsingDB(m.alias)
	var res []orm.Params
	_, err := o.Raw(sql).Values(&res)
	if err != nil {
		panic("Func MySQL=>Exec Error:" + err.Error())
	}
}

type MySQLExt struct {
	Config map[string]interface{}
	Alias  string
}

func (m MySQLExt) InitDB() {
	apiServer := m.Config["api_server"].(string)
	dbSlice := strings.Split(m.Config["sync_db_sour"].(string), ":")
	dbIp := dbSlice[0]
	dbPort := dbSlice[1]
	dbService := dbSlice[2]
	dbUser := dbSlice[3]
	dbPass := dbSlice[4]
	dbNewPass := GetPassword(apiServer, dbUser, dbPass)
	ds := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbNewPass, dbIp, dbPort, dbService)
	orm.RegisterDataBase(m.Alias, "mysql", ds)
}

func (m MySQLExt) Exec(sql string) {
	o := orm.NewOrmUsingDB(m.Alias)
	var res []orm.Params
	_, err := o.Raw(sql).Values(&res)
	if err != nil {
		panic("Func MySQLExt=>Exec Error:" + err.Error())
	}
	fmt.Println(res)
}
