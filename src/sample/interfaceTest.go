package sample

/**
 * 当方法是 “指针接受者” 时，Interface必须用指针
 * 当方法是 “ 值 接受者” 时，Interface都可以
 * */
import (
	"fmt"
)

type mInterface interface {
	getName() string
	setName(string)
}

type mStruct struct {
	name string
}

func (m *mStruct) setName(s string) {
	m.name = s
}

func (m *mStruct) getName() string {
	return m.name
}

// InterfaceTest 接口必须是指针
func InterfaceTest() {
	var s mStruct
	fmt.Printf("%T\n", s)

	var i mInterface = &s
	i.setName("Tom")
	fmt.Printf("name: %s\n", i.getName())

	i = new(mStruct)
	i.setName("Jerry")
	fmt.Printf("name: %s\n", i.getName())
}
