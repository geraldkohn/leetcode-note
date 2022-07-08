package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	index := 0
	sort.Ints(nums)
	ret := make([][]int, 0)
	retIndex := 0
	for index <= len(nums)-3 { //固定一个指针, 从头遍历. 这里是为了给左右指针留出位置
		left := index + 1
		right := len(nums) - 1
		for left < right { //左指针小于右指针, 才能有三个不同的数
			if nums[index]+nums[left]+nums[right] < 0 { //和小于0, 就需要把左指针向右移动
				//跳过重复的值
				newLeft := left
				//检查数组是否越界
				for newLeft != right && nums[newLeft] == nums[left] {
					newLeft++
				}
				if newLeft == right {
					break
				}
				left = newLeft
			} else if nums[index]+nums[left]+nums[right] > 0 { //和大于0, 就需要把右指针先左移动
				//跳过重复的值
				newRight := right
				//检查数组是否越界
				for newRight != left && nums[newRight] == nums[right] {
					newRight--
				}
				if newRight == left {
					break
				}
				right = newRight
			} else {
				ret = append(ret, make([]int, 0))
				ret[retIndex] = append(ret[retIndex], nums[index], nums[left], nums[right])
				sort.Ints(ret[retIndex])
				retIndex++

				//跳过重复的值
				newLeft := left
				//检查数组是否越界
				for newLeft != right && nums[newLeft] == nums[left] {
					newLeft++
				}
				if newLeft == right {
					break
				}
				left = newLeft
				//跳过重复的值
				newRight := right
				//检查数组是否越界
				for newRight != left && nums[newRight] == nums[right] {
					newRight--
				}
				if newRight == left {
					break
				}
				right = newRight
			}
		}
		newIndex := index
		//检查数组是否越界
		for newIndex != len(nums) && nums[newIndex] == nums[index] {
			newIndex++
		}
		if newIndex == len(nums) {
			break
		}
		index = newIndex
	}
	return ret
}

// func main() {
// 	// m := []int{-1, 0, 1, 2, -1, -4}
// 	// m := []int{}
// 	// m := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
// 	m := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
// 	ret := threeSum(m)
// 	fmt.Println(ret)
// }
