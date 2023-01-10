package numutils

import (
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
	result := validateNum(number)
	if result {
		return true
	}

	result = validateMobile(number)
	if result {
		return true
	}
	return false
}

func validateNum(number string) bool {
	reg := regexp.MustCompile(NumberReg)
	result := reg.MatchString(number)
	if result {
		return true
	}
	return false
}

func validateMobile(number string) bool {
	reg := regexp.MustCompile(MobileRex)
	result := reg.MatchString(number)
	if result {
		return true
	}
	return false
}
