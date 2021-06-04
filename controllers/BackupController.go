package controllers

import (
	"easebase_web/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

type BackupController struct {
	BaseController
}

type BackupControllerByParId struct {
	BaseController
}
type BackupServerController struct {
	BaseController
}

type BackupLogController struct {
	BaseController
}

type BackupTaskController struct {
	BaseController
}

//query backup
func (c *BackupController) Get() {
	dbTag := c.GetString("db_tag")
	dbEnv := c.GetString("db_env")
	dbType := c.GetString("db_type")

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

	o := orm.NewOrm()
	var backups []orm.Params
	st := fmt.Sprintf(
		`select
                  a.id,a.comments,a.db_tag,a.expire,a.run_time,
                  concat(b.server_ip,':',b.server_port),a.api_server,
                  CASE a.STATUS WHEN '1' THEN '启用' WHEN '0' THEN '禁用' END  as  status,
                  CASE a.task_status WHEN '1' THEN '<span style=''color: red''>运行中</span>' WHEN '0' THEN '停止' END as  task_status,
				  a.updater,
				  a.last_update_date      
              from t_db_backup_config a,t_server b,t_db_source c
              where a.server_id=b.id  AND a.db_id=c.id and b.status='1'  %s `, vWhere)
	fmt.Println(st)
	_, err := o.Raw(st).Values(&backups)
	if err == nil {
		c.SuccessJson("BackupController->Get", &backups)
	} else {
		c.ErrorJson("BackupController->Get", 500, err.Error(), nil)
	}

}

//insert backup
func (c *BackupController) Put() {
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
	id, err := o.Insert(&backup)
	fmt.Println(id, err)
	if err != nil {
		c.ErrorJson("BackupController->Put", 500, err.Error(), nil)
	} else {
		c.SuccessJson("BackupController->Put", "id="+strconv.FormatInt(id, 10))
	}
}

//update backup
func (c *BackupController) Post() {
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
func (c *BackupController) Delete() {
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

//query backup by id
func (c *BackupControllerByParId) Get() {
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

//get backup servers
func (c *BackupServerController) Get() {
	o := orm.NewOrm()
	var backups []orm.Params
	st := fmt.Sprintf(`SELECT '' AS dmm,'请选择...'  AS dmmc UNION ALL SELECT CONCAT(id,'') AS id,server_desc FROM t_server WHERE server_type='1' ORDER BY dmm+0`)
	fmt.Println(st)
	_, err := o.Raw(st).Values(&backups)
	if err == nil {
		c.SuccessJson("BackupServerController->Get", &backups)
	} else {
		fmt.Println(err.Error())
		c.ErrorJson("BackupServerController->Get", 500, err.Error(), nil)
	}
}

//query backup log
func (c *BackupLogController) Get() {
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

//query backup task
func (c *BackupTaskController) Get() {
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
