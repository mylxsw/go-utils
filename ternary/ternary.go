package ternary

func IfElse[T any](condition bool, positive, negative T) T {
	if condition {
		return positive
	}

	return negative
}

func If[T any](condition bool, positive, negative T) T {
	return IfElse(condition, positive, negative)
}

func IfLazy[T any](condition bool, positiveFunc, negativeFunc func() T) T {
	return IfElseLazy(condition, positiveFunc, negativeFunc)
}

func IfElseLazyPositive[T any](condition bool, positiveFunc func() T, negative T) T {
	if condition {
		return positiveFunc()
	}

	return negative
}

func IfElseLazyNegative[T any](condition bool, positive T, negativeFunc func() T) T {
	if condition {
		return positive
	}

	return negativeFunc()
}

func IfElseLazy[T any](condition bool, positiveFunc, negativeFunc func() T) T {
	if condition {
		return positiveFunc()
	}

	return negativeFunc()
}

func IfElseLazy2[T any, K any](condition bool, positiveFunc, negativeFunc func() (T, K)) (T, K) {
	if condition {
		return positiveFunc()
	}

	return negativeFunc()
}

func IfElseLazy3[T any, K any, M any](condition bool, positiveFunc, negativeFunc func() (T, K, M)) (T, K, M) {
	if condition {
		return positiveFunc()
	}

	return negativeFunc()
}
