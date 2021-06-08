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

type SyncLogDetailController struct {
	BaseController
}

type SyncTaskController struct {
	BaseController
}

//query sync
func (c *SyncController) Get() {
	syncTag := c.GetString("syncTag")
	marketId := c.GetString("marketId")
	syncYwlx := c.GetString("syncYwlx")
	syncType := c.GetString("syncType")
	status := c.GetString("status")

	fmt.Println("syncTag=", syncTag)
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
	fmt.Println(st)
	_, err := o.Raw(st).Values(&sync)
	if err == nil {
		c.SuccessJson("SyncController->Get", &sync)
	} else {
		c.ErrorJson("SyncController->Get", 500, err.Error(), nil)
	}

}

//insert sync
func (c *SyncController) Put() {
	ServerId := c.GetString("server_id")
	SourDbId := c.GetString("sour_db_id")
	DestDbId := c.GetString("dest_db_id")
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
	SyncId, _ := c.GetInt("id")
	ServerId := c.GetString("server_id")
	SourDbId := c.GetString("sour_db_id")
	DestDbId := c.GetString("dest_db_id")
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
		Id:             SyncId,
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
	_, err := o.Update(&sync)
	if err != nil {
		c.ErrorJson("更新同步元数据出错!", 500, err.Error(), nil)
	} else {
		c.SuccessJson("同步元数据变更成功!", "")
	}
}

//delete server
func (c *SyncController) Delete() {
	SyncpId, _ := c.GetInt("id")
	o := orm.NewOrm()
	sync := models.TDbSyncConfig{Id: SyncpId}
	_, err := o.Delete(&sync)
	if err != nil {
		c.ErrorJson("SyncController->Delete", 500, err.Error(), nil)
	} else {
		c.SuccessJson("SyncController->Delete", nil)
	}
}

//query sync by id
func (c *SyncControllerByParId) Get() {
	id := c.Ctx.Input.Param(":id")
	int32, err := strconv.Atoi(id)
	if err != nil {
		c.ErrorJson("SyncControllerByParId->Get", 500, err.Error(), nil)
	}
	o := orm.NewOrm()
	sync := models.TDbSyncConfig{Id: int32}
	err = o.Read(&sync)
	if err != nil {
		c.ErrorJson("SyncControllerByParId->Get", 500, err.Error(), nil)

	} else {
		c.SuccessJson("SyncControllerByParId->Get", &sync)
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
	syncTag := c.GetString("sync_tag")
	marketId := c.GetString("market_id")
	syncYwlx := c.GetString("sync_ywlx")
	beginDate := c.GetString("begin_date")
	endDate := c.GetString("end_date")

	vWhere := ""
	if syncTag != "" {
		st := fmt.Sprintf(` and (instr(a.sync_tag,'%s')>0 or instr(a.comments,'%s')>0)`, syncTag, syncTag)
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

	if beginDate != "" {
		st := fmt.Sprintf(` and b.create_date>='%s'`, beginDate)
		vWhere = vWhere + st
	}

	if endDate != "" {
		st := fmt.Sprintf(` and a.create_date<='%s'`, endDate)
		vWhere = vWhere + st
	}

	o := orm.NewOrm()
	var logs []orm.Params
	st := fmt.Sprintf(
		`SELECT b.id,
                    c.dmmc as market_name,
                    a.comments,
                    b.sync_tag,
                    cast(b.create_date as char) as create_date,
                    b.duration,
                    b.amount
            FROM  t_db_sync_config a,t_db_sync_task_log b,t_dmmx c
            WHERE a.sync_tag=b.sync_tag 
              and c.dm='05' 
              and instr(a.sync_col_val,c.dmm)>0
              and a.status='1'  %s`, vWhere)
	fmt.Println(st)
	_, err := o.Raw(st).Values(&logs)
	if err == nil {
		c.SuccessJson("BackupLogController->Get", &logs)
	} else {
		c.ErrorJson("BackupLogController->Get", 500, err.Error(), nil)
	}
}

//query backup log detail
func (c *SyncLogDetailController) Get() {
	syncTag := c.GetString("sync_tag")
	createDate := c.GetString("create_date")

	vWhere := ""
	if syncTag != "" {
		st := fmt.Sprintf(` and a.sync_tag='%s'`, syncTag)
		vWhere = vWhere + st
	}

	if createDate != "" {
		st := fmt.Sprintf(` and b.create_date>='%s' and b.create_date<='%s'`, createDate[0:13], createDate)
		vWhere = vWhere + st
	}

	o := orm.NewOrm()
	var syncs []orm.Params
	st := fmt.Sprintf(
		`SELECT 
					 a.comments,
					 b.sync_tag,
					 b.sync_table,
					 CAST(b.create_date AS CHAR) as create_date, 
					 b.amount,
					 b.duration 
                FROM  t_db_sync_config a,t_db_sync_task_log_detail b
                WHERE  a.sync_tag=b.sync_tag 
                   AND a.status='1' %s`, vWhere)
	fmt.Println(st)
	_, err := o.Raw(st).Values(&syncs)
	if err == nil {
		c.SuccessJson("BackupLogDetailController->Get", &syncs)
	} else {
		c.ErrorJson("BackupLogDetailController->Get", 500, err.Error(), nil)
	}

}

//query sync task
func (c *SyncTaskController) Get() {
	marketId := c.GetString("market_id")
	o := orm.NewOrm()
	var sync []orm.Params
	st := fmt.Sprintf(
		`SELECT sync_tag as dmm,comments as dmmc
                   FROM t_db_sync_config  
                    WHERE STATUS=1  and INSTR(sync_col_val,'%s')>0  ORDER BY sync_col_val,comments`, marketId)
	_, err := o.Raw(st).Values(&sync)
	if err == nil {
		fmt.Println("sync=", &sync)
		c.SuccessJson("SyncTaskController->Get", &sync)
	} else {
		c.ErrorJson("SyncTaskController->Get", 500, err.Error(), nil)
	}
}
