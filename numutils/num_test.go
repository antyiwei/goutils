package numutils

import "testing"

func TestSubPhone(t *testing.T) {
	t.Log("17140800086", SubPhone("17140800086"))
	t.Log("%2B8617140800086", SubPhone("%2B8617140800086"))
	t.Log("+8617140800086", SubPhone("+8617140800086"))
	t.Log("8617140800086", SubPhone("8617140800086"))

}
