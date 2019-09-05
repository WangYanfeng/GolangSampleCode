package sample

/**
 * 1. slice create:
 *     []int{*,*,*}, make([]int, len, cap)
 * 2. append(slice, *)
 * 3. len(), cap()
 * 4. 切片[:]
 * 5. for range 创建每个元素的副本*/

import "fmt"

//SliceTest test
func SliceTest() {
	slice := make([]int, 1, 1) //slice := []int{0, 1, 2, 3, 4}
	slice[0] = 1
	slice = append(slice, 2)
	fmt.Printf("Type: %T\n", slice)

	for index, value := range slice {
		fmt.Printf("slice[%d] address:%p. range address:%p.\n", index, &slice[index], &value)
	}

	slice = append(slice, 3)
	printSlice(slice[1:3])
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
