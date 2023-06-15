package util

import (
	"fmt"
	"os"
)

// FileExists 判断文件夹是否存在
func FileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

// FileExistsOrCreate 文件是否存在，不存在就递归创建文件夹
func FileExistsOrCreate(dir string) error {
	if dir != "" {
		if !FileExists(dir) {
			// 创建文件夹
			return os.MkdirAll(dir, os.ModePerm)
		}
	}
	return fmt.Errorf("创建文件失败")
}
