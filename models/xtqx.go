package models

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

// Model Struct
type TXtqx struct {
	Id             string    `orm:"pk"`
	Name           string    `orm:"size(20)"`
	ParentId       string    `orm:"size(10)"`
	Url            string    `orm:"size(100)"`
	Status         string    `orm:"size(1)"`
	Icon           string    `orm:"size(50)"`
	Creator        string    `orm:"size(10)"`
	Updater        string    `orm:"size(10)"`
	CreateDate     time.Time `orm:"auto_now_add;type(datetime) ;description(创建时间)"`
	LastUpdateDate time.Time `orm:"auto_now;type(datetime) ;description(修改时间)"`
}
