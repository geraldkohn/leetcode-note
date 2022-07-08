package main

import (
	"fmt"
	"sort"
)

/**
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）
*/

//用双指针法
func fourSum(nums []int, target int) [][]int {
	firstIndex := 0
	secondIndex := firstIndex + 1
	ret := make([][]int, 0)
	retIndex := 0
	sort.Ints(nums)
	for firstIndex <= len(nums)-4 {
		for secondIndex <= len(nums)-3 {
			left := secondIndex + 1
			right := len(nums) - 1
			for left < right {
				if nums[firstIndex]+nums[secondIndex]+nums[left]+nums[right] < target {
					newLeft := left
					for newLeft != right && nums[newLeft] == nums[left] {
						newLeft++
					}
					if newLeft == right {
						break
					}
					left = newLeft
				} else if nums[firstIndex]+nums[secondIndex]+nums[left]+nums[right] > target {
					newRight := right
					for newRight != left && nums[newRight] == nums[right] {
						newRight--
					}
					if newRight == left {
						break
					}
					right = newRight
				} else {
					ret = append(ret, make([]int, 0))
					ret[retIndex] = append(ret[retIndex], nums[firstIndex], nums[secondIndex], nums[left], nums[right])
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
			newSecondIndex := secondIndex
			//检查数组是否越界
			for newSecondIndex != len(nums)-2 && nums[newSecondIndex] == nums[secondIndex] {
				newSecondIndex++
			}
			if newSecondIndex == len(nums)-2 {
				break
			}
			secondIndex = newSecondIndex
		}
		newFirstIndex := firstIndex
		//检查数组是否越界
		for newFirstIndex != len(nums)-3 && nums[newFirstIndex] == nums[firstIndex] {
			newFirstIndex++
		}
		if newFirstIndex == len(nums)-3 {
			break
		}
		firstIndex = newFirstIndex
		secondIndex = firstIndex + 1
	}
	return ret
}

func main() {
	// m := []int{-1, 0, 1, 2, -1, -4}
	// m := []int{}
	// m := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	// m := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m := []int{1, 0, -1, 0, -2, 2}
	ret := fourSum(m, 0)
	fmt.Println(ret)
}
