/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

package leetcode

import "fmt"

// @lc code=start
func twoSum(nums []int, target int) []int {
	rec := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if value, ok := rec[nums[i]]; ok {
			return []int{value, i}
		}
		rec[target-nums[i]] = i
	}
	fmt.Println(rec)
	return []int{0, 0}
}

// @lc code=end
