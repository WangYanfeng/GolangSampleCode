package sample

/**
 * 1. 构造
 * 		struct{v1,v2, ...} / struct{key:value, ...}
 * 		new()
 * 		自定义构造函数
 * 2. 继承: 组合
 * 		匿名字段 父类，才可以直接引用 父类变量。
 * */

import (
	"fmt"
)

type people struct {
	name string
	age  int
}

func (p *people) setName(n string) {
	p.name = n
}

// Student class contain people class
type Student struct {
	people //匿名字段
	grade  int
}

// NewStudent : 构造函数
func NewStudent(funcs ...SetStuOptFunc) *Student {
	student := new(Student) //Student pointer
	for _, f := range funcs {
		f(student)
	}
	return student
}

// SetStuOptFunc :
type SetStuOptFunc func(*Student)

// SetGrade :
func SetGrade(g int) SetStuOptFunc {
	return func(s *Student) {
		s.grade = g
	}
}

// SetPeople :
func SetPeople(p people) SetStuOptFunc {
	return func(s *Student) {
		s.people = p
	}
}

// ClassTest : main call test
func ClassTest() {
	var tom = Student{
		people: people{
			name: "Tom",
			age:  10,
		},
		grade: 3,
	}

	fmt.Printf("%T : %+v\n", tom, tom)
	tom.setName("TTom")
	fmt.Printf("call father func. name: %s, age: %d\n", tom.name, tom.age)

	p := people{"Jerry", 16}
	setGrade := SetGrade(5)
	setPeople := SetPeople(p)
	jerry := NewStudent(setGrade, setPeople)
	fmt.Printf("%T : %+v\n", jerry, jerry)
}
