package leetcode

import "testing"

func Test_maxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	if r := maxSubArray(nums); r != 6 {
		t.Log(r)
		t.Fail()
	}
}
