package intersect

import (
	"reflect"
	"testing"
)

type TestCase struct {
	Nums1 []int
	Nums2 []int
	Want  []int
}

func TestIntersect(t *testing.T) {
	testCases := []TestCase{
		{
			Nums1: []int{1, 2, 2, 1},
			Nums2: []int{2, 2},
			Want:  []int{2, 2},
		},
		{
			Nums1: []int{4, 9, 5},
			Nums2: []int{9, 4, 9, 8, 4},
			Want:  []int{4, 9},
		},
	}

	for _, testCase := range testCases {
		result := Intersect(testCase.Nums1, testCase.Nums2)
		if !reflect.DeepEqual(result, testCase.Want) {
			t.Fatalf(`Intersect(%#v, %#v), want = %#v`, testCase.Nums1, testCase.Nums2, testCase.Want)
		}
	}
}
