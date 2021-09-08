package mssql2mysql

func preProcess(cfg map[string]interface{}) map[string]interface{} {
	InitSourDB(cfg)
	InitTargetDB(cfg)
	return cfg
}

func Sync(api *string, tag *string, show *bool) {

	//配置参数
	var cfg map[string]interface{}

	res := GetConfig(*api, *tag)
	if res.Code == 200 {
		cfg = res.Data[0]
	} else {
		panic(res.Msg)
	}

	//预处理参数
	cfg = preProcess(cfg)

	//输出配置
	if *show {
		ShowConfig(cfg)
	}

	//从配置中获取数据库连接
	mssql := cfg["sourceDB"].(Mssql)

	//获取结果集
	rows := mssql.QueryParam("select top 10 CountDate,SiteKey,DateKey,SiteType,Area,NetArea from dbo.summary_thirty")
	//输出列名
	mssql.GetColumnNames(rows)

	//输出列值
	mssql.GetRowContents(rows)

}
