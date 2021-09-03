package main

import (
	"easebase_web/agent/dbSyncer/utils"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
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

		//检查源表是否存在
		if utils.CheckTabExists("sourceDB", tabName) == 0 {
			panic(fmt.Sprintf(`DB: %s ,Table: %s not exists!`, cfg["db_sour_service"], tabName))
		}

		//检查源表是否有主键
		if utils.CheckTabPkExists("sourceDB", tabName) == 0 {
			panic(fmt.Sprintf(`DB: %s ,Table %s not exists primary key!`, cfg["db_sour_service"], tabName))
		}

		//检查目标表是否存在
		if utils.CheckTabExists("targetDB", tabName) > 0 {
			fmt.Println(fmt.Sprintf(`DB: %s ,Table: %s already exists,skip create table!`, cfg["db_dest_service"], tabName))
		} else {
			//获取表定义
			st := utils.GetTabDef(tabName)
			//目标库建表
			utils.CreateTable("targetDB", tabName, st)
		}
	}
}

func fullSync(cfg map[string]interface{}) {
	tabList := strings.Split(cfg["sync_table"].(string), ",")
	batchSize, _ := strconv.Atoi(cfg["batch_size"].(string))
	startTime := utils.GetTime()
	for _, tab := range tabList {
		//获取表名，列名，同步时长
		tabName := strings.Split(tab, ":")[0]
		colName := strings.Split(tab, ":")[1]
		syncTime, _ := strconv.Atoi(strings.Split(tab, ":")[2])
		if colName == "" {
			panic(fmt.Sprintf(`DB: %s ,Table: %s sync column not configuration!`, cfg["db_sour_service"], tabName))
		}

		//全量同步表
		if utils.CheckTabExists("targetDB", tabName) > 0 && utils.CheckTabDataExists("targetDB", tabName) == 0 {
			fmt.Println(fmt.Sprintf(`DB: %s ,Table:'%s' starting full sync,batch size:%s...`, cfg["db_dest_service"], tabName, batchSize))
			header := utils.GetTabHeader("sourceDB", tabName)
			cols := utils.GetSyncCols("sourceDB", tabName)
			total := utils.CheckTabDataExists("sourceDB", tabName)
			counter := 0

			//构造表同步条件
			minRq, days := utils.GetTabMinRqDays("sourceDB", tabName, colName)
			for i := 0; i <= days; i += syncTime {
				beginDate := fmt.Sprintf(`DATE_ADD('%s',INTERVAL %d DAY)`, minRq, i)
				endDate := fmt.Sprintf(`DATE_ADD('%s',INTERVAL %d DAY)`, minRq, i+syncTime)
				vWhere := fmt.Sprintf(`where %s between %s and %s `, colName, beginDate, endDate)

				//源库:获取表数据
				rs := utils.ExecSQL("sourceDB", fmt.Sprintf(`select %s from %s %s`, cols, tabName, vWhere))
				values := ""
				for _, r := range rs {
					cols := strings.Split(strings.Replace(cols, "`", "", -1), ",")
					value := ""
					for xh := range cols {
						if r[cols[xh]] == nil {
							value = value + "null,"
						} else if utils.IsDigit(r[cols[xh]].(string)) == true || utils.IsFloat(r[cols[xh]].(string)) == true {
							value = value + fmt.Sprintf(`%s,`, r[cols[xh]].(string))
						} else {
							value = value + fmt.Sprintf(`'%s',`, r[cols[xh]].(string))
						}
					}
					counter += 1
					values += fmt.Sprintf(`(%s),`, value[0:len(value)-1])
					//目标库：打包执行
					if counter%batchSize == 0 {
						fmt.Printf("\r%s",
							fmt.Sprintf(`Synce Table %s ,progress:%d/%d[%.2f%%],Time:%ds...`,
								tabName,
								total,
								counter,
								float64(counter)/float64(total)*100,
								utils.GetSeconds(startTime, utils.GetTime())))
						utils.ExecSQL("targetDB", fmt.Sprintf(`%s values %s`, header, values[0:len(values)-1]))
						values = ""
					}
				}
				//目标库：执行最后一批
				if len(values) > 0 {
					fmt.Printf("\r%s",
						fmt.Sprintf(`Synce Table %s ,progress:%d/%d[%.2f%%],Time:%ds...`,
							tabName,
							total,
							counter,
							float64(counter)/float64(total)*100,
							utils.GetSeconds(startTime, utils.GetTime())))
					utils.ExecSQL("targetDB", fmt.Sprintf(`%s values %s`, header, values[0:len(values)-1]))
				}
			}
			fmt.Println("\nsync complete!")
		} else {
			fmt.Println(fmt.Sprintf(`Table: %s already exists data,skip full sync!`, tabName))
		}
	}
}

func incrSync(cfg map[string]interface{}) {
	tabList := strings.Split(cfg["sync_table"].(string), ",")
	batchSize, _ := strconv.Atoi(cfg["batch_size_incr"].(string))
	startTime := utils.GetTime()
	for _, tab := range tabList {
		//获取表名，列名，同步时长
		tabName := strings.Split(tab, ":")[0]
		colName := strings.Split(tab, ":")[1]
		timeCol := cfg["sync_time_type"].(string)
		syncTime := strings.Split(tab, ":")[2]

		//增量同步表
		if utils.CheckTabExists("targetDB", tabName) > 0 && utils.CheckTabDataExists("targetDB", tabName) > 0 {
			counter := 0
			header := utils.GetTabHeader("sourceDB", tabName)
			cols := utils.GetSyncCols("sourceDB", tabName)
			vWhere := utils.GetTabIncrFilter(timeCol, colName, syncTime)
			vPkCols := utils.GetTabPkCols("sourceDB", tabName)
			vPkNames := utils.GetTabPkNames("sourceDB", tabName)

			//源库:获取表数据
			rs := utils.ExecSQL("sourceDB", fmt.Sprintf(`select %s , %s from %s %s`, vPkCols, cols, tabName, vWhere))
			total := len(rs)
			fmt.Print(fmt.Sprintf(`Table: %s Incr sync recent %s %s data,total %d rows!`, tabName, syncTime, timeCol, total))
			var values = ""
			var valuesD []string
			for _, r := range rs {
				cols := strings.Split(strings.Replace(cols, "`", "", -1), ",")
				vPkWhere := utils.GetTabPkWhere(vPkNames, r["pk"].(string))
				valuesD = append(valuesD, fmt.Sprintf(`delete from %s where %s`, tabName, vPkWhere))
				value := ""
				for xh := range cols {
					if r[cols[xh]] == nil {
						value = value + "null,"
					} else if utils.IsDigit(r[cols[xh]].(string)) == true || utils.IsFloat(r[cols[xh]].(string)) == true {
						value = value + fmt.Sprintf(`%s,`, r[cols[xh]].(string))
					} else {
						value = value + fmt.Sprintf(`'%s',`, r[cols[xh]].(string))
					}
				}
				counter += 1
				values += fmt.Sprintf(`(%s),`, value[0:len(value)-1])
				//目标库：打包执行
				if counter%batchSize == 0 {
					//删除表数据
					fmt.Printf("\n%s\n", fmt.Sprintf(`Delete Table %s data %d rows!`, tabName, len(valuesD)))
					for _, del := range valuesD {
						utils.ExecSQL("targetDB", del)
					}

					fmt.Printf("\r%s",
						fmt.Sprintf(`Increment Sync Table %s ,progress:%d/%d[%.2f%%],Time:%ds...`,
							tabName,
							total,
							counter,
							float64(counter)/float64(total)*100,
							utils.GetSeconds(startTime, utils.GetTime())))
					utils.ExecSQL("targetDB", fmt.Sprintf(`%s values %s`, header, values[0:len(values)-1]))
					values = ""
					valuesD = valuesD[0:0]
				}
			}
			//目标库：执行最后一批
			if len(values) > 0 {
				//删除数据
				fmt.Printf("\n%s\n", fmt.Sprintf(`Delete Table %s last batch data %d rows!`, tabName, len(valuesD)))
				for _, del := range valuesD {
					utils.ExecSQL("targetDB", del)
				}

				fmt.Printf("\r%s",
					fmt.Sprintf(`Increment Sync Table %s ,progress:%d/%d[%.2f%%],Time:%ds...`,
						tabName,
						total,
						counter,
						float64(counter)/float64(total)*100,
						utils.GetSeconds(startTime, utils.GetTime())))
				utils.ExecSQL("targetDB", fmt.Sprintf(`%s values %s`, header, values[0:len(values)-1]))
			}
			fmt.Println("\nsync complete!")
		}
	}
}

func main() {
	//全局配置
	var cfg map[string]interface{}

	//命令行获取tag参数 -sync_tag=sync_mysql_mysql_flow_218_Bi -api_server=10.16.46.121:9000
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

	//全量同步
	fullSync(cfg)

	//增量同步
	incrSync(cfg)

}
