package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	Nums []int
	K    int
	Want []int
}

func TestRotateWithCopy(t *testing.T) {
	var testCases = []TestCase{
		{
			Nums: []int{1, 2, 3, 4, 5, 6, 7},
			K:    3,
			Want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			Nums: []int{1, 2, 3},
			K:    2,
			Want: []int{2, 3, 1},
		},
		{
			Nums: []int{-1, -100, 3, 99},
			K:    6,
			Want: []int{3, 99, -1, -100},
		},
	}

	for _, testCase := range testCases {
		rotateWithCopy(testCase.Nums, testCase.K)
		if !reflect.DeepEqual(testCase.Nums, testCase.Want) {
			t.Fatalf(`rotateWithCopy = %#v, want = %#v`, testCase.Nums, testCase.Want)
		}
	}
}

func TestRotateWithReverse(t *testing.T) {
	var testCases = []TestCase{
		{
			Nums: []int{1, 2, 3, 4, 5, 6, 7},
			K:    3,
			Want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			Nums: []int{1, 2, 3},
			K:    2,
			Want: []int{2, 3, 1},
		},
		{
			Nums: []int{-1, -100, 3, 99},
			K:    6,
			Want: []int{3, 99, -1, -100},
		},
	}

	for _, testCase := range testCases {
		rotateWithReverse(testCase.Nums, testCase.K)
		if !reflect.DeepEqual(testCase.Nums, testCase.Want) {
			t.Fatalf(`rotateWithReverse = %#v, want = %#v`, testCase.Nums, testCase.Want)
		}
	}
}
