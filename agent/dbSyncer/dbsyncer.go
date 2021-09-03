package main

import (
	"easebase_web/agent/dbSyncer/mysql2mysql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
)

//示例 ： -type=mysql2mysql -api=10.16.46.121:9000 -tag=sync_mysql_mysql_flow_218_Bi

func main() {

	//解析参数
	api := flag.String("api", "", "接口服务地址,格式：IP:PORT")
	tye := flag.String("type", "", "同步类型[支持:mysql2mysql]")
	tag := flag.String("tag", "", "数据同步标识")
	show := flag.Bool("config", false, "是否显示配置信息")
	flag.Parse()

	//开启同步
	if *tye == "mysql2mysql" {
		mysql2mysql.Sync(api, tag, show)
	}

}
