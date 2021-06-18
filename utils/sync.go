package utils

import (
	"easebase_web/models"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"time"
)

//检测同步服务器状态
func CheckSyncServerStatus(syncTag string) int {
	o := orm.NewOrm()
	var res []orm.Params
	st := fmt.Sprintf(`select count(0) as rec from t_db_sync_config a,t_server b where a.server_id=b.id and a.sync_tag='%s' and b.status='0'`, syncTag)
	_, err1 := o.Raw(st).Values(&res)
	if err1 != nil {
		panic("Func CheckSyncServerStatus Error:" + err1.Error())
	}
	val, err2 := strconv.Atoi(res[0]["rec"].(string))
	if err2 != nil {
		panic("Func CheckSyncServerStatus Error:" + err2.Error())
	}
	return val
}

//检测同步标识
func CheckSyncDbTag(syncTag string) int {
	o := orm.NewOrm()
	var res []orm.Params
	st := fmt.Sprintf(`select count(0) as rec from t_db_sync_config where sync_tag='%s'`, syncTag)
	_, err1 := o.Raw(st).Values(&res)
	if err1 != nil {
		panic("Func CheckSyncDbTag Error:" + err1.Error())
	}
	val, err2 := strconv.Atoi(res[0]["rec"].(string))
	if err2 != nil {
		panic("Func CheckSyncDbTag Error:" + err2.Error())
	}
	return val
}

//检测任务是否禁用
func CheckSyncTaskStatus(syncTag string) int {
	o := orm.NewOrm()
	var res []orm.Params
	st := fmt.Sprintf(`select count(0) as rec from t_db_sync_config a,t_server b where a.server_id=b.id and a.sync_tag='%s' and a.status='0'`, syncTag)
	_, err1 := o.Raw(st).Values(&res)
	if err1 != nil {
		panic("Func CheckSyncTaskStatus Error:" + err1.Error())
	}
	val, err2 := strconv.Atoi(res[0]["rec"].(string))
	if err2 != nil {
		panic("Func CheckSyncTaskStatus Error:" + err2.Error())
	}
	return val
}

//获取同步配置
func GetSyncConfig(dbTag string) ([]orm.Params, error) {
	o := orm.NewOrm()
	var cfg []orm.Params

	if CheckSyncServerStatus(dbTag) > 0 {
		return nil, errors.New("同步服务器已禁用")
	}

	if CheckSyncDbTag(dbTag) == 0 {
		return nil, errors.New(fmt.Sprintf(`同步标识:'%s'不存在`, dbTag))
	}

	if CheckSyncTaskStatus(dbTag) > 0 {
		return nil, errors.New("同步任务已禁用")
	}

	st := fmt.Sprintf(
		`select a.sync_tag,a.sync_ywlx,
					(select dmmc from t_dmmx where dm='08' and dmm=a.sync_ywlx) as sync_ywlx_name,
					a.sync_type,
					(select dmmc from t_dmmx where dm='09' and dmm=a.sync_type) as sync_type_name,
					CASE WHEN c.service='' THEN 
				       CONCAT(c.ip,':',c.port,':',a.sync_schema,':',c.user,':',c.password)
					ELSE
					   CONCAT(c.ip,':',c.port,':',c.service,':',c.user,':',c.password)
					END AS sync_db_sour,                          
					CASE WHEN d.service='' THEN 
					   CONCAT(d.ip,':',d.port,':',IFNULL(a.sync_schema_dest,a.sync_schema),':',d.user,':',d.password)
					ELSE
					   CONCAT(d.ip,':',d.port,':',d.service,':',d.user,':',d.password)
					END AS sync_db_dest,                          
					a.server_id,b.server_desc,a.run_time,a.api_server,
					lower(a.sync_table) AS sync_table,a.batch_size,a.batch_size_incr,a.sync_gap,a.sync_col_name,a.sync_repair_day,
					a.sync_col_val,a.sync_time_type,a.script_path,a.script_file,a.comments,
					a.status,b.server_ip,b.server_port,b.server_user,b.server_pass,c.proxy_server,
					(select dmmc from t_dmmx where dm='36' and dmm='01') as proxy_local_port,
					(select key_value from t_sys_settings where key_code='send_server') as send_server,
					(select key_value from t_sys_settings where key_code='send_port') as send_port,
					(select key_value from t_sys_settings where key_code='sender') as sender,
					(select key_value from t_sys_settings where key_code='sendpass') as sendpass,
					(select key_value from t_sys_settings where key_code='receiver') as receiver
			from t_db_sync_config a,t_server b,t_db_source c,t_db_source d
			  where a.server_id=b.id AND a.sour_db_id=c.id  AND a.dest_db_id=d.id
				and a.sync_tag ='%s' ORDER BY a.id,a.sync_ywlx`, dbTag)
	_, err := o.Raw(st).Values(&cfg)
	return cfg, err
}

//写同步汇总日志
func SaveSyncTotal(cfg map[string]string) ReturnMsg {
	template2 := "2006-01-02"
	CreateDate, _ := time.Parse(template2, cfg["create_date"])
	o := orm.NewOrm()
	total := models.TDbSyncTaskLog{
		SyncTag:    cfg["sync_tag"],
		Duration:   cfg["duration"],
		Amount:     cfg["amount"],
		CreateDate: CreateDate,
		Creator:    "DBA",
	}
	_, err := o.Insert(&total)
	if err != nil {
		return ReturnMsg{Code: -1, Msg: err.Error()}
	} else {
		return ReturnMsg{Code: 0, Msg: "insert success"}
	}
}

//写同步明细日志
func SaveSyncDetail(cfg map[string]string) ReturnMsg {
	o := orm.NewOrm()
	template2 := "2006-01-02"
	CreateDate, _ := time.Parse(template2, cfg["create_date"])
	vv := fmt.Sprintf(` where sync_tag='%s' and tab_name='%s'`, cfg["db_tag"], cfg["db_name"])
	if CheckTabDataExists("t_db_sync_tab_config", vv) == 0 {
		detail := models.TDbSyncTabConfig{
			SyncTag:     cfg["sync_tag"],
			DbName:      cfg["db_name"],
			SchemaName:  cfg["schema_name"],
			TabName:     cfg["tab_name"],
			SyncCols:    cfg["sync_cols"],
			SyncIncrCol: cfg["sync_incr_col"],
			SyncTime:    cfg["sync_time"],
			Status:      cfg["status"],
			CreateDate:  CreateDate,
			Creator:     "DBA",
		}
		_, err := o.Insert(&detail)
		if err != nil {
			return ReturnMsg{Code: -1, Msg: err.Error()}
		} else {
			return ReturnMsg{Code: 0, Msg: "insert t_db_sync_tab_config success"}
		}
	}

	detail := models.TDbSyncTaskLogDetail{
		SyncTag:    cfg["sync_tag"],
		SyncTable:  cfg["sync_table"],
		Duration:   cfg["duration"],
		Amount:     cfg["amount"],
		CreateDate: CreateDate,
		Creator:    "DBA",
	}
	_, err := o.Insert(&detail)
	if err != nil {
		return ReturnMsg{Code: -1, Msg: err.Error()}
	} else {
		return ReturnMsg{Code: 0, Msg: "insert success"}
	}
}
