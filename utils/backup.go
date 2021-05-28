package utils

import (
	"easebase_web/models"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"time"
)

type ReturnMsg struct {
	Code int
	Msg  string
}

//检测备份服务器状态
func CheckBackupServerStatus(dbTag string) int {
	o := orm.NewOrm()
	var res []orm.Params
	st := fmt.Sprintf(`select count(0) as rec from t_db_backup_config a,t_server b where a.server_id=b.id and a.db_tag='%s' and b.status='0'`, dbTag)
	fmt.Println(st)
	_, err1 := o.Raw(st).Values(&res)
	if err1 != nil {
		panic("Func CheckServerBackupStatus Error:" + err1.Error())
	}
	val, err2 := strconv.Atoi(res[0]["rec"].(string))
	if err2 != nil {
		panic("Func CheckServerBackupStatus Error:" + err2.Error())
	}
	return val
}

//检测备份标识
func CheckBackupDbTag(dbTag string) int {
	o := orm.NewOrm()
	var res []orm.Params
	st := fmt.Sprintf(`select count(0) as rec from t_db_backup_config where db_tag='%s'`, dbTag)
	fmt.Println(st)
	_, err1 := o.Raw(st).Values(&res)
	if err1 != nil {
		panic("Func CheckBackupDbTag Error:" + err1.Error())
	}
	val, err2 := strconv.Atoi(res[0]["rec"].(string))
	if err2 != nil {
		panic("Func CheckBackupDbTag Error:" + err2.Error())
	}
	return val
}

//检测任务是否禁用
func CheckBackupTaskStatus(dbTag string) int {
	o := orm.NewOrm()
	var res []orm.Params
	st := fmt.Sprintf(`select count(0) as rec from t_db_backup_config a,t_server b where a.server_id=b.id and a.db_tag='%s' and a.status='0'`, dbTag)
	fmt.Println(st)
	_, err1 := o.Raw(st).Values(&res)
	if err1 != nil {
		panic("Func CheckBackupTaskStatus Error:" + err1.Error())
	}
	val, err2 := strconv.Atoi(res[0]["rec"].(string))
	if err2 != nil {
		panic("Func CheckBackupTaskStatus Error:" + err2.Error())
	}
	return val
}

//获取备份配置
func GetBackupConfig(dbTag string) ([]orm.Params, error) {
	o := orm.NewOrm()
	var cfg []orm.Params

	if CheckBackupServerStatus(dbTag) > 0 {
		return nil, errors.New("备份服务器已禁用")
	}

	if CheckBackupDbTag(dbTag) == 0 {
		return nil, errors.New(fmt.Sprintf(`备份标识:'%s'不存在`, dbTag))
	}

	if CheckBackupTaskStatus(dbTag) > 0 {
		return nil, errors.New("备份任务已禁用")
	}

	st := fmt.Sprintf(
		`select      
					a.db_tag,
					c.ip       AS db_ip,
					c.port     AS db_port,
					c.service  AS db_service,
					c.user     AS db_user,
					c.password AS db_pass,
					a.expire,a.bk_base,a.script_path,a.script_file,a.bk_cmd,a.run_time,
					b.server_ip,b.server_port,b.server_user,b.server_pass,b.server_os,
					a.comments,a.backup_databases,a.api_server,a.status
			from t_db_backup_config a,t_server b,t_db_source c
			where a.server_id=b.id AND a.db_id=c.id AND a.db_tag='%s' AND b.status='1'`, dbTag)
	_, err := o.Raw(st).Values(&cfg)
	return cfg, err
}

//写备份汇总日志
func SaveBackupTotal(cfg map[string]string) ReturnMsg {
	fmt.Println("SaveBackupTotal=>config:", cfg)
	// parse datetime column
	template1 := "2006-01-02 15:04:05"
	template2 := "2006-01-02"
	StartTime, _ := time.Parse(template1, cfg["start_time"])
	EndTime, _ := time.Parse(template1, cfg["end_time"])
	CreateDate, _ := time.Parse(template2, cfg["create_date"])
	fmt.Println("CreateDate=", CreateDate)

	// check tab exists
	vv := fmt.Sprintf(` where db_tag='%s' and create_date='%s'`, cfg["db_tag"], cfg["create_date"])
	fmt.Println("SaveBackupTotal=", vv)
	o := orm.NewOrm()
	if CheckTabExists("t_db_backup_total", vv) == 0 {
		total := models.TDbBackupTotal{
			DbTag:         cfg["db_tag"],
			BkBase:        cfg["bk_base"],
			TotalSize:     cfg["total_size"],
			StartTime:     StartTime,
			EndTime:       EndTime,
			ElapsedBackup: cfg["elapsed_backup"],
			ElapsedGzip:   cfg["elapsed_gzip"],
			CreateDate:    CreateDate,
			Status:        cfg["status"],
			Creator:       "DBA",
			Updater:       "DBA",
		}
		_, err := o.Insert(&total)
		if err != nil {
			return ReturnMsg{Code: -1, Msg: err.Error()}
		} else {
			return ReturnMsg{Code: 0, Msg: "insert success"}
		}
	} else {
		st := `update t_db_backup_total set 
					bk_base         = ?,
					total_size      = ?,
					start_time      = ?,
					end_time        = ?,
					elapsed_backup  = ?,
					elapsed_gzip    = ?,
                    status          = ?,
					create_date     = ?
				where db_tag = ? and create_date=?`

		_, err := o.Raw(st,
			cfg["bk_base"],
			cfg["total_size"],
			StartTime,
			EndTime,
			cfg["elapsed_backup"],
			cfg["elapsed_gzip"],
			cfg["status"],
			cfg["create_date"],
			cfg["db_tag"],
			cfg["create_date"]).Exec()
		if err != nil {
			return ReturnMsg{Code: -1, Msg: err.Error()}
		} else {
			return ReturnMsg{Code: 0, Msg: "update success"}
		}
	}

}

//写备份明细日志
func SaveBackupDetail(cfg map[string]string) ReturnMsg {
	fmt.Println("SaveBackupDetail=>config:", cfg)
	// parse datetime column
	template1 := "2006-01-02 15:04:05"
	template2 := "2006-01-02"
	StartTime, _ := time.Parse(template1, cfg["start_time"])
	EndTime, _ := time.Parse(template1, cfg["end_time"])
	CreateDate, _ := time.Parse(template2, cfg["create_date"])
	fmt.Println("CreateDate=", CreateDate)

	// check tab exists
	vv := fmt.Sprintf(` where db_tag='%s' and db_name='%s' and create_date='%s'`, cfg["db_tag"], cfg["db_name"], cfg["create_date"])
	fmt.Println("SaveBackupDetail=", vv)
	o := orm.NewOrm()
	if CheckTabExists("t_db_backup_detail", vv) == 0 {
		detail := models.TDbBackupDetail{
			DbTag:         cfg["db_tag"],
			DbName:        cfg["db_name"],
			FileName:      cfg["file_name"],
			BkPath:        cfg["bk_path"],
			DbSize:        cfg["db_size"],
			StartTime:     StartTime,
			EndTime:       EndTime,
			ElapsedBackup: cfg["elapsed_backup"],
			ElapsedGzip:   cfg["elapsed_gzip"],
			Status:        cfg["status"],
			Error:         cfg["error"],
			CreateDate:    CreateDate,
			Creator:       "DBA",
			Updater:       "DBA",
		}
		_, err := o.Insert(&detail)
		if err != nil {
			return ReturnMsg{Code: -1, Msg: err.Error()}
		} else {
			return ReturnMsg{Code: 0, Msg: "insert success"}
		}
	} else {
		st := `update t_db_backup_detail set 
					db_name         = ?,
                    file_name       = ?,
					bk_path         = ?,
					db_size         = ?,
					start_time      = ?,
					end_time        = ?,
					elapsed_backup  = ?,
					elapsed_gzip    = ?,
                    status          = ?,
                    error           = ?,
					create_date     = ?
				where db_tag = ? and create_date=?`

		_, err := o.Raw(st,
			cfg["db_name"],
			cfg["file_name"],
			cfg["bk_path"],
			cfg["db_size"],
			StartTime,
			EndTime,
			cfg["elapsed_backup"],
			cfg["elapsed_gzip"],
			cfg["status"],
			cfg["error"],
			CreateDate,
			cfg["db_tag"],
			cfg["create_date"]).Exec()
		if err != nil {
			return ReturnMsg{Code: -1, Msg: err.Error()}
		} else {
			return ReturnMsg{Code: 0, Msg: "update success"}
		}
	}

}
