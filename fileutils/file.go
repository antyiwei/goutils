package fileutils

import "os"

// 判断文件是否存在，不存在则创建并添加，否则只尾添加。
func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil { //文件或者目录存在
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
