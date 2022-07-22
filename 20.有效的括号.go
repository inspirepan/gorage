/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

package leetcode

// @lc code=start
func isValid(s string) bool {
	var stack = []rune{}
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			if char == '}' && stack[len(stack)-1] == '{' || char == ']' && stack[len(stack)-1] == '[' || char == ')' && stack[len(stack)-1] == '(' {

				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end

// test case need cover with quote
// ""()()()"\n"((()"\n"()()))""
