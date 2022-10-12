package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	// k := 3
	nums := []int{-1, -100, 3, 99}
	k := 6
	// nums := []int{1, 2, 3}
	// k := 2
	fmt.Println("Input nums:", nums, ", k:", k)

	rotateWithCopy(nums, k)
	// rotateWithReverse(nums, k)
	fmt.Println("Rotated nums:", nums)
}

// 36ms (beats 80,06%), 8MB (beats 91,21%)
func rotateWithCopy(nums []int, k int) {
	oldNums := make([]int, len(nums))
	copy(oldNums, nums)

	for i := 0; i < len(oldNums); i++ {
		newIndex := (i + k) % len(oldNums)
		nums[newIndex] = oldNums[i]
	}
}

// 69ms (beats 34,12%), 8,3MB (beats 64,37%)
func rotateWithReverse(nums []int, k int) {
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k%len(nums)-1)
	reverse(nums, k%len(nums), len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start += 1
		end -= 1
	}
}
