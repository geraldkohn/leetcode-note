package main

/**
实现 strStr() 函数。
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

说明：
当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。
*/

func strStr(haystack string, needle string) int {
	if needle == "" {
		return -1
	}
	if len(haystack) < len(needle) {
		return -1
	}
	haystackIndex := 0
	h := []byte(haystack)
	n := []byte(needle)
	lenN := len(n)
	for haystackIndex <= len(h)-lenN {
		ok := strEqual_28(h[haystackIndex:haystackIndex+lenN], n)
		if ok {
			return haystackIndex
		}
		haystackIndex++
	}
	return -1
}

//传入的数组大小需要相等
func strEqual_28(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// func main() {
// 	// h := "mississippi"
// 	// n := "issip"
// 	h := "woood"
// 	n := "wool"
// 	ret := strStr(h, n)
// 	fmt.Println(ret)
// }
