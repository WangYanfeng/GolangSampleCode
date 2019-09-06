package sample

/**
 * os package
 * 1. 文件是否存在
 * 		判断os.State() 返回的error：os.IsNotExist(err)
 *
 * 2. 检查文件权限
 * 		判断os.Open(...) 返回的error：os.IsPermission(err)
 *
 * 3. 文件信息 FileInfo
 * 		os.State()
 * 		.Name() / .Size() / .Mode() / .IsDir() / .Sys() / .ModTime()
 *
 * 4. 文件读写
 *		os.Open() / os.OpenFile()
 *		os.Create()
 *		os.Truncate()
 *
 * 5. 文件描述符 os.File
 * 		fp.Read() / os.ReadAtList() / ReadFull() / fp.Write() 
 * 		fp.Seek()
 * 		fp.Sync() / fp.Close()
 *
 * 6. 操作
 * 		os.Rename() / os.Remove()
 * 		os.Copy()
 * 		os.TempDir()
 * 		os.Getwd()
 *
 * 7. 修改文件信息
 * 		os.Chown()
 * 		os.Chtimes()
 * 		os.Link() / os.Symlink() / os.Lstat() / os.Lchown()
 *
 * ioutil package
 * 1. 快写文件
 * 		ioutil.WriteFile()
 * 2. 快读文件
 * 		ioutil.ReadFile() / ioutil.ReadAll(fp)
 * 3. 临时文件
 * 		ioutil.TempDir() / ioutil.TempFile()
 *
 * archive/zip package
 * 		zip.OpenReader() / zip.NewWriter()
 */

import (
	"fmt"
	"os"
)

const fileName = "\\static\\config.ini"

func readText(file string) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Fprintf(os.Stderr, "file %s is not permission.\n", file)
			return
		} else if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "file %s is not exist.\n", file)
			return
		}
	}
	fmt.Printf("%s size: %dB\n", fileInfo.Name(), fileInfo.Size())

	fp, err := os.Open(file)
	// os.OpenFile(file, os.O_RDONLY,0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s open error: %s.\n", fileInfo.Name(), err.Error())
		return
	}
	defer fp.Close()

	// buffer读取
	// 切片读取
}

// FileTest : test file
func FileTest() {
	str, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	readText(str + fileName)
}
