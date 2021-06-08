package models

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

// Model Struct
type TDbSyncConfig struct {
	Id             int       `orm:"description(主键)"`
	SyncTag        string    `orm:"size(100);description(同步标识)"`
	SyncYwlx       string    `orm:"size(3);description(同步业务类型)"`
	Comments       string    `orm:"size(100);description(任务描述)"`
	SyncType       string    `orm:"size(1);description(同步数据类型)"`
	SourDbId       string    `orm:"size(20);description(源数据源ID)"`
	DestDbId       string    `orm:"size(20);description(目标数据源ID)"`
	ServerId       string    `orm:"size(20);description(同步服务器ID)"`
	RunTime        string    `orm:"size(100);description(运行时间)"`
	SyncSchema     string    `orm:"size(100);description(源数据库名)"`
	SyncSchemaDest string    `orm:"size(100);description(目标数据库名)"`
	SyncTable      string    `orm:"size(2000);description(同步表列表)"`
	BatchSize      int       `orm:"description(全量批大小)"`
	BatchSizeIncr  int       `orm:"description(增量批大小)"`
	SyncGap        int       `orm:"description(同步间隔)"`
	SyncColVal     string    `orm:"size(100);description(同步新增列值)"`
	SyncColName    string    `orm:"size(100);description(同步新增列名)"`
	SyncRepairDay  int       `orm:"description(自动修复天数)"`
	SyncTimeType   string    `orm:"size(20);description(同步时间类型)"`
	ScriptPath     string    `orm:"size(200);description(同步路径)"`
	ScriptFile     string    `orm:"size(100);description(同步文件名)"`
	ApiServer      string    `orm:"size(200);description(API服务器)"`
	Status         string    `orm:"size(1);description(同步状态)"`
	TaskStatus     string    `orm:"size(1);description(任务状态)"`
	Creator        string    `orm:"size(10);description(同步表列表)"`
	TaskCreateTime time.Time `orm:"type(datetime) ;description(任务创建时间)"`
	Updater        string    `orm:"size(10);description(同步表列表)"`
	CreateDate     time.Time `orm:"auto_now_add;type(datetime) ;description(创建时间)"`
	LastUpdateDate time.Time `orm:"auto_now;type(datetime) ;description(修改时间)"`
}

type TDbSyncTabConfig struct {
	Id             int       `orm:"description(主键)"`
	SyncTag        string    `orm:"size(100);description(同步标识)"`
	DbName         string    `orm:"size(100);description(数据库名称)"`
	SchemaName     string    `orm:"size(100);description(数据模式名称)"`
	TabName        string    `orm:"size(100);description(表名称)"`
	SyncCols       string    `orm:"size(2000);description(同步列列表)"`
	SyncIncrCol    string    `orm:"size(100);description(增量同步列)"`
	SyncTime       string    `orm:"size(100);description(同步时间)"`
	Status         string    `orm:"size(1);description(同步状态)"`
	Creator        string    `orm:"size(10);description(同步表列表)"`
	Updater        string    `orm:"size(10);description(同步表列表)"`
	CreateDate     time.Time `orm:"auto_now_add;type(datetime) ;description(创建时间)"`
	LastUpdateDate time.Time `orm:"auto_now;type(datetime) ;description(修改时间)"`
}

type TDbSyncTaskLog struct {
	Id         int
	SyncTag    string    `orm:"size(100);description(同步标识)"`
	Duration   int       `orm:"description(同步时长(s))"`
	Amount     int       `orm:"description(同步数量)"`
	Creator    string    `orm:"size(10)"`
	CreateDate time.Time `orm:"type(datetime);description(创建时间)"`
}

type TDbSyncTaskLogDetail struct {
	Id         int
	SyncTag    string    `orm:"size(100);description(同步标识)"`
	SyncTable  string    `orm:"size(100);description(同步表)"`
	Duration   int       `orm:"description(同步时长(s))"`
	Amount     int       `orm:"description(同步数量)"`
	Creator    string    `orm:"size(10)"`
	CreateDate time.Time `orm:"type(datetime);description(创建时间)"`
}
