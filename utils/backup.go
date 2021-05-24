package utils

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

func GetBackupConfig(dbTag string) ([]orm.Params, error) {
	o := orm.NewOrm()
	var cfg []orm.Params
	st := fmt.Sprintf(
		`SELECT      
                    a.db_tag,
                    c.ip       AS db_ip,
                    c.port     AS db_port,
                    c.service  AS db_service,
                    c.user     AS db_user,
                    c.password AS db_pass,
                    a.expire,a.bk_base,a.script_path,a.script_file,a.bk_cmd,a.run_time,
                    b.server_ip,b.server_port,b.server_user,b.server_pass,b.server_os,
                    a.comments,a.backup_databases,a.api_server,a.status
            FROM t_db_backup_config a,t_server b,t_db_source c
            WHERE a.server_id=b.id AND a.db_id=c.id AND a.db_tag='%s' AND b.status='1'`, dbTag)
	_, err := o.Raw(st).Values(&cfg)
	return cfg, err
}
