package sample

/**
 * Map 在多线程中，同时读写是线程不安全的。需要用sync.Map
 * 创建
 * 		make(map[string] string)
 * 创建+初始化
 * 		m := map[string]string{}
 * 取值
 * 		if v, ok := m1["a"]; ok {}
 * 遍历
 * 		for k, v := range m1 {
 * 删除
 * 		delete
 * */

import "fmt"

// PersonInfo 是一个包含个人详细信息的类型
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

// MapTest test map
func MapTest() {
	var personDB map[string]PersonInfo //声明
	personDB = make(map[string]PersonInfo)

	personDB["1234"] = PersonInfo{
		ID:      "1234",
		Name:    "Tom",
		Address: "Room 203",
	}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101"}

	person, ok := personDB["1234"]
	// ok是一个返回的bool型，返回true表示找到了对应的数据
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}
	delete(personDB, "1234")
}
