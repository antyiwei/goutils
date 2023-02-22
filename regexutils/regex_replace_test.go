package regexutils

import (
	"log"
	"testing"
)

func TestReplaceStringByRegex(t *testing.T) {
	orgStr := `<p>直接点击<a href="https://www.yidaoerp.cn/user/login" target="_blank">立即注册</a>登陆使用</p>`
	dstStr, err := ReplaceStringByRegex(orgStr, "<[^a>]+>", "")
	log.Println(dstStr, err)
}
