package utils

import (
	"encoding/json"
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
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/idata?charset=utf8")
}
