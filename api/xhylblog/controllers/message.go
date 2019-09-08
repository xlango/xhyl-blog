package controllers

import "xhylblog/service"

//  邮件短信提醒
type MessageController struct {
	BaseController
}

// URLMapping ...
func (c *MessageController) URLMapping() {
	c.Mapping("Email", c.Email)
}

// GetOne ...
// @Title Get One
// @Description get Article by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :id is empty
// @router /email [get]
func (c *MessageController) Email() {
	service:=&service.MessageServcie{}
	service.WsHandler(c.Ctx.ResponseWriter, c.Ctx.Request)
}