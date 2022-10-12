package singlenumber

// 28ms (beats 62,85%), 7,4MB (beats 9,52%)
func UsingMap(nums []int) int {
	counts := make(map[int]int)

	for _, num := range nums {
		count := counts[num]
		if count == 1 {
			delete(counts, num)
			continue
		}
		counts[num] = 1
	}

	var result int
	for num := range counts {
		result = num
		break
	}

	return result
}

// 21ms (beats 79,67%), 6,3MB (beats 62,85%)
func UsingXor(nums []int) int {
	result := 0

	for _, num := range nums {
		result ^= num
	}

	return result
}
