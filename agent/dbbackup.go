package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"os"
	_ "os"
	"os/exec"
	"strconv"
	_ "strings"
	"time"
)

type result struct {
	Name string                   `json:"Name"`
	Code int                      `json:"Code"`
	Msg  string                   `json:"Msg"`
	Data []map[string]interface{} `json:"Data"`
}

func GetTag() string {
	var tag string
	for idx, args := range os.Args {
		if idx == 1 {
			tag = args
		}
	}
	return tag
}

func getDate() string {
	return time.Now().Format("2006-01-02 03:04:05")[0:10]
}

func getTimeString(t time.Time) string {
	return t.Format("2006-01-02 03:04:05")
}

func getTime() time.Time {
	return time.Now()
}

func getSeconds(t1 time.Time, t2 time.Time) int {
	return int(t2.Sub(t1).Seconds())
}

func InitDB(cfg map[string]interface{}) {
	ds := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", cfg["db_user"], cfg["new_pass"], cfg["db_ip"], cfg["db_port"])
	orm.RegisterDataBase("default", "mysql", ds)
}

func GetConfig(tag string) result {
	var res result
	url := fmt.Sprintf(`http://10.16.47.114:9000/api/backup/config/%s`, tag)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &res)
	return res
}

func getDatabases(pdb string) []orm.Params {
	o := orm.NewOrm()
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

func GetPassword(key interface{}, password interface{}) string {
	var res result
	url := `http://10.16.47.114:9000/api/public/decode`
	tag := map[string]interface{}{
		"key":      key,
		"password": password,
	}
	data, _ := json.Marshal(tag)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &res)
	return res.Data[0]["password"].(string)
}

func GetFIleContents(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func GetFIleSize(filename string) int {
	fi, err := os.Stat(filename)
	if err != nil {
		fmt.Print(err)
	}
	return int(fi.Size())
}

func preProcess(cfg map[string]interface{}) map[string]interface{} {
	cfg["bk_path"] = fmt.Sprintf(`%s/%s`, cfg["bk_base"], getDate())
	cfg["new_pass"] = GetPassword(cfg["db_user"], cfg["db_pass"])
	InitDB(cfg)
	return cfg
}

func WriteLogTotal(cfg map[string]string) result {
	// 请求URL
	url := fmt.Sprintf(`http://10.16.47.114:9000/api/backup/log/%s`, cfg["db_tag"])

	//备份汇总日志
	total := map[string]string{
		"db_tag":         cfg["db_tag"],
		"bk_base":        cfg["bk_base"],
		"total_size":     cfg["total_size"],
		"start_time":     cfg["start_time"],
		"end_time":       cfg["end_time"],
		"create_date":    cfg["create_date"],
		"elapsed_backup": cfg["elapsed_backup"],
		"elapsed_gzip":   cfg["elapsed_gzip"],
		"status":         cfg["status"],
		"creator":        "DBA",
		"updater":        "DBA",
	}

	data, _ := json.Marshal(total)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("body:",body)
	//fmt.Println(string(body))
	var res result
	_ = json.Unmarshal(body, &res)
	//fmt.Println("Msg=", res)
	return res
}

func WriteLogDetail(cfg map[string]string) result {
	// 请求URL
	url := fmt.Sprintf(`http://10.16.47.114:9000/api/backup/detail/%s`, cfg["db_tag"])

	//备份明细日志
	detail := map[string]interface{}{
		"db_tag":         cfg["db_tag"],
		"db_name":        cfg["db_name"],
		"file_name":      cfg["file_name"],
		"bk_path":        cfg["bk_path"],
		"db_size":        cfg["db_size"],
		"start_time":     cfg["start_time"],
		"end_time":       cfg["end_time"],
		"elapsed_backup": cfg["elapsed_backup"],
		"elapsed_gzip":   cfg["elapsed_gzip"],
		"status":         cfg["status"],
		"error":          cfg["error"],
		"create_date":    cfg["create_date"],
		"creator":        "DBA",
		"updater":        "DBA",
	}
	data, _ := json.Marshal(detail)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	var res result
	_ = json.Unmarshal(body, &res)
	//fmt.Println("Msg=", res)
	return res
}

func doBackup(cfg map[string]interface{}) {
	//create backup directory
	err := os.MkdirAll(cfg["bk_path"].(string), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf(`Directory '%s' created!`, cfg["bk_path"].(string)))
	}

	// init sum var
	DbSizeHj := 0
	ElapsedBackupHj := 0
	ElapsedGzipHj := 0

	//loop backup db
	dbs := getDatabases("coupons")
	BackupBeginTime := getTime()
	GStatus := "0"
	fmt.Println("Starting backup database please wait...")
	for _, v := range dbs {
		error := ""
		status := "0"
		const ShellToUse = "bash"
		db := v["schema_name"].(string)
		fmt.Println(fmt.Sprintf(`Performing backup database: %s...`, db))
		StartTime := getTime()
		FileName := fmt.Sprintf(`%s/%s_%s.sql`, cfg["bk_path"], db, getDate())
		ErrName := fmt.Sprintf(`/tmp/%s_%s.err`, db, getDate())
		FullGzipName := fmt.Sprintf(`%s.gz`, FileName)
		GzipName := fmt.Sprintf(`%s_%s.sql.gz`, db, getDate())
		Command := fmt.Sprintf(`%s -u%s -p%s -h%s --port %s --single-transaction -routines --force --master-data=2 --databases %s -r %s &>%s`,
			cfg["bk_cmd"],
			cfg["db_user"],
			cfg["new_pass"],
			cfg["db_ip"],
			cfg["db_port"],
			db,
			FileName,
			ErrName)
		//fmt.Println(Command)
		cmd := exec.Command(ShellToUse, "-c", Command)
		_, err := cmd.Output()
		if err != nil {
			error = GetFIleContents(ErrName)
			status = "1"
			GStatus = "1"
		}
		EndTime := getTime()

		// gzip file
		Command = fmt.Sprintf(`cd %s && gzip -f %s`, cfg["bk_path"], FileName)
		cmd = exec.Command(ShellToUse, "-c", Command)
		_, err = cmd.Output()
		if err != nil {
			error = err.Error()
			status = "1"
		}
		ZipTime := getTime()

		// prepare detail interface parameter
		detail := map[string]string{
			"db_tag":         cfg["db_tag"].(string),
			"db_name":        db,
			"create_date":    getDate(),
			"file_name":      GzipName,
			"bk_path":        cfg["bk_path"].(string),
			"db_size":        strconv.Itoa(GetFIleSize(FullGzipName)),
			"start_time":     getTimeString(StartTime),
			"end_time":       getTimeString(ZipTime),
			"elapsed_backup": strconv.Itoa(getSeconds(StartTime, EndTime)),
			"elapsed_gzip":   strconv.Itoa(getSeconds(EndTime, ZipTime)),
			"status":         status,
			"error":          error,
		}
		// write backup detail log
		WriteLogDetail(detail)

		//sum static
		DbSizeHj = DbSizeHj + GetFIleSize(FullGzipName)
		ElapsedBackupHj = ElapsedBackupHj + getSeconds(EndTime, StartTime)
		ElapsedGzipHj = ElapsedGzipHj + getSeconds(ZipTime, EndTime)
	}

	// prepare total interface parameter
	total := map[string]string{
		"db_tag":         cfg["db_tag"].(string),
		"bk_base":        cfg["bk_base"].(string),
		"create_date":    getDate(),
		"start_time":     getTimeString(BackupBeginTime),
		"end_time":       getTimeString(getTime()),
		"total_size":     strconv.Itoa(DbSizeHj),
		"elapsed_backup": strconv.Itoa(ElapsedBackupHj),
		"elapsed_gzip":   strconv.Itoa(ElapsedGzipHj),
		"status":         GStatus,
	}

	// write backup total log
	WriteLogTotal(total)

}

func main() {
	//全局配置
	var cfg map[string]interface{}

	//命令行获取tag参数
	tag := GetTag()

	//调用接口读取配置
	res := GetConfig(tag)
	if res.Code == 200 {
		cfg = res.Data[0]
		//for k, v := range cfg  {
		//	fmt.Println(k,"=", v)
		//}
	} else {
		panic(res.Msg)
	}

	//预处理参数
	cfg = preProcess(cfg)

	//执行备份
	doBackup(cfg)

}
