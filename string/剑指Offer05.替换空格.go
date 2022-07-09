package main

/**
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
*/

func replaceSpace(s string) string {
	temp := make([]byte, 2*len(s))
	b := []byte(s)
	b = append(b, temp...)
	left := len(s) - 1
	right := len(b) - 1
	for left >= 0 && right >= 0 {
		if b[left] == ' ' {
			b[right] = '0'
			b[right-1] = '2'
			b[right-2] = '%'
			left -= 1
			right -= 3
		} else {
			b[right] = b[left]
			left -= 1
			right -= 1
		}
	}
	if right < 0 {
		right = 0
	}
	b = b[right:]
	return string(b)
}

// func main() {
// 	// s := "we are happy."
// 	s := "              "
// 	newStr := replaceSpace(s)
// 	fmt.Println(newStr)
// }
