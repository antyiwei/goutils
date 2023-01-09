package poolutils

import (
	"log"
	"testing"
	"time"
)

func TestWorkPool(t *testing.T) {

	var opeFunc = func(i interface{}) {

		td := (i.(int) % 3)
		time.Sleep(time.Duration(td) * time.Second)
		log.Println(i)
	}
	workName := "task1"
	NewWorkPool(workName, 20, opeFunc)

	workName2 := "task2"
	NewWorkPool(workName2, 200, opeFunc)

	go func() {
		for i := 0; i < 1000; i++ {
			WPool(workName).Invoke(i)
		}
	}()

	go func() {
		for i := 1000; i < 2000; i++ {
			WPool(workName2).Invoke(i)
		}
	}()
	time.Sleep(1 * time.Hour)
}
