package main_test

import "fmt"

//map没有这个key时, 访问这个key的value就会返回value类型的默认值

// func main() {
// 	test1()
// 	test2()
// 	test3()
// }

func test1() {
	m := make(map[int]string)
	m[1] = "adf"
	if m[100] == "" {
		fmt.Println("string默认类型: 空")
	}
}

func test2() {
	m := make(map[int]bool)
	m[1] = true
	if m[100] == false {
		fmt.Println("bool默认类型: false")
	}
}

func test3() {
	m := make(map[int]int)
	m[1] = 1
	if m[100] == 0 {
		fmt.Println("int默认类型: 0")
	}
}
