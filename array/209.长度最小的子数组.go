package main

/**
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

示例：

输入：s = 7, nums = [2,3,1,2,4,3] 输出：2 解释：子数组 [4,3] 是该条件下的长度最小的子数组。

进阶：
如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(nlog(n)) 时间复杂度的解法。
*/

// 滑动窗口法
func minSubArrayLen(target int, nums []int) int {
	i := 0
	j := 0
	result := len(nums) + 1
	sum := 0
	for j < len(nums) {	//右指针从左遍历到右
		sum += nums[j]	//sum加上右指针位置的元素
		for sum >= target {	//只要满足和>=目标值, 左指针就会向右移动, 此时右指针停止移动
			result = min(result, j - i + 1)	//更新最短数组的长度
			sum -= nums[i]	//去掉左指针当前位置的值
			i++ //左指针移动
		}
		j++
	}
	if result == len(nums) + 1 {
		return 0
	} else {
		return result
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// func main() {
// 	a := minSubArrayLen(7, []int{2,3,1,2,4,3})
// 	fmt.Println(a)
// }
