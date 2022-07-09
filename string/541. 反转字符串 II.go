package main

/**
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
*/

func reverseStr(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(b); i += (2 * k) {
		if i+k >= len(b) {
			reverse(b, i, len(b)-1)
		} else if (i+2*k) >= len(b) && i+k <= len(b)-1 {
			reverse(b, i, i+k-1)
		} else {
			reverse(b, i, i+k-1)
		}
	}
	return string(b)
}

func reverse(s []byte, begin int, end int) { //左闭右闭
	left := begin
	right := end
	for left < right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}
