package main

/**
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
如果可以，返回 true ；否则返回 false 。
magazine 中的每个字符只能在 ransomNote 中使用一次。
*/

func canConstruct(ransomNote string, magazine string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	byteCount(ransomNote, &m1)
	byteCount(magazine, &m2)
	for k1, v1 := range m1 {
		v2 := m2[k1]
		if v1 > v2 {
			return false
		}
	}
	return true
}

func byteCount(str string, m *map[byte]int) {
	for i := 0; i < len(str); i++ {
		(*m)[str[i]]++
	}
}