package controllers

import (
	"github.com/astaxie/beego/logs"
	"os"
	"xhylblog/utils"
)

type FileController struct {
	BaseController
}

// URLMapping ...
func (c *FileController) URLMapping() {
	c.Mapping("Upload", c.Upload)
	c.Mapping("Download", c.Download)
}



// upload ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router /upload [post]
func (c *FileController) Upload() {
	// 读取文件信息
	f, h, err := c.GetFile("file")
	// 文件目录
	path := c.GetString("path")

	if path == "" {
		logs.Error("未填写文件路径", err)
		c.Error(err.Error())
	}

	if err != nil {
		logs.Error("读取文件错误", err)
		c.Error(err.Error())
	}

	// 延迟关闭文件
	defer f.Close()

	// 保存文件, 本地文件路径static/upload/上传文件名
	// 需要提前创建好static/upload目录
	uploadPath:="static/file/"+path+"/"
	exist, err :=utils.PathExists(uploadPath)
	if err != nil {
		logs.Error("get dir error![%v]\n", err)
		return
	}

	if exist {
		logs.Info("has dir![%v]\n", uploadPath)
	} else {
		logs.Info("no dir![%v]\n", uploadPath)
		// 创建文件夹
		err := os.Mkdir(utils.CurrentDir+uploadPath, os.ModePerm)
		if err != nil {
			logs.Error("mkdir failed![%v]\n", err)
		} else {
			logs.Info("mkdir success!\n")
		}
	}

	fileErr := c.SaveToFile("file", "static/file/"+path+"/"+h.Filename)
	if fileErr!=nil {
		logs.Error("保存文件错误", err)
		c.Error(fileErr.Error())
	}
	c.Ok("上传成功")
}


// donwload ...
// @Title Get One
// @Description get Article by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :id is empty
// @router /download/:filename [get]
func (c *FileController) Download() {
	// 业务逻辑处理，例如先检测用户权限
	filename := c.Ctx.Input.Param(":filename")
	// 下载服务器上当前目录/data/file.zip文件， 下载后的文件名为:压缩包1.zip
	c.Ctx.Output.Download("static/file/", filename)
}