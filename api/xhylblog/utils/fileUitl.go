package utils

import "os"

var CurrentDir string

func init() {
	//解决windows无法debug启动的问题
	CurrentDir = GetCurrentExeDir()
	if len(CurrentDir) > 0 {
		CurrentDir = CurrentDir + string(os.PathSeparator)
	}
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(CurrentDir+path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
