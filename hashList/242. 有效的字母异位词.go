package main

/**
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

示例 1:
输入: s = "anagram", t = "nagaram"
输出: true

示例 2:
输入: s = "rat", t = "car"
输出: false

提示:
1 <= s.length, t.length <= 5 * 104
s 和 t 仅包含小写字母
*/

func isAnagram(s string, t string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m1[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		m2[t[i]]++
	}
	for i := byte('a'); i <= byte('z'); i++ {
		cnt1, ok1 := m1[i]
		cnt2, ok2 := m2[i]
		if ok1 == ok2 { //都没有这个字母或者都有这个字母
			if cnt1 == cnt2 { //都没有数量都是默认的0, 都有数量需要相等
				continue
			} else { //数量不相等, 输出false
				return false
			}
		} else { //一个有一个没有, 就返回false
			return false
		}
	}
	return true
}
