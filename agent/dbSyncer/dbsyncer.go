package main

import (
	"easebase_web/agent/dbSyncer/utils"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func preProcess(cfg map[string]interface{}) map[string]interface{} {
	utils.InitSourDB(cfg)
	utils.InitTargetDB(cfg)
	return cfg
}

func syncDDL(cfg map[string]interface{}) {
	tabList := strings.Split(cfg["sync_table"].(string), ",")
	for _, tab := range tabList {
		//获取表名，列名，同步时长
		tabName := strings.Split(tab, ":")[0]
		//colName := strings.Split(tab,":")[1]
		//syncTime:= strings.Split(tab,":")[2]

		//检查源表是否存在
		if utils.CheckTabExists("sourceDB", tabName) == 0 {
			panic(fmt.Sprintf(`DB: %s , Table:'%s' not exists!`, cfg["db_sour_service"], tabName))
		}

		//检查源表是否有主键
		if utils.CheckTabPkExists("sourceDB", tabName) == 0 {
			panic(fmt.Sprintf(`DB: %s ,Table '%s' not exists primary key!`, cfg["db_sour_service"], tabName))
		}

		//检查目标表是否存在
		if utils.CheckTabExists("targetDB", tabName) > 0 {
			fmt.Println(fmt.Sprintf(`DB: %s ,Table:'%s' already exists,skip create table!`, cfg["db_dest_service"], tabName))
		} else {
			//获取表定义
			st := utils.GetTabDef(tabName)

			//目标库建表
			utils.CreateTable("targetDB", tabName, st)

		}

		//全量同步表
		if utils.CheckTabExists("targetDB", tabName) > 0 && utils.CheckTabDataExists("targetDB", tabName) == 0 {
			fmt.Println(fmt.Sprintf(`DB: %s , Table:'%s' starting full sync...`, cfg["db_dest_service"], tabName))

			header := utils.GetTabHeader("sourceDB", tabName)
			cols := utils.GetSyncCols("sourceDB", tabName)
			fmt.Println("header=", header)
			fmt.Println("cols=", cols)
			//mysql :=utils.MySQLExt{Config: cfg, Alias: "test"}
			//mysql.InitDB()
			//mysql.Exec("select now() as rq")

			//目标库
			//utils.ExecSQL("targetDB",fmt.Sprintf(`delete from  %s`,tabName))

			//获取表所有数据
			fmt.Println("Exec SQL:", fmt.Sprintf(`select %s from %s`, cols, tabName))
			rs := utils.ExecSQL("sourceDB", fmt.Sprintf(`select %s from %s`, cols, tabName))
			for _, r := range rs {
				fmt.Println(r)
			}
		}

	}
}

func main() {

	//全局配置
	var cfg map[string]interface{}

	//命令行获取tag参数
	//-sync_tag=sync_mysql_mysql_flow_218_Bi -api_server=10.16.47.114:9000 -config
	api := flag.String("api_server", "", "接口服务地址,格式：IP:PORT")
	tag := flag.String("sync_tag", "", "数据同步标识")
	show := flag.Bool("config", false, "是否显示配置信息")
	flag.Parse()

	//调用接口读取配置
	res := utils.GetConfig(*api, *tag)
	if res.Code == 200 {
		cfg = res.Data[0]
	} else {
		panic(res.Msg)
	}

	//预处理参数
	cfg = preProcess(cfg)

	//输出配置
	if *show {
		utils.ShowConfig(cfg)
	}

	//同步DDL
	syncDDL(cfg)

}
