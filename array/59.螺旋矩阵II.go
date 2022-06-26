package main

/**
给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
*/

func generateMatrix(n int) [][]int {
	//初始化
	array := make([][]int, n)
	for i := 0; i < n; i++ {
		array[i] = make([]int, n)
	}
	trackArray := make([][]int, n)
	for i := 0; i < n; i++ {
		trackArray[i] = make([]int, n)
	}
	x := 0
	y := 0
	trend := "up"
	counter := 1

	for trend != "over" {
		array[x][y] = counter
		counter++
		trackPass(&trackArray, x, y)
		x, y, trend = nextLocation(&trackArray, trend, x, y)
	}
	return array
}

// 标记走过的数组
func trackPass(trackArray *[][]int, x int, y int) {
	(*trackArray)[x][y] = 1
}

// 查看下一个位置在哪, 顺序就是: 右, 下, 左, 上
func nextLocation(trackArray *[][]int, trend string, x int, y int) (int, int, string) {
	for i := 1; i <= 4; i++ {
		switch trend {
		case "right":
			if y <= len(*trackArray)-2 && (*trackArray)[x][y+1] != 1 {
				return x, y + 1, "right"
			} else {
				trend = "down"
			}
		case "down":
			if x <= len(*trackArray)-2 && (*trackArray)[x+1][y] != 1 {
				return x + 1, y, "down"
			} else {
				trend = "left"
			}
		case "left":
			if y >= 1 && (*trackArray)[x][y-1] != 1 {
				return x, y - 1, "left"
			} else {
				trend = "up"
			}
		case "up":
			if x >= 1 && (*trackArray)[x-1][y] != 1 {
				return x - 1, y, "up"
			} else {
				trend = "right"
			}
		}
	}
	return -1, -1, "over"
}

// func main() {
// 	n := 6
// 	a := generateMatrix(n)
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			fmt.Print(a[i][j], " ")
// 		}
// 		fmt.Println()
// 	}
// }
