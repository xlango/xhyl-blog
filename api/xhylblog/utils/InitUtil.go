package utils

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

)

//初始化日志输出文件
func InitLog()  {
	config := make(map[string]interface{})
	config["filename"] = "main.log"
	config["level"] = logs.LevelError
	configStr, err := json.Marshal(config)
	if err != nil {
		logs.Error("logs init maral failed, err:", err.Error())
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
}

//初始化mysql数据库
func InitMysql()  {
	//mysql注册
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"),30,30)
}

//初始化postgresql数据库
func InitPostgreSql()  {
	//postgresql注册
	orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
	orm.RegisterDataBase("default", "postgres", "postgres://fusionpbx:mypassword@192.168.10.250:5432/fusionpbx?sslmode=disable")
}

//初始化swagger
func InitSwagger()  {
	//swagger生成
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}