/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

package leetcode

import "strings"

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	var builder strings.Builder
	flag := true
	for i := 0; i < minLen; i++ {
		curr := strs[0][i]
		for _, str := range strs {
			if str[i] != curr {
				flag = false
				break
			}
		}
		if !flag {
			return builder.String()
		}
		builder.WriteByte(curr)
	}
	return builder.String()
}

// @lc code=end
