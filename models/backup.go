package models

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

// Model Struct
type TDbBackupConfig struct {
	Id              int
	ServerId        string `orm:"size(20)"`
	DbId            string `orm:"size(10)"`
	DbType          string `orm:"size(1)"`
	DbTag           string `orm:"size(100)"`
	Expire          int
	BkBase          string    `orm:"size(100)"`
	ScriptPath      string    `orm:"size(200)"`
	ScriptFile      string    `orm:"size(100)"`
	BkCmd           string    `orm:"size(100)"`
	RunTime         string    `orm:"size(100)"`
	Comments        string    `orm:"size(100)"`
	BackupDatabases string    `orm:"size(1000)"`
	ApiServer       string    `orm:"size(200)"`
	Status          string    `orm:"size(1)"`
	TaskStatus      string    `orm:"size(1)"`
	Creator         string    `orm:"size(10)"`
	Updater         string    `orm:"size(10)"`
	CreateDate      time.Time `orm:"auto_now_add;type(datetime) ;description(创建时间)"`
	LastUpdateDate  time.Time `orm:"auto_now;type(datetime) ;description(修改时间)"`
}

type TDbBackupTotal struct {
	Id             int
	DbTag          string    `orm:"size(100)"`
	BkBase         string    `orm:"size(200)"`
	TotalSize      string    `orm:"size(50)"`
	StartTime      time.Time `orm:"auto_now_add;type(datetime) ;description(开始时间)"`
	EndTime        time.Time `orm:"auto_now_add;type(datetime) ;description(结束时间)"`
	ElapsedBackup  string    `orm:"size(100)"`
	ElapsedGzip    string    `orm:"size(1000)"`
	Status         string    `orm:"size(1)"`
	Creator        string    `orm:"size(10)"`
	Updater        string    `orm:"size(10)"`
	CreateDate     time.Time `orm:"type(date);description(创建时间)"`
	LastUpdateDate time.Time `orm:"auto_now;type(datetime) ;description(修改时间)"`
}

type TDbBackupDetail struct {
	Id             int
	DbTag          string    `orm:"size(100)"`
	DbName         string    `orm:"size(100)"`
	FileName       string    `orm:"size(200)"`
	BkPath         string    `orm:"size(200)"`
	DbSize         string    `orm:"size(50)"`
	StartTime      time.Time `orm:"auto_now_add;type(datetime) ;description(开始时间)"`
	EndTime        time.Time `orm:"auto_now_add;type(datetime) ;description(结束时间)"`
	ElapsedBackup  string    `orm:"size(100)"`
	ElapsedGzip    string    `orm:"size(1000)"`
	Status         string    `orm:"size(1)"`
	Error          string    `orm:"size(2000)"`
	Creator        string    `orm:"size(10)"`
	Updater        string    `orm:"size(10)"`
	CreateDate     time.Time `orm:"type(date);description(创建时间)"`
	LastUpdateDate time.Time `orm:"auto_now;type(datetime) ;description(修改时间)"`
}
