package sample

/**
 * TypeOf() :-> .Name() .NumField() .NumMethod() .Field() .Method()
 * ValueOf() :-> .Field() .FieldByName() .FieldByIndex() .MethodByName()
 * Elem()
 * CanSet() / Set() / CanAddr() / Addr()
 * New() / NewAt()
 * IsNil() / IsVaild()
 *
 * 从 Value 到实例
 *
 * 从 Value 的指针到值
 *
 *  Type 指针和值的相互转换
 *
 *
 * */

import (
	"fmt"
	"reflect"
)

type user struct { //定义一个结构类型
	ID   int
	Name string
	Age  int
}

func (u user) hello() { //定义一个结构方法
	fmt.Println("Hello world")
}

type manager struct {
	user  //嵌入User结构，User就是Manager结构的匿名字段
	title string
}

func info(o interface{}) { //定义一个方法，参数为空接口
	t := reflect.TypeOf(o)         //获取接收到的接口类型
	fmt.Println("Type:", t.Name()) //获取名称

	v := reflect.ValueOf(o) //获取接口的字段
	fmt.Println("Fields 获取结构字段:")

	//获取结构字段
	for i := 0; i < t.NumField(); i++ { //for循环，取出所拥有的字段
		f := t.Field(i)               //获取值字段
		val := v.Field(i).Interface() //获取字段的值
		//v.FieldByName("Id")
		//v.FieldByIndex(0)
		//v.MethodByName("Hello")
		fmt.Printf("%6s: %v=%v\n", f.Name, f.Type, val)
	}

	fmt.Println("Method 通过接口获取结构的方法:")
	//通过接口获取结构的方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}

	fmt.Println("--------------------")
	m := manager{user: user{1, "mm", 27}, title: "name"} //注意初始化方式
	t2 := reflect.TypeOf(m)                              //传递结构

	fmt.Printf("%#v\n", t2.Field(0))                  //获取索引为0的字段信息，即User字段信息
	fmt.Printf("%#v\n", t2.FieldByIndex([]int{0, 0})) //根据索引取出ID的字段信息（）
}

// ReflectTest :
func ReflectTest() {
	u := user{1, "OK", 12} //实例化一个结构
	info(u)                //调用Info函数
}
