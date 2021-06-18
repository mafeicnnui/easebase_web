package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	_ "os"
	"os/exec"
	"strconv"
	"strings"
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

func getDate2() string {
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	//hour   := time.Now().Format("15")
	//min    := time.Now().Format("04")
	//second := time.Now().Format("05")
	return year + month + day
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

func GetConfig(ApiServer string, DbTag string) result {
	var res result
	url := fmt.Sprintf(`http://%s/api/backup/config/%s`, ApiServer, DbTag)
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

func GetPassword(cfg map[string]interface{}) string {
	var res result
	url := fmt.Sprintf(`http://%s/api/public/decode`, cfg["api_server"])
	tag := map[string]interface{}{
		"key":      cfg["db_user"],
		"password": cfg["db_pass"],
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

func IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

func preProcess(cfg map[string]interface{}) map[string]interface{} {
	cfg["bk_path"] = fmt.Sprintf(`%s/%s`, cfg["bk_base"], getDate2())
	cfg["new_pass"] = GetPassword(cfg)
	InitDB(cfg)
	return cfg
}

func WriteLogTotal(cfg map[string]string) result {
	// 请求URL
	url := fmt.Sprintf(`http://%s/api/backup/log/%s`, cfg["api_server"], cfg["db_tag"])

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
	var res result
	_ = json.Unmarshal(body, &res)
	return res
}

func WriteLogDetail(cfg map[string]string) result {
	// 请求URL
	url := fmt.Sprintf(`http://%s/api/backup/detail/%s`, cfg["api_server"], cfg["db_tag"])

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
	var res result
	_ = json.Unmarshal(body, &res)
	return res
}

func doBackup(cfg map[string]interface{}) {
	//create backup directory
	if IsDir(cfg["bk_path"].(string)) {
		fmt.Println(fmt.Sprintf(`Directory '%s' already exists!`, cfg["bk_path"].(string)))
	} else {
		err := os.MkdirAll(cfg["bk_path"].(string), os.ModePerm)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf(`Directory '%s' created!`, cfg["bk_path"].(string)))
		}
	}

	// init sum var
	DbSizeHj := 0
	ElapsedBackupHj := 0
	ElapsedGzipHj := 0

	//loop backup db
	dbs := getDatabases(cfg["backup_databases"].(string))
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
			"api_server":     cfg["api_server"].(string),
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
		ElapsedBackupHj = ElapsedBackupHj + getSeconds(StartTime, EndTime)
		ElapsedGzipHj = ElapsedGzipHj + getSeconds(EndTime, ZipTime)
	}

	// prepare total interface parameter
	total := map[string]string{
		"db_tag":         cfg["db_tag"].(string),
		"api_server":     cfg["api_server"].(string),
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
	fmt.Println("Backup database complete!")

}

func showConfig(cfg map[string]interface{}) {
	fmt.Println(strings.Repeat("-", 140))
	fmt.Println(fmt.Sprintf(`%-20s%-120s`, "配置项", "配置值"))
	fmt.Println(strings.Repeat("-", 140))
	for k, v := range cfg {
		fmt.Println(fmt.Sprintf(`%-20s=%-120s`, k, v))
	}
	fmt.Println(strings.Repeat("-", 140))
}

func main() {

	//全局配置
	var cfg map[string]interface{}

	//命令行获取tag参数
	//-db_tag=mysql_10_2_39_40_3306 -api_server=10.16.47.114:9000
	api := flag.String("api_server", "", "接口服务地址,格式：IP:PORT")
	tag := flag.String("db_tag", "", "数据备份标识")
	show := flag.Bool("config", false, "是否显示配置信息")
	flag.Parse()

	//调用接口读取配置
	res := GetConfig(*api, *tag)
	if res.Code == 200 {
		cfg = res.Data[0]
		if *show {
			showConfig(cfg)
		}
	} else {
		panic(res.Msg)
	}

	//预处理参数
	cfg = preProcess(cfg)

	//执行备份
	doBackup(cfg)

}
