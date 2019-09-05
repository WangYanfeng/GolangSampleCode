package sample

/**
 * 1. 双向链表list:
 *		New()
 *		PushBack() / PushFront()
 *		InsertAfter() / InsertBefore()
 *		Remove()
 *
 *		Front() / Back()
 *   list.Element:
 *		Next() Prev()
 * 2. 堆heap
 *
 * 3. 环形链表ring
 * */

import (
	"container/list"
	"fmt"
)

func listTest() {
	ll := list.New()
	ll.PushBack("Hi")
	ll.PushFront("OK")

	for p := ll.Front(); p != nil; p = p.Next() {
		//p *list.Element
		fmt.Println(p.Value)
	}
}

// ContainerTest : container package
func ContainerTest() {
	listTest()
}
