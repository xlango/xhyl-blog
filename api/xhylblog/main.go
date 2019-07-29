package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "xhylblog/routers"
	"xhylblog/utils"

	"github.com/astaxie/beego"
)

func main() {
	//初始化日志
	utils.InitLog()

	//swagger生成
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//mysql注册
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"),30,30)

	beego.Run()
}
