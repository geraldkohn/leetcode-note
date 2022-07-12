package main

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

示例 1：
输入：s = "()"
输出：true

示例 2：
输入：s = "()[]{}"
输出：true

示例 3：
输入：s = "(]"
输出：false

示例 4：
输入：s = "([)]"
输出：false

示例 5：
输入：s = "{[]}"
输出：true
*/

func isValid(s string) bool {
	leftStack := make([]byte, 0)
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] == '(' || b[i] == '[' || b[i] == '{' {
			leftStack = append(leftStack, b[i])
		} else {
			if len(leftStack) == 0 {
				return false
			}
			last := leftStack[len(leftStack)-1]
			leftStack = leftStack[:len(leftStack)-1]
			switch b[i] {
			case ')':
				if last != '(' {
					return false
				}
			case ']':
				if last != '[' {
					return false
				}
			case '}':
				if last != '{' {
					return false
				}
			}
		}
	}
	return len(leftStack) == 0
}

// func main() {
// 	s := "("
// 	ok := isValid(s)
// 	fmt.Println(ok)
// }
