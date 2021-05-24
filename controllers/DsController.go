package controllers

import (
	"easebase_web/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

type DsController struct {
	BaseController
}

type DsControllerByParId struct {
	BaseController
}

type DsServerController struct {
	BaseController
}

type DsTestController struct {
	BaseController
}

//query ds
func (c *DsController) Get() {
	dsName := c.GetString("ds_name")
	marketId := c.GetString("market_id")
	dbType := c.GetString("db_type")
	dbEnv := c.GetString("db_env")
	fmt.Println(dsName, marketId, dbType, dbEnv)
	vWhere := ""
	if dsName != "" {
		st := fmt.Sprintf(` and instr(BINARY CONCAT(a.db_desc,b.dmmc,':/',ip,':',PORT,'/',service),'%s')>0 `, dsName)
		vWhere = vWhere + st
	}

	if marketId != "" {
		st := fmt.Sprintf(` and a.market_id='%s'`, marketId)
		vWhere = vWhere + st
	}

	if dbType != "" {
		st := fmt.Sprintf(` and a.db_type='%s'`, dbType)
		vWhere = vWhere + st
	}

	if dbEnv != "" {
		st := fmt.Sprintf(` and a.db_env='%s'`, dbEnv)
		vWhere = vWhere + st
	}

	o := orm.NewOrm()
	var users []orm.Params
	st := fmt.Sprintf(
		`select  a.id,
				d.dmmc AS market_name,
				a.db_desc,
				c.dmmc AS db_env,
				CONCAT(SUBSTR(CONCAT(b.dmmc,':/',ip,':',PORT,'/',service),1,40),'...')  AS ds_name,              
				user,
				CASE STATUS WHEN '1' THEN '是' WHEN '0' THEN '否' END as  status,                    
				updater,
				a.last_update_date AS last_update_date         
		from t_db_source a,t_dmmx b,t_dmmx c,t_dmmx d
		where a.db_type=b.dmm AND b.dm='02'
		   and a.db_env=c.dmm  AND c.dm='03' 
		   and a.market_id=d.dmm  AND d.dm='05' 
		   %s  ORDER BY a.db_desc`, vWhere)

	_, err := o.Raw(st).Values(&users)
	if err == nil {
		c.SuccessJson("DsController->Get", &users)
	} else {
		c.ErrorJson("DsController->Get", 500, err.Error(), nil)
	}

}

//insert ds
func (c *DsController) Put() {
	MarketId := c.GetString("market_id")
	InstType := c.GetString("inst_type")
	DbType := c.GetString("db_type")
	DbEnv := c.GetString("db_env")
	DbDesc := c.GetString("db_desc")
	Ip := c.GetString("ip")
	Port := c.GetString("port")
	Service := c.GetString("service")
	User := c.GetString("user")
	Password := c.GetString("password")
	Status := c.GetString("status")
	ProxyStatus := c.GetString("proxy_status")
	ProxyServer := c.GetString("proxy_server")

	o := orm.NewOrm()
	ds := models.TDbSource{
		MarketId:    MarketId,
		InstType:    InstType,
		DbType:      DbType,
		DbEnv:       DbEnv,
		DbDesc:      DbDesc,
		Ip:          Ip,
		Port:        Port,
		Service:     Service,
		User:        User,
		Password:    Password,
		Status:      Status,
		ProxyStatus: ProxyStatus,
		ProxyServer: ProxyServer,
		Creator:     "DBA",
		Updater:     "DBA",
	}
	id, err := o.Insert(&ds)
	fmt.Println(id, err)
	if err != nil {
		c.ErrorJson("DsController->Put", 500, err.Error(), nil)
	} else {
		c.SuccessJson("DsController->Put", "id="+strconv.FormatInt(id, 10))
	}
}

//update user
func (c *DsController) Post() {
	dsId, _ := c.GetInt("id")
	MarketId := c.GetString("market_id")
	InstType := c.GetString("inst_type")
	DbType := c.GetString("db_type")
	DbEnv := c.GetString("db_env")
	DbDesc := c.GetString("db_desc")
	Ip := c.GetString("ip")
	Port := c.GetString("port")
	Service := c.GetString("service")
	User := c.GetString("user")
	Password := c.GetString("password")
	Status := c.GetString("status")
	ProxyStatus := c.GetString("proxy_status")
	ProxyServer := c.GetString("proxy_server")

	o := orm.NewOrm()
	ds := models.TDbSource{
		Id:          dsId,
		MarketId:    MarketId,
		InstType:    InstType,
		DbType:      DbType,
		DbEnv:       DbEnv,
		DbDesc:      DbDesc,
		Ip:          Ip,
		Port:        Port,
		Service:     Service,
		User:        User,
		Password:    Password,
		Status:      Status,
		ProxyStatus: ProxyStatus,
		ProxyServer: ProxyServer,
		Creator:     "DBA",
		Updater:     "DBA",
	}
	_, err := o.Update(&ds)
	if err != nil {
		c.ErrorJson("更新数据源出错!", 500, err.Error(), nil)
	} else {
		c.SuccessJson("数据源变更成功!", "")
	}
}

//delete user
func (c *DsController) Delete() {
	dsId, _ := c.GetInt("id")
	o := orm.NewOrm()
	ds := models.TDbSource{Id: dsId}
	_, err := o.Delete(&ds)
	if err != nil {
		c.ErrorJson("DsController->Delete", 500, err.Error(), nil)
	} else {
		c.SuccessJson("DsController->Delete", nil)
	}
}

//query ds by id
func (c *DsControllerByParId) Get() {
	id := c.Ctx.Input.Param(":id")
	int32, err := strconv.Atoi(id)
	if err != nil {
		c.ErrorJson("UserControllerByParId->Get", 500, err.Error(), nil)
	}
	o := orm.NewOrm()
	u := models.TDbSource{Id: int32}
	err = o.Read(&u)
	if err != nil {
		c.ErrorJson("DsControllerByParId->Get", 500, err.Error(), nil)

	} else {
		c.SuccessJson("DsControllerByParId->Get", &u)
	}
}

//get ds server
func (c *DsServerController) Get() {
	o := orm.NewOrm()
	var dbServers []orm.Params
	st := fmt.Sprintf(`SELECT '' AS dmm,'请选择...'  AS dmmc 
						  	  union all
						   	  SELECT id,db_desc FROM t_db_source 
							   WHERE  (db_type IN(0)  AND USER IN('puppet','easylife','apptong') OR db_type NOT IN (0))
							    AND STATUS=1 ORDER BY dmm+0`)
	fmt.Println(st)
	_, err := o.Raw(st).Values(&dbServers)
	if err == nil {
		c.SuccessJson("BackupServerController->Get", &dbServers)
	} else {
		fmt.Println(err.Error())
		c.ErrorJson("BackupServerController->Get", 500, err.Error(), nil)
	}

}
