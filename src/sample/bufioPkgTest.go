package sample

/**
 * 带缓存的I/O操作
 * 1. NewReader() / NewReaderSize()
 * 2. Read() / ReadLine() / ReadString() / ReadSlice() ReadBytes()
 * 3. Reset() / Discard()
 * 4. Peek()
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

// BufIOTest : bufio read from stdin
func BufIOTest() {
	// 准备从标准输入读取数据。
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println(inputReader.Size())
	fmt.Println("Please input your name:")
	// 读取数据直到碰到 \n 为止。
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		// 异常退出。
		os.Exit(1)
	} else {
		// 用切片操作删除最后的 \n 。
		name := input[:len(input)-2]
		fmt.Printf("Hello, %s! What can I do for you?\n", name)
	}

	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			continue
		}
		input = input[:len(input)-2]
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

// buf := bytes.NewBuffer(make([]byte, 0, 512))
// buf.ReadFrom

// io.LimitReader
// io.Coyp

// bufio

// ioutil.ReadAll()

// string.NewReader()

// bytes.NewBuffer()
