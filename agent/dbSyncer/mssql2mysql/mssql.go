package mssql2mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"strconv"
	"time"
)

type Mssql struct {
	ip       string
	port     string
	user     string
	password string
	database string
	DB       *sql.DB
}

//建立连接
func (m *Mssql) Connect() {
	port, _ := strconv.Atoi(m.port)
	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s;encrypt=disable", m.ip, port, m.database, m.user, m.password)
	fmt.Println(connString)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	fmt.Println("db connect success!")
	m.DB = conn
}

//关闭连接
func (m *Mssql) Close() {
	m.DB.Close()
}

//返回所有列名信息
func (m *Mssql) GetColumnNames(rows *sql.Rows) {
	cols, err := rows.Columns()
	fmt.Println("cols=", cols)
	if err != nil {
		log.Fatal("GetColumnNames failed:", err.Error())
	}
	var res = make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		res[i] = new(interface{})
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
	fmt.Println()
}

//打印一行记录
func (m *Mssql) PrintRow(res []interface{}) {
	for _, val := range res {
		switch v := (*(val.(*interface{}))).(type) {
		case nil:
			fmt.Print("NULL")
		case bool:
			if v {
				fmt.Print("True")
			} else {
				fmt.Print("False")
			}
		case []byte:
			fmt.Print(string(v))
		case time.Time:
			fmt.Print(v.Format("2016-01-02 15:05:05.999"))
		default:
			fmt.Print(v)
		}
		fmt.Print("\t")
	}
	fmt.Println()
}

//输出结果集
func (m *Mssql) GetRowContents(rows *sql.Rows) {
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal("GetColumnNames failed:", err.Error())
	}
	var res = make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		res[i] = new(interface{})
	}

	for rows.Next() {
		rows.Scan(res...) //将查到的数据写入到这行中
		m.PrintRow(res)   //打印此行
	}
}

//插入数据
func (m *Mssql) Insert(xh int, xm string, age int) {
	stmt, err := m.DB.Prepare(`insert into xs(xh,xm,age) values(?,?,?)`)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()
	stmt.Exec(xh, xm, age)
	fmt.Println("insert success!")
}

//更新数据
func (m *Mssql) Update(xh int, xm string, age int) {
	stmt, err := m.DB.Prepare("update xs set xm=? , age=? where xh=?")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(xm, age, xh)
	if err != nil {
		log.Fatal("Exec failed:", err.Error())
	}
	fmt.Println("update success!")
}

//删除数据
func (m *Mssql) Delete(xh int) {
	stmt, err := m.DB.Prepare("delete from xs  where xh=?")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(xh)
	if err != nil {
		log.Fatal("Delete failed:", err.Error())
	}
	fmt.Println("delete success!")
}

//执行DML语句
func (m *Mssql) Exec(sql string) {
	_, err := m.DB.Exec(sql)
	if err != nil {
		log.Fatal("Execute failed:", err.Error())
	}
	fmt.Println("Execute success!")
}

//手动管理事务
func (m *Mssql) ExecTrans() {
	tx, err := m.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO xs(xh,xm,age) VALUES (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(i, "dog"+strconv.Itoa(i), 20+i)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}

//查询表数据
func (m *Mssql) Query() {
	stmt, err := m.DB.Prepare(`select * from [xs]`)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}

	//输出列名
	m.GetColumnNames(rows)

	//输出列值
	m.GetRowContents(rows)

}

//查询表数据参数
func (m *Mssql) QueryParam(statement string) *sql.Rows {
	stmt, err := m.DB.Prepare(statement)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	return rows
}

func main() {
	var m = Mssql{"10.2.39.9", "1433", "hopsonone", "hopsonone123", "test", nil}
	m.Connect()
	m.Query()
	m.Insert(3, "wang.wu", 32)
	m.Update(3, "wang.qi", 52)
	m.Delete(3)
	m.ExecTrans()
	m.Close()
}
