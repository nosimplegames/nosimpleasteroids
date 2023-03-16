package utils

type List[T any] struct {
	Elements []T

	Current *T
}
