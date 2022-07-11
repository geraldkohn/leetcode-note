package main

/**
给你一个字符串 s ，颠倒字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s 中可能会存在前导空格、尾随空格或者单词间的多个空格。
返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
*/

func reverseWords(s string) string {
	b := make([]byte, 0)
	b = append(b, []byte(s)...)
	//用双指针法移除数组元素, 复杂度O(n)
	b = sub1Reverse(b)
	//反转字符串
	b = sub2Reverse(b)
	//反转单词
	b = sub3Reverse(b)
	return string(b)
}

func sub1Reverse(b []byte) []byte {
	//用双指针法移除数组元素, 复杂度O(n)
	left := 0  //这个就是用于更新元素的指针, 更新完left位置的元素left就加
	right := 0 //这个是用于查看哪个需要被删除的指针,
	//保留的元素需要赋值给left指针位置的元素
	count := 0
	for right < len(b) && left < len(b) {
		if b[right] == ' ' && count == 0 {
			b[left] = b[right]
			left++
			count++
		} else if b[right] != ' ' {
			b[left] = b[right]
			left++
			count = 0 //设置为零
		} else {
			count++ //空格位置的计数器加一
		}
		right++
	}
	b = b[:left]
	if b[len(b)-1] == ' ' {
		b = b[:len(b)-1]
	}
	if b[0] == ' ' {
		b = b[1:]
	}
	return b
}

func sub2Reverse(b []byte) []byte {
	left := 0
	right := len(b) - 1
	for left < right {
		temp := b[left]
		b[left] = b[right]
		b[right] = temp
		left++
		right--
	}
	return b
}

func sub3Reverse(b []byte) []byte {
	temp := make([]byte, 0)
	tempBeginIndex := 0
	for i := 0; i < len(b); i++ {
		if i == len(b)-1 { //到了最后一个的时候
			temp = append(temp, b[i]) //写入到临时数组里
			temp = sub2Reverse(temp)  //反转单词
			k := 0
			j := tempBeginIndex
			for j < tempBeginIndex+len(temp) && k < len(temp) {
				b[j] = temp[k]
				k++
				j++
			}
			break
		} else if b[i] == ' ' { //到了空格
			temp = sub2Reverse(temp) //反转单词
			k := 0
			j := tempBeginIndex
			for j < tempBeginIndex+len(temp) && k < len(temp) {
				b[j] = temp[k]
				k++
				j++
			}
			temp = make([]byte, 0) //重置
			tempBeginIndex = i + 1 //空格的下一个就是单词的起始位置
		} else if b[i] != ' ' {
			temp = append(temp, b[i]) //加入到临时数组
		}
	}
	return b
}

// func main() {
// 	s := []byte("    hello world    hi ")
// 	str := reverseWords(string(s))
// 	fmt.Println(str + "->end")
// }
