package regexutils

import (
	"errors"
	"regexp"
)

func ReplaceStringByRegex(str, rule, replace string) (string, error) {
	reg, err := regexp.Compile(rule)
	if reg == nil || err != nil {
		return "", errors.New("正则MustCompile错误:" + err.Error())
	}

	//log.Println(reg.FindString(str))
	return reg.ReplaceAllString(str, replace), nil
}
