package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	_ "strings"
)

type result struct {
	Name string                   `json:"Name"`
	Code int                      `json:"Code"`
	Msg  string                   `json:"Msg"`
	Data []map[string]interface{} `json:"Data"`
}

type Backup struct {
	DbTag      string
	CreateDate string
}

func Get(Url string) result {
	resp, err := http.Get(Url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var res result
	fmt.Println("body:", body)
	_ = json.Unmarshal(body, &res)
	return res
}

func PostFormMultiParam(Url string) {
	var r http.Request
	r.ParseForm()
	r.Form.Add("db_tag", "mysql_172.17.194.53_3306")
	r.Form.Add("create_date", "2021-05-25")
	body := strings.NewReader(r.Form.Encode())
	clt := http.Client{}
	resp, err := clt.Post(Url, "application/x-www-form-urlencoded", body)
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	respBody := string(content)
	fmt.Println("respBody=", respBody)

}

func PostMultiParam(Url string) {
	param := url.Values{}
	param.Add("db_tag", "test")
	param.Add("create_date", "22")
	resp, _ := http.PostForm(Url, param)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func PostStructByJson(Url string) {
	param := Backup{
		DbTag:      "mysql_172.17.194.53_3306",
		CreateDate: "2021-05-26",
	}
	data, _ := json.Marshal(param)
	req, _ := http.NewRequest("POST", Url, bytes.NewReader(data))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func PostMapByJson(Url string) {
	//备份汇总日志
	total := map[string]string{
		"db_tag":         "mysql_172.17.194.53_3306",
		"db_type":        "1",
		"bk_base":        "/home/hopson/apps/usr/webserver/dba/db_backup/mysqldump_easylife_uat",
		"total_size":     "132.4m",
		"start_time":     "2021-05-19 00:30:05",
		"end_time":       "2021-05-19 00:31:20",
		"create_date":    "2021-05-26",
		"elapsed_backup": "793",
		"elapsed_gzip":   "234",
		"status":         "1",
		"creator":        "DBA",
		"updater":        "DBA",
	}

	data, _ := json.Marshal(total)
	req, _ := http.NewRequest("POST", Url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("body:",body)
	fmt.Println(string(body))
	var res result
	_ = json.Unmarshal(body, &res)
	fmt.Println("Msg=", res)

	//备份明细日志
	detail := map[string]string{
		"db_tag":         "mysql_172.17.194.53_3306",
		"db_name":        "hopsonone_bill",
		"file_name":      "hopsonone_bill_20210519.sql.gz",
		"bk_path":        "/home/hopson/apps/usr/webserver/dba/db_backup/mysqldump_bi/20210519",
		"db_size":        "356M",
		"start_time":     "2021-05-19 00:30:05",
		"end_time":       "2021-05-19 00:31:20",
		"elapsed_backup": "793",
		"elapsed_gzip":   "234",
		"status":         "1",
		"error":          "error",
		"create_date":    "2021-05-26",
		"creator":        "DBA",
		"updater":        "DBA",
	}

	url := "http://10.16.47.114:9000/api/backup/detail/mysql_172.17.194.53_3306"
	data, _ = json.Marshal(detail)
	req, _ = http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	resp, _ = client.Do(req)
	body, _ = ioutil.ReadAll(resp.Body)
	//fmt.Println("body:",body)
	fmt.Println(string(body))
	var res2 result
	_ = json.Unmarshal(body, &res2)
	fmt.Println("Msg=", res2)

}

func main() {
	//url := "http://10.16.47.114:9000/api/backup/config/mysql_172.17.194.53_3306"
	//res := Get(url)
	//for k, v := range res.Data[0]  {
	//	fmt.Println(k,"=", v)
	//}

	url := "http://10.16.47.114:9000/api/backup/log/mysql_172.17.194.53_3306"
	PostMapByJson(url)
	//PostFormMultiParam(url)
	//PostMultiParam(url)
	//PostStructByJson(url)

}
