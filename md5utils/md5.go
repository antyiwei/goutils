package md5utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func MD5(str string) string {
	w := md5.New()
	io.WriteString(w, str)               //将str写入到w中
	return fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
}

// 文件MD5,快速但占用内存高
func FileMD5Fast(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	all, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return MD5(string(all)), nil
}

// 文件MD5,大文件使用,内存占用低
func FileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}
