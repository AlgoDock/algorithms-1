/* https://leetcode.com/problems/remove-duplicates-from-sorted-array/#/description
Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

For example,
Given input array nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the new length.
*/
package leetcode

func removeDuplicates(nums []int) int {
	n, end := len(nums), 0
	if n < 2 {
		return n
	}
	for i := 1; i < n; i++ {
		if nums[i] != nums[end] {
			end++
			nums[end] = nums[i]
		}
	}
	return end + 1
}
