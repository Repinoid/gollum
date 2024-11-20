package gollum

import (
	"math/rand/v2"
)

func CreateArray(arrayLength, maxint int) (arra []int) {
	arra = make([]int, arrayLength)
	for i := range arra {
		arra[i] = rand.IntN(maxint)
	}
	return
}
