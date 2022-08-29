package batchutils

func GroupRun(count, mod int, f func(start, end int) error) error {
	group := count / mod
	if count%mod != 0 {
		group += 1
	}
	for i := 0; i < group; i++ {
		start, end := i*mod, mod*(i+1)
		if i == group-1 && count%mod != 0 {
			end = count
		}
		err := f(start, end)
		if err != nil {
			return err
		}
	}
	return nil
}
