package main

/**
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

示例 1：
输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
输出：2
解释：
两个元组如下：
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0

示例 2：
输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
输出：1
*/

// func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
// 	count := 0
// 	for i1 := 0; i1 < len(nums1); i1++ {
// 		for i2 := 0; i2 < len(nums2); i2++ {
// 			for i3 := 0; i3 < len(nums3); i3++ {
// 				for i4 := 0; i4 < len(nums4); i4++ {
// 					if nums1[i1] + nums2[i2] + nums3[i3] + nums4[i4] == 0 {
// 						count++
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return count
// }
//以上时间超限 O(n^4)

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m1 := make(map[int]int) //两项和出现的次数
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			m1[nums1[i]+nums2[j]]++
		}
	}
	m2 := make(map[int]int)
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			m2[nums3[i]+nums4[j]]++
		}
	}
	count := 0
	for sum1, count1 := range m1 {
		for sum2, count2 := range m2 {
			if sum2+sum1 == 0 {
				count += count2 * count1
			}
		}
	}
	return count
}
