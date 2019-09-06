package sample

/**
 * 1. init()最先调用
 * 2. 多重赋值
 * 3. 匿名变量
 * 4. 数组传参，传拷贝
 *
 * 5. panic() / recover() recover仅在延迟函数defer中有效。
 *     Go没有try catch，recover类似于catch
 * */

import "fmt"

// NUM is a const value
const NUM = 10

func init() {
	fmt.Println("> sample package init")
}

// VariablesTest test
func VariablesTest() {
	var i = 10
	var j = 1
	k := 5
	// 多重复值
	j, k = k, j
	fmt.Println("Hello, 世界！", i, j, k)

	//匿名变量
	_, _, nickName := getName()
	fmt.Printf("pointer: %p, value:%s\n", &nickName, nickName)

	//数组传值，传备份
	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println(&(array), array)
}

func getName() (firstName int, lastName, nickName string) {
	// 常量定义
	return NUM, "Chan", "Chibi Maruko"
}

func modify(array [5]int) {
	array[0] = 10
	fmt.Println(&array, array)
}
