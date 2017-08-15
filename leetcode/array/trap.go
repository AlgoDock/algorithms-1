/* https://leetcode.com/problems/trapping-rain-water/description/
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.

For example,
Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.

http://www.leetcode.com/static/images/problemset/rainwatertrap.png
*/

package leetcode

func trap(height []int) int {
	// 2 pointers
	var (
		left, right       = 0, len(height) - 1
		res               = 0
		leftMax, rightMax = 0, 0
	)

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}
