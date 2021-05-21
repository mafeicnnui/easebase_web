package models

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

// Model Struct
type TServer struct {
	Id             int
	MarketId       string    `orm:"size(100);description(项目编码)"`
	ServerType     string    `orm:"size(10);description(服务器类型)"`
	ServerDesc     string    `orm:"size(50);description(服务器描述)"`
	ServerAuth     string    `orm:"size(50);description(认证方式)"`
	ServerIp       string    `orm:"size(50);description(服务器IP)"`
	ServerPort     string    `orm:"size(20);description(服务器端口)"`
	ServerUser     string    `orm:"size(100);description(用户)"`
	ServerPass     string    `orm:"size(200);description(密码)"`
	ServerKeys     string    `orm:"size(1);description(密钥)"`
	ServerOs       string    `orm:"size(50);description(操作系统类型)"`
	ServerCpu      string    `orm:"size(20);description(cpu核数)"`
	ServerMem      string    `orm:"size(20);description(内存大小)"`
	Status         string    `orm:"size(50);description(服务器状态)"`
	Creator        string    `orm:"size(10);description(创建者)"`
	Updater        string    `orm:"size(10);description(更新者)"`
	CreateDate     time.Time `orm:"auto_now_add;type(datetime) ;description(创建时间)"`
	LastUpdateDate time.Time `orm:"auto_now;type(datetime) ;description(修改时间)"`
}
