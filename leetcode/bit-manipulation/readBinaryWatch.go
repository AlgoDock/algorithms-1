/* https://leetcode.com/problems/binary-watch/#/description
A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59).

Each LED represents a zero or one, with the least significant bit on the right.

https://upload.wikimedia.org/wikipedia/commons/8/8b/Binary_clock_samui_moon.jpg

For example, the above binary watch reads "3:25".

Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.

Example:

    Input: n = 1
    Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
Note:
    The order of output does not matter.
    The hour must not contain a leading zero, for example "01:00" is not valid, it should be "1:00".
    The minute must be consist of two digits and may contain a leading zero, for example "10:2" is not valid, it should be "10:02".
*/

package lbm

import "fmt"

func readBinaryWatch(num int) []string {
	H, S := make([][]int, 4), make([][]int, 6)
	for i := 0; i < 12; i++ {
		c := countBits(i)
		H[c] = append(H[c], i)
		S[c] = append(S[c], i)
	}
	for i := 12; i < 60; i++ {
		c := countBits(i)
		S[c] = append(S[c], i)
	}

	r := []string{}
	for i := 0; i <= num; i++ {
		for h := 0; i < 4 && h < len(H[i]); h++ {
			for s := 0; num-i < 6 && s < len(S[num-i]); s++ {
				r = append(r, fmt.Sprintf("%d:%02d", H[i][h], S[num-i][s]))
			}
		}

	}
	return r
}

func countBits(n int) int {
	c := 0
	for n > 0 {
		n = n & (n - 1)
		c++
	}
	return c
}
