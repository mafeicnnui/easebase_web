package models

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

// Model Struct
type TUser struct {
	Id             int
	Name           string    `orm:"size(20)"`
	EmpNo          string    `orm:"size(10)"`
	Gender         string    `orm:"size(1)"`
	Email          string    `orm:"size(50)"`
	Phone          string    `orm:"size(20)"`
	DeptNo         string    `orm:"size(1)"`
	ExpireDate     time.Time `orm:"type(date)"`
	Password       string    `orm:"size(100)"`
	Status         string    `orm:"size(1)"`
	LoginName      string    `orm:"size(20)"`
	Creator        string    `orm:"size(10)"`
	Updater        string    `orm:"size(10)"`
	CreateDate     time.Time `orm:"auto_now_add;type(datetime) ;description(创建时间)"`
	LastUpdateDate time.Time `orm:"auto_now;type(datetime) ;description(修改时间)"`
}
