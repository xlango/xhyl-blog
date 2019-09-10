package utils

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

/**
创建定时任务
 */
func CreateCronTask(taskName string,time string,task  func() error) (tk *toolbox.Task,err error){
	tk = toolbox.NewTask(taskName, time, task)
	err = tk.Run()
	if err != nil {
		logs.Error(err.Error())
		return
	}
	logs.Info("create cron task ",taskName)
	return tk,nil
}

/**
开始定时任务
 */
func StartCronTask(taskName string,tk *toolbox.Task) (err error) {
	toolbox.AddTask(taskName, tk)
	toolbox.StartTask()
	logs.Info("start cron task ",taskName)
	return
}

/**
关闭定时任务
 */
func StopCronTask(taskName string,tk *toolbox.Task) (err error) {
	toolbox.AddTask(taskName, tk)
	toolbox.StopTask()
	logs.Info("stop cron task ",taskName)
	return
}