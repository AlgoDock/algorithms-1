/* https://leetcode.com/problems/word-break/description/
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words. You may assume the dictionary does not contain duplicate words.

For example, given
s = "leetcode",
dict = ["leet", "code"].

Return true because "leetcode" can be segmented as "leet code".

UPDATE (2017/1/4):
The wordDict parameter had been changed to a list of strings (instead of a set of strings). Please reload the code definition to get the latest changes.
*/

package ldp

func wordBreak(s string, wordDict []string) bool {
	maps := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		maps[word] = true
	}

	var helper func(s string) bool
	helper = func(s string) bool {
		if res, ok := maps[s]; ok {
			return res
		}

		for i := 1; i < len(s); i++ {
			if res, ok := maps[s[:i]]; ok && res {
				if word := s[i:]; helper(word) {
					maps[word] = true
					return true
				}
			}
		}
		maps[s] = false
		return false
	}

	return helper(s)
}
