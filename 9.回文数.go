/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

package leetcode

import (
	"strconv"
)

// @lc code=start
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

// @lc code=end
