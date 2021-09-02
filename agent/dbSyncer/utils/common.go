package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//请求返回的数据格式
type result struct {
	Name string                   `json:"Name"`
	Code int                      `json:"Code"`
	Msg  string                   `json:"Msg"`
	Data []map[string]interface{} `json:"Data"`
}

type resultApi struct {
	Name string `json:"Name"`
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
	Data string `json:"Data"`
}

//测试函数
func Test() {
	//测试
	fmt.Println(strings.Repeat("-", 120))
	for _, v := range GetSourDatabases("dbo_summary_day") {
		fmt.Println(fmt.Sprintf(`%-40s%-40s`, v["table_schema"], v["table_name"]))
	}
}

//初始化源数据库
func InitSourDB(cfg map[string]interface{}) {
	apiServer := cfg["api_server"].(string)
	dbSlice := strings.Split(cfg["sync_db_sour"].(string), ":")
	dbIp := dbSlice[0]
	dbPort := dbSlice[1]
	dbService := dbSlice[2]
	dbUser := dbSlice[3]
	dbPass := dbSlice[4]
	dbNewPass := GetPassword(apiServer, dbUser, dbPass)
	cfg["db_sour_ip"] = dbIp
	cfg["db_sour_port"] = dbPort
	cfg["db_sour_service"] = dbService
	cfg["db_sour_user"] = dbUser
	cfg["db_sour_password"] = dbNewPass
	cfg["db_sour_string"] = dbIp + ":" + dbPort + "/" + dbService
	ds := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbNewPass, dbIp, dbPort, dbService)
	orm.RegisterDataBase("sourceDB", "mysql", ds)
}

//初始化目标数据库
func InitTargetDB(cfg map[string]interface{}) {
	apiServer := cfg["api_server"].(string)
	dbSlice := strings.Split(cfg["sync_db_dest"].(string), ":")
	dbIp := dbSlice[0]
	dbPort := dbSlice[1]
	dbService := dbSlice[2]
	dbUser := dbSlice[3]
	dbPass := dbSlice[4]
	dbNewPass := GetPassword(apiServer, dbUser, dbPass)
	cfg["db_dest_ip"] = dbIp
	cfg["db_dest_port"] = dbPort
	cfg["db_dest_service"] = dbService
	cfg["db_dest_user"] = dbUser
	cfg["db_dest_password"] = dbNewPass
	cfg["db_dest_string"] = dbIp + ":" + dbPort + "/" + dbService

	ds := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbNewPass, dbIp, dbPort, dbService)
	orm.RegisterDataBase("targetDB", "mysql", ds)
}

/*
   功能：获取当前时间字符串，格式:YYYY-MM-DD HH24:MI:SS
   入口:无
   出口:当前时间字符串
*/
func getTimeStr() string {
	return time.Now().Format("2006-01-02 03:04:05")[0:10]
}

/*
   功能：获取当前时间字符串，格式:YYYY-MM-DD
   入口:无
   出口:当前时间字符串
*/
func GetDate() string {
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	return year + month + day
}

/*
   功能：获取当前时间字符串
   入口:无
   出口:当前时间字符串
*/
func GetTimeString(t time.Time) string {
	return t.Format("2006-01-02 03:04:05")
}

/*
   功能：获取当前时间
   入口:无
   出口:当前时间
*/
func GetTime() time.Time {
	return time.Now()
}

/*
   功能：t2-t1时间差（秒）
   入口:t1,t2时间串
   出口:秒
*/
func GetSeconds(t1 time.Time, t2 time.Time) int {
	return int(t2.Sub(t1).Seconds())
}

/*
   功能：解密函数
   入口:API地址，用户名，加密串
   出口:解密串
*/
func GetPassword(apiServer string, dbUser string, dbPass string) string {
	var res resultApi
	url := fmt.Sprintf(`http://%s/api/public/decrypt`, apiServer)
	tag := map[string]string{
		"key":      dbUser,
		"password": dbPass,
	}
	data, err := json.Marshal(tag)
	if err != nil {
		panic(err)
	}
	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err2)
	}
	err3 := json.Unmarshal(body, &res)
	if err3 != nil {
		panic(err3)
	}
	return res.Data
}

/*
   功能：获取同步配置
   入口:API地址，同步标识号
   出口:同步配置
*/
func GetConfig(ApiServer string, SyncTag string) result {
	var res result
	url := fmt.Sprintf(`http://%s/api/sync/config/%s`, ApiServer, SyncTag)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &res)
	return res
}

/*
   功能：打印同步配置
   入口:同步配置
*/
func ShowConfig(cfg map[string]interface{}) {
	fmt.Println(strings.Repeat("-", 140))
	fmt.Println(fmt.Sprintf(`%-20s%-120s`, "配置项", "配置值"))
	fmt.Println(strings.Repeat("-", 140))
	var keys []string
	for k := range cfg {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(fmt.Sprintf(`%-20s  =  %-120s`, k, cfg[k]))
	}
	fmt.Println(strings.Repeat("-", 140))
}

/*
   功能：检测表在当前库下是否存在
   入口:表名
   出口:1 存在 ,0 不存在
*/
func CheckTabExists(pDs, pTab string) int {
	o := orm.NewOrmUsingDB(pDs)
	var res []orm.Params
	st := fmt.Sprintf(
		`SELECT count(0) as rec
				  FROM information_schema.tables
				 WHERE  table_schema=database() and table_name='%s' order by table_name`, pTab)
	_, err1 := o.Raw(st).Values(&res)
	if err1 != nil {
		panic("Func CheckTabExists Error:" + err1.Error())
	}
	val, err2 := strconv.Atoi(res[0]["rec"].(string))
	if err2 != nil {
		panic("Func CheckTabExists Error:" + err2.Error())
	}
	return val
}

/*
   功能：检测表是否存在主键
   入口:表名
   出口:1 存在 ,0 不存在
*/
func CheckTabPkExists(pDs string, pTab string) int {
	o := orm.NewOrmUsingDB(pDs)
	var res []orm.Params
	st := fmt.Sprintf(
		`select count(0) as rec from information_schema.columns
              where table_schema=database() and table_name='%s' and column_key='PRI'`, pTab)
	_, err1 := o.Raw(st).Values(&res)
	if err1 != nil {
		panic("Func CheckTabPkExists Error:" + err1.Error())
	}
	val, err2 := strconv.Atoi(res[0]["rec"].(string))
	if err2 != nil {
		panic("Func CheckTabPkExists Error:" + err2.Error())
	}
	return val
}

/*
   功能：检测表是否存在主数据
   入口:表名
   出口:1 存在 ,0 不存在
*/
func CheckTabDataExists(pDs string, pTab string) int {
	o := orm.NewOrmUsingDB(pDs)
	var res []orm.Params
	st := fmt.Sprintf(`select count(0) as rec from %s limit 1`, pTab)
	_, err1 := o.Raw(st).Values(&res)
	if err1 != nil {
		panic("Func CheckTabDataExists Error:" + err1.Error())
	}
	val, err2 := strconv.Atoi(res[0]["rec"].(string))
	if err2 != nil {
		panic("Func CheckTabDataExists Error:" + err2.Error())
	}
	return val
}

/*
   功能：获取mysql表定义
   入口:表名
   出口:表定义
*/
func GetTabDef(pTab string) string {
	o := orm.NewOrmUsingDB("sourceDB")
	var res []orm.Params
	st := fmt.Sprintf(`show create table %s`, pTab)
	_, err1 := o.Raw(st).Values(&res)
	if err1 != nil {
		panic("Func GetTabDef Error:" + err1.Error())
	}
	val := res[0]["Create Table"].(string)
	return val
}

func GetTargetDatabases(pdb string) []orm.Params {
	o := orm.NewOrmUsingDB("targetDB")
	var dbs []orm.Params
	st := fmt.Sprintf(
		`SELECT schema_name 
				  FROM information_schema.schemata 
				 WHERE schema_name not IN('information_schema','performance_schema','test','sys','mysql')
				  and instr(schema_name,'%s')>0`, pdb)
	_, err := o.Raw(st).Values(&dbs)
	if err != nil {
		panic(err.Error())
	}
	return dbs
}

func GetSourDatabases(tabName string) []orm.Params {
	o := orm.NewOrmUsingDB("sourceDB")
	var rs []orm.Params
	st := fmt.Sprintf(
		`SELECT table_schema,table_name  
				  FROM information_schema.tables
				 WHERE  table_schema=database() and table_name='%s' order by table_name`, tabName)
	_, err := o.Raw(st).Values(&rs)
	if err != nil {
		panic(err.Error())
	}
	return rs
}

/*
   功能：创建表
   入口:表名,创建语句
   出口:无
*/
func CreateTable(pDs string, tName string, sql string) {
	o := orm.NewOrmUsingDB(pDs)
	var rs []orm.Params
	_, err := o.Raw(sql).Values(&rs)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(fmt.Sprintf(`Table: %s createed`, tName))
	}
}

/*
   功能：执行SQL获取结果集
   入口:表名,创建语句
   出口:无
*/
func ExecSQL(pDs string, sql string) []orm.Params {
	o := orm.NewOrmUsingDB(pDs)
	var rs []orm.Params
	_, err := o.Raw(sql).Values(&rs)
	if err != nil {
		panic(err.Error())
	}
	return rs
}

func GetTabHeader(pDs string, pTab string) string {
	o := orm.NewOrmUsingDB(pDs)
	var rs []orm.Params
	st := fmt.Sprintf(`select group_concat(concat('[[',column_name,']]')) AS column_name 
                               from information_schema.columns 
                                where table_schema=database() and table_name='%s' 
                                 order by ordinal_position`, pTab)
	_, err := o.Raw(st).Values(&rs)
	if err != nil {
		panic(err.Error())
	}
	cols := rs[0]["column_name"].(string)
	cols = strings.Replace(cols, "[[", "`", -1)
	cols = strings.Replace(cols, "]]", "`", -1)
	cols = fmt.Sprintf("insert into %s(%s)", pTab, cols)
	//fmt.Println("cols=",cols)
	return cols
}

func GetSyncCols(pDs string, pTab string) string {
	o := orm.NewOrmUsingDB(pDs)
	var rs []orm.Params
	st := fmt.Sprintf(`select group_concat(concat('[[',column_name,']]')) AS column_name 
                               from information_schema.columns 
                                where table_schema=database() and table_name='%s' 
                                 order by ordinal_position`, pTab)
	_, err := o.Raw(st).Values(&rs)
	if err != nil {
		panic(err.Error())
	}
	cols := rs[0]["column_name"].(string)
	cols = strings.Replace(cols, "[[", "`", -1)
	cols = strings.Replace(cols, "]]", "`", -1)
	return cols
}

func GetTabMinRqDays(pDs string, pTab string, pCol string) (string, int) {
	o := orm.NewOrmUsingDB(pDs)
	var rs []orm.Params
	st := fmt.Sprintf(`select min(modifytime) AS min_rq,
                                     timestampdiff(day,min(%s),max(%s)) AS days 
                             from %s`, pCol, pCol, pTab)
	_, err := o.Raw(st).Values(&rs)
	if err != nil {
		panic(err.Error())
	}
	rq := rs[0]["min_rq"].(string)
	days, _ := strconv.Atoi(rs[0]["days"].(string))
	return rq, days
}

func IsDigit(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}

func IsFloat(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}
	return true
}
