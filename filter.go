package filter

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type AllowedTypes interface {
	constraints.Ordered
}

func Filter[T AllowedTypes](slice1, slice2 []T) []T {
	combinedSlice := append(slice1, slice2...)
	temp := make(map[T]bool)
	fmt.Println("Filter putting in work")
	var filtered []T
	for _, v := range combinedSlice {
		_, ok := temp[v]
		if !ok {
			temp[v] = true
			filtered = append(filtered, v)
		}
	}
	return filtered
}
