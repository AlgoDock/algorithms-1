package lmath

import "testing"

func Test_isUgly(t *testing.T) {
	nums := []int{1, 14, 7, 30, 123456789}
	results := []bool{true, false, false, true, false}

	for i := range nums {
		if r := isUgly(nums[i]); r != results[i] {
			t.Fatal(nums[i], r)
		}
	}
}
