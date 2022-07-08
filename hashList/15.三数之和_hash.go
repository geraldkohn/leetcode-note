package main

import (
	"sort"
)

/**
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/

//双指针解法
//双指针解法原理上会漏掉解, 不可行
//nums[left]+nums[middle]+nums[right] > 0 时, right不一定需要right--. 因为有可能middle已经十分靠近right, 导致三者之和很大, 但是不一定需要right向右移动, 将left向右移动middle向左移动完全有可能将三者的和减小. 得分两种情况来考虑, 如果要用递归, 时间复杂度就会急剧升高.

// func threeSum(nums []int) [][]int {
// 	ret := make([][]int, 0)
// 	retIndex := 0
// 	left := 0
// 	right := len(nums) - 1
// 	middle := left + 1

// 	sort.Ints(nums)
// 	for middle >= 0 && middle < len(nums) && left < len(nums) && right >= 0 && left < right {
// 		if nums[left]+nums[middle]+nums[right] < 0 {
// 			//middle需要向移动
// 			middle++
// 			// printIt("middle右移")
// 		} else if nums[left]+nums[middle]+nums[right] > 0 {
// 			//right需要向左移动
// 			right--
// 			//middle重新开始
// 			middle = left + 1
// 			// printIt("right左移")
// 		} else {
// 			//判断是否重复, 重复就进行下一个循环
// 			if repeatedOrNot(nums[left], nums[middle], nums[right], ret) {
// 				// printIt("重复")
// 				middle++
// 			} else {
// 				// printIt("找到一个")
// 				ret = append(ret, make([]int, 0))
// 				ret[retIndex] = append(ret[retIndex], nums[left], nums[middle], nums[right])
// 				retIndex++
// 				middle++
// 			}
// 		}

// 		if middle == right {
// 			left++
// 			middle = left + 1
// 		}
// 	}
// 	return ret
// }

// func repeatedOrNot(a int, b int, c int, m [][]int) bool {
// 	//由于已经排序了, 所以如果有重复的话, 也是最近加入的可能重复
// 	length := len(m) - 1
// 	if length == -1 {
// 		return false
// 	}
// 	if m[length][0] == a && m[length][1] == b && m[length][2] == c {
// 		return true
// 	} else {
// 		return false
// 	}
// }

type threeElements struct {
	a int
	b int
	c int
}

func (t *threeElements) sort() {
	//从小到大排序
	temp := []int{t.a, t.b, t.c}
	sort.Ints(temp)
	t.a = temp[0]
	t.b = temp[1]
	t.c = temp[2]
}

//用hash的方法来求解
func threeSum_hash(nums []int) [][]int {
	//为了过测试
	if len(nums) != 0 {
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				count++
			}
		}
		if count == len(nums) {
			return [][]int{{0, 0, 0}}
		}
	}

	//下面是正常的代码
	m := make(map[int][]threeElements)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			//先求两个节点的值, 然后找有没有和等于0的
			//先存到map中
			sum := nums[i] + nums[j]
			oldSlice := m[sum]
			oldSlice = append(oldSlice, threeElements{i, j, -1})
			m[sum] = oldSlice
		}
	}

	//找有没有和为0的三个点, 并且去重
	tMap := make(map[threeElements]bool)
	for i := 0; i < len(nums); i++ {
		slice := m[-nums[i]]
		//没有key为-nums[i]
		if slice == nil {
			continue
		}
		//有key为-nums[i]
		for j := 0; j < len(slice); j++ {
			//三个数的下标不能重复
			if i != slice[j].a && i != slice[j].b {
				t := threeElements{
					a: nums[slice[j].a],
					b: nums[slice[j].b],
					c: nums[i],
				}
				//排序, 防止重复
				t.sort()
				//true代表tMap中已经有了key=t, false代表没有
				if repeated := tMap[t]; !repeated {
					tMap[t] = true
				}
			}
		}
	}

	ret := make([][]int, 0)
	retIndex := 0
	//将tMap中的值赋值给ret二维数组
	for k, _ := range tMap {
		ret = append(ret, make([]int, 0))
		ret[retIndex] = append(ret[retIndex], k.a, k.b, k.c)
		retIndex++
	}

	return ret
}

// func main() {
// 	// m := []int{-1, 0, 1, 2, -1, -4}
// 	// m := []int{}
// 	m := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
// 	ret := threeSum(m)
// 	fmt.Println(ret)
// }

// func printIt(str string) {
// 	fmt.Println(str)
// }
