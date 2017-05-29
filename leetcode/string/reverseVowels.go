/* https://leetcode.com/problems/reverse-vowels-of-a-string/#/description
Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:
Given s = "hello", return "holle".

Example 2:
Given s = "leetcode", return "leotcede".

Note:
The vowels does not include the letter "y".
*/
package leetcode

func isVowels(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func reverseVowels(s string) string {
	r := []rune(s)
	i, j := 0, len(r)-1
	for i < j {
		for i < j && !isVowels(r[i]) {
			i++
		}
		for i < j && !isVowels(r[j]) {
			j--
		}
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}
