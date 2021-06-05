package routers

import (
	"easebase_web/controllers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins: true,
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))

	beego.Router("/", &controllers.IndexController{})
	beego.Router("/tree", &controllers.TreeController{})
	beego.Router("/tree2", &controllers.TreeController2{})
	beego.Router("/dm", &controllers.DmController{})
	//user
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/:id:int", &controllers.UserControllerByParId{})
	beego.Router("/user/role/:userId:int", &controllers.UserRoleControllerByParId{})
	//menu
	beego.Router("/menu", &controllers.MenuController{})
	beego.Router("/menu/:id:int", &controllers.MenuControllerByParId{})
	beego.Router("/menu/parent", &controllers.MenuParentController{})
	beego.Router("/menu/qx", &controllers.MenuQxController{})
	beego.Router("/menu/role/qx", &controllers.MenuRoleQxController{})
	//role
	beego.Router("/role", &controllers.RoleController{})
	beego.Router("/role/:id:int", &controllers.RoleControllerByParId{})
	beego.Router("/role/privileges/:roleId:int", &controllers.RolePrivilegesControllerByParId{})

	//db source
	beego.Router("/ds", &controllers.DsController{})
	beego.Router("/ds/:id:int", &controllers.DsControllerByParId{})
	beego.Router("/ds/server", &controllers.DsServerController{})

	//server
	beego.Router("/server", &controllers.ServerController{})
	beego.Router("/server/:id:int", &controllers.ServerControllerByParId{})

	//db backup
	beego.Router("/backup", &controllers.BackupController{})
	beego.Router("/backup/:id:int", &controllers.BackupControllerByParId{})
	beego.Router("/backup/server", &controllers.BackupServerController{})
	beego.Router("/backup/log", &controllers.BackupLogController{})
	beego.Router("/backup/task", &controllers.BackupTaskController{})

	//db sync
	beego.Router("/sync", &controllers.SyncController{})
	beego.Router("/sync/:id:int", &controllers.SyncControllerByParId{})
	beego.Router("/sync/server", &controllers.SyncServerController{})
	beego.Router("/sync/log", &controllers.SyncLogController{})
	beego.Router("/sync/task", &controllers.SyncTaskController{})

	// public api
	beego.Router("/api/public/:flag", &controllers.PublicApiController{})

	// agent api
	beego.Router("/api/backup/:flag/:tag", &controllers.BackupApiController{})

}
