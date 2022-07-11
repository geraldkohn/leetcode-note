package main

import "fmt"

/**
给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。

示例 1:
输入: s = "abab"
输出: true
解释: 可由子串 "ab" 重复两次构成。

示例 2:
输入: s = "aba"
输出: false

示例 3:
输入: s = "abcabcabcabc"
输出: true
解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)
*/

func repeatedSubstringPattern(s string) bool {
	b := []byte(s)
	for i := 0; i < len(b) - 1; i++ {	
		//这里的len(b) - 1是为了防止以自身为子字符串
		ok := subRepeated_459(b, i + 1)
		if ok {
			return true
		}
	}
	return false
}

//是否有重复的子字符串
//0到length左闭右开
func subRepeated_459(a []byte, length int) bool {
	if length > len(a) || len(a) % length != 0 {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i%length] != a[i] {
			return false
		}
	}
	return true
}

func main() {
	// s := "abababababab"
	s := " "
	ok := repeatedSubstringPattern(s)
	fmt.Println(ok)
}