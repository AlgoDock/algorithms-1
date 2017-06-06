/* https://leetcode.com/problems/heaters/#/description
Winter is coming! Your first job during the contest is to design a standard heater with fixed warm radius to warm all the houses.

Now, you are given positions of houses and heaters on a horizontal line, find out minimum radius of heaters so that all houses could be covered by those heaters.

So, your input will be the positions of houses and heaters seperately, and your expected output will be the minimum radius standard of heaters.

Note:
    Numbers of houses and heaters you are given are non-negative and will not exceed 25000.
    Positions of houses and heaters you are given are non-negative and will not exceed 10^9.
    As long as a house is in the heaters' warm radius range, it can be warmed.
    All the heaters follow your radius standard and the warm radius will the same.

Example 1:
    Input: [1,2,3],[2]
    Output: 1
    Explanation: The only heater was placed in the position 2, and if we use the radius 1 standard, then all the houses can be warmed.

Example 2:
    Input: [1,2,3,4],[1,4]
    Output: 1
    Explanation: The two heater was placed in the position 1 and 4. We need to use radius 1 standard, then all the houses can be warmed.
*/

package leetcode

import "sort"

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)

	maxDist := 0
	for _, house := range houses {
		index := findHeater(heaters, house)
		curDist := abs(heaters[index] - house)
		if index+1 < len(heaters) {
			if temp := abs(heaters[index+1] - house); temp < curDist {
				curDist = temp
			}
		}
		if curDist > maxDist {
			maxDist = curDist
		}
	}
	return maxDist
}

func findHeater(heaters []int, house int) int {
	// 找到house左右两边的heater，永远返回左边
	start, end := 0, len(heaters)-1

	if house <= heaters[start] {
		return start
	} else if house >= heaters[end] {
		return end
	}

	for start <= end {
		mid := start + (end-start)/2
		switch {
		case house < heaters[mid]:
			end = mid - 1
		case house > heaters[mid]:
			start = mid + 1
		case house == heaters[mid]:
			return mid
		}
	}
	return start - 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
