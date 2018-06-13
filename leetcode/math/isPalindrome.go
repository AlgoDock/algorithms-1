/* https://leetcode.com/problems/palindrome-number/#/description
Determine whether an integer is a palindrome. Do this without extra space.

palindrome: like level、noon.
*/
package lmath

func isPalindrome(x int) bool {
	result, n := 0, x
	for x > 0 {
		result = result*10 + x%10
		x /= 10
	}
	return result == n
}
