package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "xhylblog/routers"
	"xhylblog/utils"

	"github.com/astaxie/beego"
)

func main() {
	//初始化日志
	utils.InitLog()

	//swagger生成
	utils.InitSwagger()

	//初始化mysql数据库
	utils.InitMysql()

	//跨域
	utils.CrossRequestFilter()

	//接口请求Token过滤器
	//访问接口前验证token
	utils.TokenFilter()

	beego.Run()
}
