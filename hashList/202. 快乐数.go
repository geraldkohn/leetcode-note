package main

import "strconv"

/**
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：
对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。
*/

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] { //防止死循环, 已经有过了就不能再循环了
		n, m[n] = getSum1(n), true
	}
	return n == 1
}

func getSum1(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

//以下是复杂的方法(自己写的)

func isHappySpiltInt(n int) []int {
	ret := make([]int, 0)
	str := strconv.Itoa(n)
	for i := 0; i < len(str); i++ {
		retElement, _ := strconv.Atoi(string(str[i]))
		ret = append(ret, retElement)
	}
	return ret
}

func isHappyArraySum(a *[]int) int {
	sum := 0
	for i := 0; i < len(*a); i++ {
		sum += (*a)[i]
	}
	return sum
}

func getSum2(n int) int {
	temp := isHappySpiltInt(n)
	sum := isHappyArraySum(&temp)
	return sum
}
