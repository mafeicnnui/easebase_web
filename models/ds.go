package models

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

// Model Struct
type TDbSource struct {
	Id             int
	Ip             string    `orm:"size(100);description(数据库IP)"`
	Port           string    `orm:"size(10);description(数据库端口)"`
	Service        string    `orm:"size(50);description(数据库名称)"`
	Status         string    `orm:"size(50);description(数据源状态)"`
	User           string    `orm:"size(20);description(用户名)"`
	Password       string    `orm:"size(100);description(密码)"`
	DbType         string    `orm:"size(1);description(数据库类型)"`
	DbDesc         string    `orm:"size(50);description(数据源描述)"`
	DbEnv          string    `orm:"size(1);description(数据库环境)"`
	InstType       string    `orm:"size(10);description(实例类型)"`
	ProxyStatus    string    `orm:"size(1);description(代理状态)"`
	ProxyServer    string    `orm:"size(50);description(代理服务接口)"`
	MarketId       string    `orm:"size(50);description(项目编码)"`
	Creator        string    `orm:"size(10);description(主键)"`
	Updater        string    `orm:"size(10);description(主键)"`
	CreateDate     time.Time `orm:"auto_now_add;type(datetime) ;description(创建时间)"`
	LastUpdateDate time.Time `orm:"auto_now;type(datetime) ;description(修改时间)"`
}
