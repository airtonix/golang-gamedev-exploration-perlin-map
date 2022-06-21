package array

import (
	"github.com/samber/lo"
)

func Pop[T comparable](alist *[]T) T {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

func Prepend[T comparable](list []T, item T) []T {
	return append([]T{item}, list...)
}

func Remove[T comparable](list []T, itemToRemove T) []T {
	return lo.Filter(list, func(x T, _ int) bool {
		return x != itemToRemove
	})
}

func VectorEach[T comparable](
	list []T,
	size int,
	process func(x int, y int, item T),
) {
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			process(x, y, list[x*size+y])
		}
	}
}
