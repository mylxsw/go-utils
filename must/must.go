package must

// Must 包装2个返回值的函数返回值，当 err 为非 nil 的时候自动触发 panic
func Must[T interface{}](res T, err error) T {
	if err != nil {
		panic(err)
	}

	return res
}
