package sample

/**
 * 传入 []byte 的函数，都不会修改传入的参数。返回值要么是参数的副本，要么是参数的切片
 * 1. 字符串处理
 *    转换：
 * 			toUpper() / ToLower() / Title() / ToTitle()
 * 			ToUpperSpecial()...
 *    比较:
 * 			Compare() / Equal() / EqualFold()
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
 * 			Replace() / ReplaceAll() / Map()
 * 2. bytes.Buffer
 *    创建
 * 			NewBuffer() / NewBufferString()
 *    方法
 * 			Cap() / Len() / Bytes() / Grow()
 * 			Reset() / String()
 * 			Truncate()
 *    读
 * 			ReadBytes() / ReadString() / Next()
 * 			ReadFrom(reader)
 *    写
 * 			WriteString() / WriteRune()
 * 3. bytes.Reader
 *			Len()未读取长度 / Size()总长度 / Reset()
      读
			Read() / ReadAt()
      写
			WriteTo()
 * */

import (
	"bytes"
	"fmt"
)

func bytesFunc(b []byte) {
	//To Upper Lower
	fmt.Printf("ToUpper : %s\n", bytes.ToUpper(b))
	//Compare
	s1 := "Φφϕ kKK"
	s2 := "ϕΦφ KkK"
	fmt.Printf("EqualFold :%t\n", bytes.EqualFold([]byte(s1), []byte(s2)))
	//Trim
	fmt.Printf("Trim !: %s\n", bytes.Trim(b, "!"))
	fmt.Printf("TrimPrefix Hello: %s\n", bytes.TrimPrefix(b, []byte("Hello")))
	//Split
	fmt.Printf("Split: %q\n", bytes.Split(b, []byte{','}))
	fmt.Printf("Fields: %q\n", bytes.Fields(b))
	//Join
	fmt.Printf("Join -: %q\n", bytes.Join(bytes.Fields(b), []byte("-")))
	//Index
	fmt.Printf("Contains : %t\n", bytes.Contains(b, []byte("Hello")))
	fmt.Printf("Index : %d\n", bytes.Index(b, []byte("world")))
	//Replace
	fmt.Printf("Replace : %s\n", bytes.Replace(b, []byte("Hello"), []byte("Hi"), -1))
}

func bytesReader() {
	data := "Hello World!"
	re := bytes.NewReader([]byte(data))                             //通过[]byte创建Reader
	fmt.Printf("reader Len : %d, Size : %d\n", re.Len(), re.Size()) //Len()返回未读取部分的长度 Size()返回底层数据总长度

	buf := make([]byte, 6)
	for {
		//读取数据
		n, err := re.Read(buf)
		if err != nil {
			break
		}
		fmt.Printf("read %d byte from reader: %s\n", len(buf), string(buf[:n]))
		fmt.Printf("reader Len : %d, Size : %d\n", re.Len(), re.Size()) //Len()返回未读取部分的长度 Size()返回底层数据总长度
	}
	//b, err := re.ReadByte()//一个字节一个字节的读
	re.Seek(0, 0) //设置偏移量，因为上面的操作已经修改了读取位置等信息
	off := int64(0)
	for {
		n, err := re.ReadAt(buf, off) //指定偏移量读取
		if err != nil {
			break
		}
		off += int64(n)
		fmt.Printf("read %d at offset %d : %s\n", len(buf), off, string(buf[:n]))
	}
}

func bytesBuffer() {
	data := "Hello Go!"
	//通过[]byte创建一个Buffer
	buf := bytes.NewBuffer([]byte(data))

	//fmt.Printf("Buffer:%T\n", buf)
	fmt.Printf("%s Len: %d Cap: %d\n", buf.String(), buf.Len(), buf.Cap())

	bys := buf.Bytes() //Bytes()返回buffer中数据的切片
	for _, v := range bys {
		fmt.Printf(string(v) + "-")
	}
	fmt.Println()

	length := buf.Len()
	for i := 0; i < length; i++ {
		tmp := buf.Next(1) //Next()读出缓存中前n字节数据的切片
		fmt.Print(string(tmp) + "-")
	}

	buf.Reset() //重设缓冲，丢弃全部内容

	//通过string创建Buffer
	buf2 := bytes.NewBufferString("This is buffer2.")
	fmt.Printf("\n--------------------\n%s\n", buf2.String())
	line, _ := buf2.ReadBytes(' ') //读取第一个 分隔符 及其之前的内容，返回遇到的错误
	fmt.Println("ReadBytes before ' ':", string(line))
	//效果同上，返回string
	line2, _ := buf2.ReadString(' ')
	fmt.Println("ReadString before ' ':", string(line2))
	fmt.Println("buffer remain:", buf2.String())

	//创建一个空Buffer
	buf3 := bytes.Buffer{}
	buf3.Grow(16) //自动增加缓存容量，保证有n字节剩余空间

	n, _ := buf3.WriteRune(rune('中')) //写入rune编码，返回写入的字节数和错误。
	fmt.Println("buf3 write ", n)
	n, _ = buf3.WriteString("国人")
	fmt.Println("buf3 write ", n)
	fmt.Println(buf3.String())

	buf3.Truncate(6) //将数据长度截断到n字节
	fmt.Println(buf3.String())
}

// BytesTest show the sample code of bytes packagge
func BytesTest() {
	//var b = []byte("Hello world, I'm Golang!!!")
	//var b1 = make([]byte, 2)
	//var b2 = []byte{'a', ' ', ',', 'b'}
	//bytesFunc(b2)
	//bytesFunc(b)
	bytesBuffer()
	// bytesReader()
}
