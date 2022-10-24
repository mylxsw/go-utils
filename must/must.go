package must

// Must 包装2个返回值的函数返回值，当 err 为非 nil 的时候自动触发 panic
func Must[T interface{}](res T, err error) T {
	if err != nil {
		panic(err)
	}

	return res
}

// Must 包装2个返回值的函数返回值，当 err 为非 nil 的时候自动触发 panic
func Must2[T interface{}, M interface{}](r1 T, r2 M, err error) (T, M) {
	if err != nil {
		panic(err)
	}

	return r1, r2
}

// NoError 包装一个返回值为 error 的函数，当 err 为非 nil 的时候自动触发 panic
func NoError(err error) {
	if err != nil {
		panic(err)
	}
}
