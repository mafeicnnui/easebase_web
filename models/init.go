package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	// set default database
	user, _ := web.AppConfig.String("user")
	password, _ := web.AppConfig.String("password")
	url, _ := web.AppConfig.String("url")
	db, _ := web.AppConfig.String("db")
	charset, _ := web.AppConfig.String("charset")
	timezone, _ := web.AppConfig.String("timezone")
	ds := user + ":" + password + "@tcp(" + url + ")/" + db + "?charset=" + charset + "&loc=" + timezone
	orm.RegisterDataBase("default", "mysql", ds)

	// register model
	orm.RegisterModel(new(TUser))
	orm.RegisterModel(new(TXtqx))
	orm.RegisterModel(new(TUserRole))
	orm.RegisterModel(new(TRole))
	orm.RegisterModel(new(TRolePrivs))
	orm.RegisterModel(new(TDbSource))
	orm.RegisterModel(new(TServer))
	orm.RegisterModel(new(TDbBackupConfig))
	orm.RegisterModel(new(TDbBackupTotal))
	orm.RegisterModel(new(TDbBackupDetail))

	// create table
	orm.RunSyncdb("default", false, true)
}
