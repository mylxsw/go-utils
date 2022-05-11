package ternary

func IfElse[T interface{}](condition bool, positive, negative T) T {
	if condition {
		return positive
	}

	return negative
}

func IfElseLazy[T interface{}](condition bool, positiveFunc, negativeFunc func() T) T {
	if condition {
		return positiveFunc()
	}

	return negativeFunc()
}
