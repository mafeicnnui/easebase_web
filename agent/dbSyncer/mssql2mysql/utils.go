package mssql2mysql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strings"
)

/*
   功能：请求返回的数据格式
*/
type result struct {
	Name string                   `json:"Name"`
	Code int                      `json:"Code"`
	Msg  string                   `json:"Msg"`
	Data []map[string]interface{} `json:"Data"`
}

/*
   功能：请求返回的数据格式(适用部分API)
*/
type resultApi struct {
	Name string `json:"Name"`
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
	Data string `json:"Data"`
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
		if strings.Contains(k, "DB") {
			fmt.Println(fmt.Sprintf(`%-20s  =  %-120s`, k, reflect.TypeOf(cfg[k])))
		} else {
			fmt.Println(fmt.Sprintf(`%-20s  =  %-120s`, k, cfg[k]))
		}

	}
	fmt.Println(strings.Repeat("-", 140))
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
   功能：初始化源数据库
   入口：配置文件
   出口: 无
*/
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
	fmt.Println(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbNewPass, dbIp, dbPort, dbService))
	mssql := Mssql{dbIp, dbPort, dbUser, dbNewPass, dbService, nil}
	mssql.Connect()
	cfg["sourceDB"] = mssql
}

/*
   功能：初始化目标数据库
   入口：配置文件
   出口: 无
*/
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
