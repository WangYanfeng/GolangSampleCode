package sample

/**
 * 1. strings
 *    转换：
 * 			toUpper() / ToLower() / Title() / ToTitle()
 *    比较:
 * 			Compare() / EqualFold()
 *    清理：
 *			Trim() / TrimLeft() / TrimRight() / TrimFunc()
			TrimSpace()
			TrimPrefix() / TrimSuffix()
 *    拆分：
 * 			Split() / SplitN() / SplitAfter() / SplitAfterN()
 * 			Fields() / FieldsFunc()
 *    合并
 * 			Join()
 * 			Repeat()
 *    子串
 * 			HasPrefix() / HasSuffix()
 * 			Contains() / ContainsRune() / ContainsAny()
 * 			Index() / IndexByte() / IndexRune() / IndexAny() / IndexFunc()
 * 			LastIndex() / LastIndexByte() / LastIndexAny() / LastIndexFunc()
 * 			Count()
 *    替换
 * 			Replace() / Map()
 * 2. string.Replacer
 *
 * 3. strings.Reader
*/

import (
	"fmt"
	"strings"
)

// StringsTest use strings package
func StringsTest() {
	// Compare比较字符串的速度比字符串内建(str1 == str2)要快
	var str1 = "hello world"
	var str2 = "world"
	res := strings.Compare(str1, str2)
	fmt.Printf("%q Compare %q: equal: %d\n", str1, str2, res)

	var isSub = strings.Contains(str1, str2)
	fmt.Printf("%q Contains %q: %t\n", str1, str2, isSub)

	var theIndex = strings.Index(str1, str2)
	fmt.Printf("%q Index %q is: %d\n", str1, str2, theIndex)

	fmt.Printf("'a,b,c' Split , : %q\n", strings.Split("a,b,c", ","))
	fmt.Printf("'a b c' Field : %q\n", strings.Fields("a b c"))

	fmt.Printf("%q Count 'l' is: %d\n", str1, strings.Count(str1, "l"))

	// 重复s字符串count次, 最后返回新生成的重复的字符串
	fmt.Printf("'%s' Repeat 5: %s\n", str2, strings.Repeat(str2, 5))

	// 在s字符串中, 把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	fmt.Printf("'%s' Replace o to *: %s\n", str1, strings.Replace(str1, "o", "*", -1))

	// 删除在s字符串的头部和尾部中指定的字符串, 并返回删除后的字符串
	fmt.Printf("',abc,' Trim , : %q\n", strings.Trim(",abc,", "c,"))
	fmt.Printf("' hello ' TrimSpace , : %q\n", strings.TrimSpace(" hello "))

	//判断字符串是否包含前缀prefix,后缀suffix
	fmt.Printf("'%s' HasPrefix 'wo' : %t\n", str2, strings.HasPrefix(str2, "wo"))
	fmt.Printf("'%s' HasSuffix 'ld' : %t\n", str2, strings.HasSuffix(str2, "ld"))

	//Join 用于将元素类型为 string 的 slice, 使用分割符号来拼接组成一个字符串
	fmt.Printf("%q Join / : %s\n", strings.Split("a,b,c", ","), strings.Join(strings.Split("a,b,c", ","), "/"))
}
