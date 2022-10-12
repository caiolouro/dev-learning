package main

import (
	"fmt"
	"leetcode/rotatearray"
)

func main() {
	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	// k := 3
	nums := []int{-1, -100, 3, 99}
	k := 6
	// nums := []int{1, 2, 3}
	// k := 2
	fmt.Println("Input nums:", nums, ", k:", k)

	// rotatearray.UsingCopy(nums, k)
	rotatearray.UsingReverse(nums, k)
	fmt.Println("Rotated nums:", nums)
}
