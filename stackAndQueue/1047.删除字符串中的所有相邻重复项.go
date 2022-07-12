package main

/**
给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
在 S 上反复执行重复项删除操作，直到无法继续删除。
在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

示例：
输入："abbaca"
输出："ca"

解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。
*/

func removeDuplicates(s string) string {
	if len(s) < 2 {
		return ""
	}
	b := []byte(s)
	ok := false
	for !ok {
		b, ok = subF_1047(b)
	}
	return string(b)
}

func subF_1047(b []byte) ([]byte, bool) {
	left := 0
	right := 1
	for right < len(b) && left < len(b) {
		if b[left] != b[right] {
			left++
			right++
		} else {
			break
		}
	}

	if right == len(b) { //没有找到相同的字符
		return b, true
	}
	if right == len(b)-1 || left == 0 { //到两头了
		if left == 0 && right != len(b)-1 {
			return b[right+1:], false
		} else if left != 0 && right == len(b)-1 {
			return b[:left], false
		} else {
			return []byte{}, true
		}
	}
	newB := append(b[:left], b[right+1:]...)
	return newB, false
}

// func main() {
// 	// s := "ababbjkkop"
// 	// s := "aaaa"
// 	s := "abbaca"
// 	ret := removeDuplicates(s)
// 	fmt.Println(ret)
// }
