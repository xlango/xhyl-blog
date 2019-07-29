package controllers

import (
	"xhylblog/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type BaseController struct {
	beego.Controller
}

//返回成功
func (c *BaseController) Ok(rs interface{}) {
	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = models.Ok(rs)
	c.ServeJSON()
}

//认证失败
func (c *BaseController) AuthErr() {
	c.Ctx.Output.SetStatus(401)
	c.Data["json"] = models.AuthErr()
	c.ServeJSON()
}

//参数错误
func (c *BaseController) ParamErr() {
	c.Ctx.Output.SetStatus(403)
	c.Data["json"] = models.ParamErr()
	c.ServeJSON()
}

//404
func (c *BaseController) Nofound() {
	c.Ctx.Output.SetStatus(404)
	c.Data["json"] = models.Nofound()
	c.ServeJSON()
}

//内部错误
func (c *BaseController) Error(rs interface{}) {
	c.Ctx.Output.SetStatus(500)
	c.Data["json"] = models.Error(rs)
	c.ServeJSON()
}

//json参数解析为结构体
func (c *BaseController) JsonParam(param interface{}) interface{}{
	data := c.Ctx.Input.RequestBody
	//json数据封装到api对象中
	err := json.Unmarshal(data, &param)
	if err != nil {
		logs.Error("json.Unmarshal is err ",err.Error())
	}

	return param
}
