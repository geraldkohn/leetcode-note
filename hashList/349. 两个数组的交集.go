package main

/**
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。

示例 1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]

示例 2：
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的
*/

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	ret := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]] = true
	} 
	for i := 0; i < len(nums2); i++ {
		m2[nums2[i]] = true
	}
	for k1, _ := range m1 {
		for k2, _ := range m2 {
			if k1 == k2 {
				ret = append(ret, k2)
				continue
			}
		}
	}
	return ret
}
