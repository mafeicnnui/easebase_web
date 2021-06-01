package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Backup struct {
	DbTag      string
	CreateDate string
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func CheckTabExists(pTab string, pWhere string) int {
	o := orm.NewOrm()
	var res []orm.Params
	st := fmt.Sprintf(`select count(0) as rec from %s  %s`, pTab, pWhere)
	_, err1 := o.Raw(st).Values(&res)
	if err1 != nil {
		panic("Func CheckTabExists Error:" + err1.Error())
	}
	val, err2 := strconv.Atoi(res[0]["rec"].(string))
	if err2 != nil {
		panic("Func CheckTabExists Error:" + err2.Error())
	}
	fmt.Println("CheckTabExists=", st, val)
	return val
}

func Decrypt(pPassword string, pKey string) ([]orm.Params, error) {
	o := orm.NewOrm()
	var res []orm.Params
	st := fmt.Sprintf(`select aes_decrypt(unhex('%s'),'%s') as password`, pPassword, reverseString(pKey))
	_, err := o.Raw(st).Values(&res)
	return res, err
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
