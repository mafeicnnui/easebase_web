package controllers

import (
	"easebase_web/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

type ServerController struct {
	BaseController
}

type ServerControllerByParId struct {
	BaseController
}

//query server
func (c *ServerController) Get() {
	serverName := c.GetString("server_name")
	vWhere := ""
	if serverName != "" {
		st := fmt.Sprintf(` and instr(binary concat(a.market_id,'|',a.server_ip,'|',a.server_port,'|',a.server_desc),'%s')>0`, serverName)
		vWhere = vWhere + st
	}
	o := orm.NewOrm()
	var users []orm.Params
	st := fmt.Sprintf(
		`SELECT  a.id,
				c.dmmc as server_type,
				a.server_desc,
				a.market_id,
				b.dmmc as market_name,
				a.server_ip,
				a.server_port,
				a.server_user,
				a.server_os,
				a.server_cpu,
				a.server_mem,
				CASE a.STATUS WHEN '1' THEN '启用' WHEN '0' THEN '禁用' END  status,
                a.last_update_date
			FROM t_server a,t_dmmx b,t_dmmx c
			WHERE a.market_id=b.dmm AND b.dm='05'
			  and a.server_type=c.dmm and c.dm='06' %s  order by a.market_id,a.server_port`, vWhere)
	fmt.Println(st)
	_, err := o.Raw(st).Values(&users)
	if err == nil {
		c.SuccessJson("DsController->Get", &users)
	} else {
		c.ErrorJson("DsController->Get", 500, err.Error(), nil)
	}

}

//insert server
func (c *ServerController) Put() {
	MarketId := c.GetString("market_id")
	ServerType := c.GetString("server_type")
	ServerDesc := c.GetString("server_desc")
	ServerAuth := c.GetString("server_auth")
	ServerIp := c.GetString("server_ip")
	ServerPort := c.GetString("server_port")
	ServerUser := c.GetString("server_user")
	ServerPass := c.GetString("server_pass")
	ServerKeys := c.GetString("server_keys")
	ServerOs := c.GetString("server_os")
	ServerCpu := c.GetString("server_cpu")
	ServerMem := c.GetString("server_mem")
	Status := c.GetString("status")

	o := orm.NewOrm()
	server := models.TServer{
		MarketId:   MarketId,
		ServerType: ServerType,
		ServerDesc: ServerDesc,
		ServerAuth: ServerAuth,
		ServerIp:   ServerIp,
		ServerPort: ServerPort,
		ServerUser: ServerUser,
		ServerPass: ServerPass,
		ServerKeys: ServerKeys,
		ServerOs:   ServerOs,
		ServerCpu:  ServerCpu,
		ServerMem:  ServerMem,
		Status:     Status,
		Creator:    "DBA",
		Updater:    "DBA",
	}
	id, err := o.Insert(&server)
	fmt.Println(id, err)
	if err != nil {
		c.ErrorJson("ServerController->Put", 500, err.Error(), nil)
	} else {
		c.SuccessJson("ServerController->Put", "id="+strconv.FormatInt(id, 10))
	}
}

//update server
func (c *ServerController) Post() {
	serverId, _ := c.GetInt("id")
	MarketId := c.GetString("market_id")
	ServerType := c.GetString("server_type")
	ServerDesc := c.GetString("server_desc")
	ServerAuth := c.GetString("server_auth")
	ServerIp := c.GetString("server_ip")
	ServerPort := c.GetString("server_port")
	ServerUser := c.GetString("server_user")
	ServerPass := c.GetString("server_pass")
	ServerKeys := c.GetString("server_keys")
	ServerOs := c.GetString("server_os")
	ServerCpu := c.GetString("server_cpu")
	ServerMem := c.GetString("server_mem")
	Status := c.GetString("status")

	o := orm.NewOrm()
	server := models.TServer{
		Id:         serverId,
		MarketId:   MarketId,
		ServerType: ServerType,
		ServerDesc: ServerDesc,
		ServerAuth: ServerAuth,
		ServerIp:   ServerIp,
		ServerPort: ServerPort,
		ServerUser: ServerUser,
		ServerPass: ServerPass,
		ServerKeys: ServerKeys,
		ServerOs:   ServerOs,
		ServerCpu:  ServerCpu,
		ServerMem:  ServerMem,
		Status:     Status,
		Creator:    "DBA",
		Updater:    "DBA",
	}
	_, err := o.Update(&server)
	if err != nil {
		c.ErrorJson("更新服务器出错!", 500, err.Error(), nil)
	} else {
		c.SuccessJson("服务器变更成功!", "")
	}
}

//delete server
func (c *ServerController) Delete() {
	dsId, _ := c.GetInt("id")
	o := orm.NewOrm()
	server := models.TServer{Id: dsId}
	_, err := o.Delete(&server)
	if err != nil {
		c.ErrorJson("ServerController->Delete", 500, err.Error(), nil)
	} else {
		c.SuccessJson("ServerController->Delete", nil)
	}
}

//query server by id
func (c *ServerControllerByParId) Get() {
	id := c.Ctx.Input.Param(":id")
	int32, err := strconv.Atoi(id)
	if err != nil {
		c.ErrorJson("ServerControllerByParId->Get", 500, err.Error(), nil)
	}
	o := orm.NewOrm()
	server := models.TServer{Id: int32}
	err = o.Read(&server)
	if err != nil {
		c.ErrorJson("ServerControllerByParId->Get", 500, err.Error(), nil)

	} else {
		c.SuccessJson("ServerControllerByParId->Get", &server)
	}
}
