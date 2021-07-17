一、概述  

   平台功能： EaseBase2.0 数据库自动化运维平台,后端采用go语言web框架beego,前端：vue前端框架 element.js

二、功能  
   
     1. 用户管理、角色管理、功能管理、菜单管理、系统设置
       
     2. 数据源管理、服务器管理、数据库库管理、慢日志管理、数据库监控
       
     3. 数据库备份、数据库同步、数据库传输、数据库归档、大数据同步
       
     4. 工单管理、图片管理、端口管理  
       
三、编译

     bee pack -be GOOS=windows
     bee pack -be GOOS=linux
     
四、数据库初始化

    ./sql/puppet2.0.sql

五、启动

    bee run
    Windows:  easebase2.0.exe
    Linux:   ./easebase2.0
        
六、配置 

   conf/app.conf
    
七、访问

    http://ip:9000/
    
八、其它

   	1、Windows 下go编译成linux可执行文件
   	
        SET CGO_ENABLED=0
        SET GOOS=linux
        SET GOARCH=amd64
        go build dbbackup.go
        chmod +x dbbackup
        ./dbbackup -api_server=10.16.47.114:9000 -db_tag=mysql_10_2_39_80_3306
   
   	2、Windows 下go编译成可执行文件
   	
        SET CGO_ENABLED=0
        SET GOOS=windows
        SET GOARCH=amd64
        go build dbbackup.go
        chmod +x dbbackup
        ./dbbackup mysql_10_2_39_40_3306
   
    3、开发环境初始化
    
        go env -w GOPROXY=https://goproxy.cn
        go list -m -u all