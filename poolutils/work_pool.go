package poolutils

import (
	"sync"

	"github.com/panjf2000/ants/v2"
)

var pools sync.Map

func NewWorkPool(workName string, size int, opeFunc func(i interface{}), options ...ants.Option) error {

	if size <= 0 {
		size = 100
	}

	if _, ok := pools.Load(workName); ok {
		return nil
	}

	pool, err := ants.NewPoolWithFunc(size, opeFunc, options...)
	if err != nil {
		return err
	}

	pools.Store(workName, pool)

	return nil

}

func WPool(workName string) *ants.PoolWithFunc {

	pool, ok := pools.Load(workName)
	if !ok {
		return nil
	}

	return pool.(*ants.PoolWithFunc)
}
