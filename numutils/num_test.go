package numutils

import (
	"log"
	"testing"
)

func TestSubPhone(t *testing.T) {
	t.Log("17140800086", SubPhone("17140800086"))
	t.Log("%2B8617140800086", SubPhone("%2B8617140800086"))
	t.Log("+8617140800086", SubPhone("+8617140800086"))
	t.Log("8617140800086", SubPhone("8617140800086"))

}

func TestValidateNumber(t *testing.T) {
	var number = "6602+057128135900"

	lenstr := len(number)
	start := lenstr - 11
	if ValidateMobile(number[start:]) {
		log.Println("是手机号")
	} else {
		if ValidateNum(number[start-1:]) {
			log.Println(number[start-1:], "是固话")
		}
	}

	got := ValidateNumber(number)
	log.Println(got)
}

func TestParseCallee(t *testing.T) {
	var number = "6602+057128135900"
	a, b := ParseCallee(number)
	log.Println(a, b)

}
