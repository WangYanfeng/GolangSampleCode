package sample

/**
 * FormatBool() / ParseBool()
 * Atoi() / Itoa()
 * FormatInt() / ParseInt()
 * FormatFloat() / ParseFloat()
 * Append...()
 * */

import (
	"fmt"
	"strconv"
)

// StrconvTest 字符串转换
func StrconvTest() {
	//获取程序运行的操作系统平台下 int 类型所占的位数，如：strconv.IntSize。
	//strconv.IntSize
	i, _ := strconv.Atoi("1024") //ParseInt("1024", 10, 0)
	fmt.Printf("\t'1024' Atoi is: %d\n", i)

	s := strconv.Itoa(1024)
	fmt.Printf("\t1024 Itoa is: %q\n", s)

	//Parse* 函数将字符串转换为其他类型
	f, _ := strconv.ParseFloat("10.24", 10)
	fmt.Printf("\t'10.24' ParseFloat is: %f\n", f)
	b, _ := strconv.ParseBool("true")
	fmt.Printf("\t'true' ParseBool is: %t\n", b) //ParseInt(s, base, bitSize)

	//Format* 函数将给定的类型变量转换为string返回
	ss := strconv.FormatBool(true)
	fmt.Printf("\ttrue FormatBool is: %s\n", ss)
	sss := strconv.FormatFloat(10.24, 'E', -1, 32)
	fmt.Printf("\t'10.24' FormatFloat is: %s\n", sss) //FormatInt(i, base)

	//Append* 函数表示将给定的类型(如bool, int等)转换为字符串后, 添加在现有的字节数组中[]byte
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 123, 10) // 10用来表示进制
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "andrew")
	str = strconv.AppendQuoteRune(str, '刘')
	fmt.Printf("\tAppend result is: %s\n", str) //FormatInt(i, base)
}
