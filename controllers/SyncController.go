package controllers

import (
	"easebase_web/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

type SyncController struct {
	BaseController
}

type SyncControllerByParId struct {
	BaseController
}
type SyncServerController struct {
	BaseController
}

type SyncLogController struct {
	BaseController
}

type SyncTaskController struct {
	BaseController
}

//query sync
func (c *SyncController) Get() {
	syncTag := c.GetString("sync_tag")
	marketId := c.GetString("market_id")
	syncYwlx := c.GetString("sync_ywlx")
	syncType := c.GetString("sync_type")
	status := c.GetString("status")

	vWhere := ""
	if syncTag != "" {
		st := fmt.Sprintf(` and instr(a.sync_tag,'%s')>0`, syncTag)
		vWhere = vWhere + st
	}
	if marketId != "" {
		st := fmt.Sprintf(` and instr(a.sync_col_val,'%s')>0`, marketId)
		vWhere = vWhere + st
	}
	if syncYwlx != "" {
		st := fmt.Sprintf(` and a.sync_ywlx='%s'`, syncYwlx)
		vWhere = vWhere + st
	}
	if syncType != "" {
		st := fmt.Sprintf(` and a.sync_type='%s'`, syncType)
		vWhere = vWhere + st
	}
	if status != "" {
		st := fmt.Sprintf(` and a.status='%s'`, status)
		vWhere = vWhere + st
	}

	o := orm.NewOrm()
	var sync []orm.Params
	st := fmt.Sprintf(
		`SELECT  a.id,
                     concat(substr(a.sync_tag,1,40),'...') as sync_tag_,             
                     a.sync_tag,
                     concat(substr(a.comments,1,30),'...') as comments,
                     CONCAT(b.server_ip,':',b.server_port) AS sync_server,
                     c.dmmc AS  sync_ywlx,
                     a.run_time,
                     CASE WHEN INSTR(api_server,',')>0 THEN 
                        SUBSTR(a.api_server,1,INSTR(a.api_server,',')-1)
                     ELSE                         
                        a.api_server
                     END AS api_server ,
                     a.api_server as back_api_server,
                     CASE a.status WHEN '1' THEN '启用' WHEN '0' THEN '禁用' END  status,
                     CASE a.task_status 
                       WHEN '1' THEN 
                     CONCAT('<span style=''color: red''>运行中(',TIMESTAMPDIFF(SECOND,a.task_create_time,NOW()),'s)</span>')
                       WHEN '0' THEN '停止' END  task_status
             FROM t_db_sync_config a,t_server b ,t_dmmx c,t_dmmx d
            WHERE a.server_id=b.id AND b.status='1' 
              AND c.dm='08' AND d.dm='09'
              AND a.sync_ywlx=c.dmm
              AND a.sync_type=d.dmm %s`, vWhere)
	_, err := o.Raw(st).Values(&sync)
	if err == nil {
		c.SuccessJson("SyncController->Get", &sync)
	} else {
		c.ErrorJson("SyncController->Get", 500, err.Error(), nil)
	}

}

//insert sync
func (c *SyncController) Put() {
	ServerId, _ := c.GetInt("server_id")
	SourDbId, _ := c.GetInt("sour_db_id")
	DestDbId, _ := c.GetInt("dest_db_id")
	SyncTag := c.GetString("sync_tag")
	SyncYwlx := c.GetString("sync_ywlx")
	SyncType := c.GetString("sync_type")
	ScriptPath := c.GetString("script_path")
	ScriptFile := c.GetString("script_file")
	RunTime := c.GetString("run_time")
	Comments := c.GetString("comments")
	SyncSchema := c.GetString("sync_schema")
	SyncSchemaDest := c.GetString("sync_schema_dest")
	SyncTable := c.GetString("sync_table")
	BatchSize, _ := c.GetInt("batch_size")
	BatchSizeIncr, _ := c.GetInt("batch_size_incr")
	SyncGap, _ := c.GetInt("sync_gap")
	SyncColVal := c.GetString("sync_col_val")
	SyncColName := c.GetString("sync_col_name")
	SyncTimeType := c.GetString("sync_time_type")
	SyncRepairDay, _ := c.GetInt("sync_repair_day")
	ApiServer := c.GetString("api_server")
	Status := c.GetString("status")

	o := orm.NewOrm()
	sync := models.TDbSyncConfig{
		ServerId:       ServerId,
		SourDbId:       SourDbId,
		DestDbId:       DestDbId,
		SyncTag:        SyncTag,
		SyncYwlx:       SyncYwlx,
		SyncType:       SyncType,
		ScriptPath:     ScriptPath,
		ScriptFile:     ScriptFile,
		RunTime:        RunTime,
		Comments:       Comments,
		SyncSchema:     SyncSchema,
		SyncSchemaDest: SyncSchemaDest,
		SyncTable:      SyncTable,
		BatchSize:      BatchSize,
		BatchSizeIncr:  BatchSizeIncr,
		SyncGap:        SyncGap,
		SyncColVal:     SyncColVal,
		SyncColName:    SyncColName,
		SyncTimeType:   SyncTimeType,
		SyncRepairDay:  SyncRepairDay,
		ApiServer:      ApiServer,
		Status:         Status,
		Creator:        "DBA",
		Updater:        "DBA",
	}
	id, err := o.Insert(&sync)
	fmt.Println(id, err)
	if err != nil {
		c.ErrorJson("SyncController->Put", 500, err.Error(), nil)
	} else {
		c.SuccessJson("SyncController->Put", "id="+strconv.FormatInt(id, 10))
	}
}

//update sync
func (c *SyncController) Post() {
	BackupId, _ := c.GetInt("id")
	ServerId := c.GetString("server_id")
	DbId := c.GetString("db_id")
	DbType := c.GetString("db_type")
	DbTag := c.GetString("db_tag")
	Expire, _ := c.GetInt("expire")
	BkBase := c.GetString("bk_base")
	ScriptPath := c.GetString("script_path")
	ScriptFile := c.GetString("script_file")
	BkCmd := c.GetString("bk_cmd")
	RunTime := c.GetString("run_time")
	Comments := c.GetString("comments")
	BackupDatabases := c.GetString("backup_databases")
	ApiServer := c.GetString("api_server")
	Status := c.GetString("status")
	TaskStatus := c.GetString("task_status")

	o := orm.NewOrm()
	backup := models.TDbBackupConfig{
		Id:              BackupId,
		ServerId:        ServerId,
		DbId:            DbId,
		DbType:          DbType,
		DbTag:           DbTag,
		Expire:          Expire,
		BkBase:          BkBase,
		ScriptPath:      ScriptPath,
		ScriptFile:      ScriptFile,
		BkCmd:           BkCmd,
		RunTime:         RunTime,
		Comments:        Comments,
		BackupDatabases: BackupDatabases,
		ApiServer:       ApiServer,
		Status:          Status,
		TaskStatus:      TaskStatus,
		Creator:         "DBA",
		Updater:         "DBA",
	}
	_, err := o.Update(&backup)
	if err != nil {
		c.ErrorJson("更新备份元数据出错!", 500, err.Error(), nil)
	} else {
		c.SuccessJson("备份元数据变更成功!", "")
	}
}

//delete server
func (c *SyncController) Delete() {
	BackupId, _ := c.GetInt("id")
	o := orm.NewOrm()
	backup := models.TDbBackupConfig{Id: BackupId}
	_, err := o.Delete(&backup)
	if err != nil {
		c.ErrorJson("BackupController->Delete", 500, err.Error(), nil)
	} else {
		c.SuccessJson("BackupController->Delete", nil)
	}
}

//query sync by id
func (c *SyncControllerByParId) Get() {
	id := c.Ctx.Input.Param(":id")
	int32, err := strconv.Atoi(id)
	if err != nil {
		c.ErrorJson("BackupControllerByParId->Get", 500, err.Error(), nil)
	}
	o := orm.NewOrm()
	backup := models.TDbBackupConfig{Id: int32}
	err = o.Read(&backup)
	if err != nil {
		c.ErrorJson("BackupControllerByParId->Get", 500, err.Error(), nil)

	} else {
		c.SuccessJson("BackupControllerByParId->Get", &backup)
	}
}

//get sync servers
func (c *SyncServerController) Get() {
	o := orm.NewOrm()
	var backups []orm.Params
	st := fmt.Sprintf(`SELECT '' AS dmm,'请选择...'  AS dmmc UNION ALL SELECT CONCAT(id,'') AS id,server_desc FROM t_server WHERE server_type='2' ORDER BY dmm+0`)
	fmt.Println(st)
	_, err := o.Raw(st).Values(&backups)
	if err == nil {
		c.SuccessJson("BackupServerController->Get", &backups)
	} else {
		fmt.Println(err.Error())
		c.ErrorJson("BackupServerController->Get", 500, err.Error(), nil)
	}
}

//query sync log
func (c *SyncLogController) Get() {
	dbTag := c.GetString("db_tag")
	dbEnv := c.GetString("db_env")
	dbType := c.GetString("db_type")
	beginDate := c.GetString("begin_date")
	endDate := c.GetString("end_date")

	vWhere := ""
	if dbTag != "" {
		st := fmt.Sprintf(` and (instr(a.db_tag,'%s')>0 or instr(a.comments,'%s')>0)`, dbTag, dbTag)
		vWhere = vWhere + st
	}
	if dbEnv != "" {
		st := fmt.Sprintf(` and c.db_env='%s'`, dbEnv)
		vWhere = vWhere + st
	}
	if dbType != "" {
		st := fmt.Sprintf(` and a.db_type='%s'`, dbType)
		vWhere = vWhere + st
	}

	if beginDate != "" {
		st := fmt.Sprintf(` and b.create_date>='%s'`, beginDate)
		vWhere = vWhere + st
	}

	if endDate != "" {
		st := fmt.Sprintf(` and a.create_date<='%s'`, endDate)
		vWhere = vWhere + st
	}

	o := orm.NewOrm()
	var backups []orm.Params
	st := fmt.Sprintf(
		`SELECT  b.id, 
                        a.comments, 
                        b.db_tag,
						cast(b.create_date as char) as create_date,
						cast(b.start_time as char) as start_time,
						cast(b.end_time as char) as end_time,
						b.total_size,
                        b.elapsed_backup, 
                        b.elapsed_gzip,
						CASE b.STATUS WHEN '1' THEN '<span style=''color: red''>失败</span>' WHEN '0' THEN '成功' END  status
				FROM  t_db_backup_config a,t_db_backup_total b,t_db_source c
				WHERE a.db_tag=b.db_tag AND a.db_id=c.id  AND a.status='1' %s
				 order by b.create_date,b.db_tag`, vWhere)
	_, err := o.Raw(st).Values(&backups)
	if err == nil {
		c.SuccessJson("BackupLogController->Get", &backups)
	} else {
		c.ErrorJson("BackupLogController->Get", 500, err.Error(), nil)
	}

}

//query sync task
func (c *SyncTaskController) Get() {
	dbEnv := c.GetString("db_env")
	dbType := c.GetString("db_type")
	vWhere := ""

	if dbEnv != "" {
		st := fmt.Sprintf(` and c.db_env='%s'`, dbEnv)
		vWhere = vWhere + st
	}
	if dbType != "" {
		st := fmt.Sprintf(` and a.db_type='%s'`, dbType)
		vWhere = vWhere + st
	}

	o := orm.NewOrm()
	var backups []orm.Params
	st := fmt.Sprintf(
		`SELECT a.db_tag as dmm,a.comments as dmmc
				 FROM t_db_backup_config a ,t_server b,t_db_source c
				   WHERE a.STATUS=1 AND a.server_id=b.id AND a.db_id=c.id
				    %s  ORDER BY c.db_type,a.db_id`, vWhere)
	_, err := o.Raw(st).Values(&backups)
	if err == nil {
		fmt.Println("backup=", &backups)
		c.SuccessJson("BackupTaskController->Get", &backups)
	} else {
		c.ErrorJson("BackupTaskController->Get", 500, err.Error(), nil)
	}

}
