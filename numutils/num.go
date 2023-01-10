package numutils

import "strings"

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
