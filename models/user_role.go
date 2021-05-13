package models

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type TUserRole struct {
	Id     int
	UserId int `description(用户ID)"`
	RoleId int `description(角色ID)"`
}
