package leetcode

import "testing"

func Test_strStr(t *testing.T) {
	haystacks := []string{"", "abc", "abcdabcdf"}
	needles := []string{"", "b", "abcdf"}
	results := []int{0, 1, 4}
	for i := 0; i < len(haystacks); i++ {
		if r := strStr(haystacks[i], needles[i]); r != results[i] {
			t.Log(haystacks[i], needles[i], results[i], "output is:", r)
			t.Fail()
		}
	}
}
