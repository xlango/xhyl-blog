package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
)

//跨域
func CrossRequestFilter()  {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type","token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
}

//接口请求Token过滤器
func TokenFilter()  {
	var FilterToken = func(ctx *context.Context) {
		logs.Info("current router path is ", ctx.Request.RequestURI)
		if ctx.Request.RequestURI != "/xhyl/user/login" && ctx.Request.RequestURI != "/xhyl/user/register" && ctx.Input.Header("token") == ""&& !ctx.Input.IsGet() {
			logs.Error("without token, unauthorized !!")
			ctx.ResponseWriter.WriteHeader(401)
			ctx.ResponseWriter.Write([]byte("no permission"))
		}
		if ctx.Request.RequestURI != "/xhyl/user/login" && ctx.Request.RequestURI != "/xhyl/user/register" && ctx.Request.Header.Get("token") != ""&& !ctx.Input.IsGet() {
			token := ctx.Request.Header.Get("token")

			if err := ValidateToken(token);err!=nil {
				ctx.ResponseWriter.WriteHeader(403)
				ctx.ResponseWriter.Write([]byte("Token authentication failed!"))
			}
		}
	}
	beego.InsertFilter("*", beego.BeforeRouter, FilterToken)
}