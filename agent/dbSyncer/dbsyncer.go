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

		//检查表是否存在
		if utils.CheckTabExists(tabName) == 0 {
			panic(fmt.Sprintf(`Table:%s not exists`, tabName))
		}

		//检查表是否有主键
		if utils.CheckTabPkExists(tabName) == 0 {
			panic(fmt.Sprintf(`Table:%s not exists primary key`, tabName))
		}

		//获取表定义
		st := utils.GetTabDef(tabName)
		fmt.Println(st)

		//目标库建表

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

	//测试
	fmt.Println(strings.Repeat("-", 120))
	for _, v := range utils.GetSourDatabases("") {
		fmt.Println(fmt.Sprintf(`%-80s`, v["schema_name"]))
	}

	//同步DDL
	syncDDL(cfg)

}
