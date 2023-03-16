package utils

type GenericCallback[T interface{}] func(T)

func ForEachBackwards[T interface{}](array []T, callback GenericCallback[T]) {
	for i := len(array) - 1; i >= 0; i-- {
		callback(array[i])
	}
}
