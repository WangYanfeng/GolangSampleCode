package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func checkFile(filename string) bool {
	_, err := os.OpenFile(filename, os.O_RDONLY, 044)
	if err != nil {
		return false
	}
	return true
}

func outputReport(result map[string]int) {
	fmt.Println(len(result))
	countMap := make(map[int][]string, len(result))
	for k, v := range result {
		countMap[v] = append(countMap[v], k)
	}

	for k, v := range countMap {
		fmt.Println(k, "-", v)
	}
}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <file1> <file2> ...\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if runtime.GOOS == "windows" {
		filenames := make([]string, 0, len(os.Args[1:]))
		for _, reg := range os.Args[1:] {
			if matches, err := filepath.Glob(reg); err == nil {
				for _, file := range matches {
					if checkFile(file) {
						filenames = append(filenames, file)
					}
				}
			}
		}

		result := analyseFiles(filenames)
		outputReport(result)

		// fmt.Printf("%v", result)
	}

}
