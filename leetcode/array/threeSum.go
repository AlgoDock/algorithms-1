/* https://leetcode.com/problems/3sum/#/description
Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note: The solution set must not contain duplicate triplets.

For example, given array S = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

package leetcode

import "sort"
import "fmt"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	maps := make(map[string]int)
	result := [][]int{}

	for i := 1; i < len(nums)-1; i++ {
		left, right := i-1, i+1
		for 0 <= left && right < len(nums) {
			if temp := nums[left] + nums[i] + nums[right]; temp == 0 {
				tempr := []int{nums[left], nums[i], nums[right]}
				key := fmt.Sprintln(tempr)
				if _, ok := maps[key]; !ok {
					result = append(result, tempr)
					maps[key] = 1
				}
				for left > 0 && nums[left] == nums[left-1] {
					left-- // skip same result
				}
				for right < len(nums)-1 && nums[right] == nums[right+1] {
					right++ // skip same result
				}
				right++
				left--
			} else if temp < 0 {
				right++
			} else {
				left--
			}
		}
	}

	return result
}
