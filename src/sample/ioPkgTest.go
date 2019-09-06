package sample

/**
 * 1. bufio : 带缓存的I/O操作
 * 		bufio Reader
 * 			bufio.NewReader(reader) / bufio.NewReaderSize(reader, size)
 * 			r.Read(b []byte) 读取数据到b中，返回长度
 * 			r.ReadLine(delim byte) / ReadString() / ReadSlice() ReadBytes() / ReadRune() 在r中查找delim，返回delim以及之前的所有数据
 * 			r.Reset() / Discard()
 * 			r.Peek(n) 返回buffer中前n个字节的切片
 * 			r.Buffered() 缓存中数据长度
 * 		bufio Writer
 * 			bufio.NewWriter(writer) / bufio.NewWriterSize(writer, size)
 * 			w.Flush()
 * 			w.Avaliable() 缓存中的可用空间
 * 			w.Buffered()
 * 			w.Write() / w.WriteString()
 * 			w.ReadFrom(reader)
 * 			w.Reset()
 * 		bufio Scanner 连续调用Scan方法将扫描数据中的“指定部分”，跳过各个“指定部分”之间的数据（如：读取多行文件）默认切分函数使用换行符（返回不包含行尾符）
 * 			停止条件：1.遇到io.EOF 2.遇到读写error 3.“指定部分”超出缓存长度
 * 			bufio.NewReader(reader)
 * 			reader.Scan()
 * 			reader.Split()
 * 			reader.Text() / reader.Bytes() / reader.Buffer()
 * 			reader.Split() 自定义分割函数
 * 2. io
 * 3. ioutil
 * 4. strings
 * 5. bytes
 *
 *
 * os.Exit() 推出
 * */

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bufioTest() {
	// 准备从标准输入读取数据
	inputReader := bufio.NewReader(os.Stdin)
	outputWriter := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)

	// bufio.NewReaderSize(os.Stdin, 1024) 设置buffer长度
	fmt.Printf("bufio.NewReader default buffer size: %d.\n", inputReader.Size())
	fmt.Printf("bufio.NewWriter default buffer size: %d.\n", outputWriter.Size())

	fmt.Println("Please input your name:")
	// 读取数据直到碰到 \n 为止。
	input, err := inputReader.ReadString('\n')
	// input, isPrefix, err := inputReader.ReadLine()
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		// 异常退出。
		os.Exit(1)
	} else {
		// 用切片操作删除最后的 \n 。
		name := input[:len(input)-2]
		outputWriter.WriteString("Hello, " + name)
		outputWriter.WriteString("! What can I do for you?\n")
		outputWriter.Flush()
	}

	for scanner.Scan() {
		input = scanner.Text()
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			continue
		}
		// 全部转换为小写。
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			// 正常退出。
			os.Exit(0)
		default:
			fmt.Println("Sorry, I didn't catch you.")
		}
	}
}

// IOTest : test all the io related pkg
func IOTest() {
	bufioTest()
}

// buf := bytes.NewBuffer(make([]byte, 0, 512))
// buf.ReadFrom

// io.LimitReader
// io.Coyp

// bufio

// ioutil.ReadAll()

// string.NewReader()

// bytes.NewBuffer()