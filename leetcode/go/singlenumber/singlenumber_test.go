package singlenumber

import "testing"

type TestCase struct {
	Nums []int
	Want int
}

func TestUsingMap(t *testing.T) {
	testCases := []TestCase{
		{
			Nums: []int{2, 2, 1},
			Want: 1,
		},
		{
			Nums: []int{4, 1, 2, 1, 2},
			Want: 4,
		},
		{
			Nums: []int{1},
			Want: 1,
		},
	}

	for _, testCase := range testCases {
		result := UsingMap(testCase.Nums)
		if result != testCase.Want {
			t.Fatalf(`UsingMap(%#v) = %d, want = %d`, testCase.Nums, result, testCase.Want)
		}
	}
}

func TestUsingXor(t *testing.T) {
	testCases := []TestCase{
		{
			Nums: []int{2, 2, 1},
			Want: 1,
		},
		{
			Nums: []int{4, 1, 2, 1, 2},
			Want: 4,
		},
		{
			Nums: []int{1},
			Want: 1,
		},
	}

	for _, testCase := range testCases {
		result := UsingXor(testCase.Nums)
		if result != testCase.Want {
			t.Fatalf(`UsingMap(%#v) = %d, want = %d`, testCase.Nums, result, testCase.Want)
		}
	}
}
