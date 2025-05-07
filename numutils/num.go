package numutils

import (
	"log"
	"regexp"
	"strings"
)

func SubPhone(phone string) string {

	if strings.Contains(phone, "86") && phone[0:2] == "86" {
		return strings.Replace(phone, "86", "", 1)
	}

	if strings.Contains(phone, "+86") && phone[0:3] == "+86" {
		return strings.Replace(phone, "+86", "", 1)
	}

	if strings.Contains(phone, "%2B86") && phone[0:5] == "%2B86" {
		return strings.Replace(phone, "%2B86", "", 1)
	}

	if strings.Contains(phone, "0086") && phone[0:4] == "0086" {
		return strings.Replace(phone, "0086", "", 1)
	}
	return phone
}

const (
	NumberReg = `^0|9\d{2,3}-?\d{7,8}}|400-?[016789]-?\d{6}$`                 // 固话
	MobileRex = `^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$` // 手机号
)

func ValidateNumber(number string) bool {
	result := ValidateNum(number)
	if result {
		return true
	}

	result = ValidateMobile(number)
	return result
}

func ValidateNum(number string) bool {
	reg := regexp.MustCompile(NumberReg)
	result := reg.MatchString(number)
	return result
}

func ValidateMobile(number string) bool {
	reg := regexp.MustCompile(MobileRex)
	result := reg.MatchString(number)
	return result
}

func ParseCallee(callee string) (string, string) {
	callees := strings.Split(callee, "+")
	if len(callees) == 2 {
		return callees[0], callees[1]
	}

	re, _ := regexp.Compile("[^a-zA-Z0-9]")
	callee = re.ReplaceAllString(callee, "")
	lenstr := len(callee)
	if lenstr < 12 {
		log.Println("callee not parsed:", callee)
		return "", callee
	}

	start := lenstr - 11
	if ValidateMobile(callee[start:]) {
		return callee[:start], callee[start:]
	}

	start = lenstr - 12
	if ValidateNum(callee[start:]) {
		return callee[:start], callee[start:]
	}
	return "", callee
}
