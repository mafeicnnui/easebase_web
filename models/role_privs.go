package models

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

// Model Struct
type TRolePrivs struct {
	Id             int
	RoleId         int
	PrivId         string    `orm:"size(11)"`
	Creator        string    `orm:"size(10)"`
	Updater        string    `orm:"size(10)"`
	CreateDate     time.Time `orm:"auto_now_add;type(datetime) ;description(创建时间)"`
	LastUpdateDate time.Time `orm:"auto_now;type(datetime) ;description(修改时间)"`
}
