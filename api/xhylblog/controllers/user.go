package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"xhylblog/models"
)

//  UserController operations for User
type UserController struct {
	BaseController
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Logout", c.Logout)
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}


// Login ...
// @Title 用户登录获取token
// @Description 用户登录获取token
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router /login [post]
func (c *UserController) Login() {
	var v models.User
	c.JsonParam(&v)
	if token, err := models.Login(v.Username,v.Password); err == nil {
		l, _ := models.GetAllUser(map[string]string{"Username":v.Username}, nil, nil, nil, 0, 0)
		rs:=models.LoginResult{
			User:l[0].(models.User),
			Token:token,
		}
		c.Ok(rs)
	} else {
		c.RequestError("账号或密码错误！")
	}
}

// Logout ...
// @Title 用户注销
// @Description 用户注销
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router /logout [post]
func (c *UserController) Logout() {
	var v models.User
	c.JsonParam(&v)
	if token, err := models.Login(v.Username,v.Password); err == nil {
		c.Ok(token)
	} else {
		c.RequestError("账号或密码错误！")
	}
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router /register [post]
func (c *UserController) Post() {
	var v models.User
	c.JsonParam(&v)
	if id, _ := models.AddUser(&v); id != 0 {
		c.Ok(v)
	} else {
		c.RequestError("该用户名已经注册！")
	}
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetUserById(id)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Ok(v)
	}
}

// GetAll ...
// @Title Get All
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Ok(l)
	}
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.User{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateUserById(&v); err == nil {
		c.Ok("OK")
	} else {
		c.Error(err.Error())
}
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteUser(id); err == nil {
		c.Ok("OK")
	} else {
		c.Error(err.Error())
	}
}
