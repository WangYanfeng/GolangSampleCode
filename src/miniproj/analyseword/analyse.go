package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func calculateWord(line string, result map[string]int) {
	line = strings.TrimSpace(line)

	noALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	words := strings.FieldsFunc(line, noALetter)

	for _, w := range words {
		w = strings.ToLower(w)
		if _, ok := result[w]; ok {
			result[w]++
		} else {
			result[w] = 1
		}
	}
}

func analyseFiles(filenames []string) map[string]int {
	result := make(map[string]int)
	for _, file := range filenames {
		fp, _ := os.Open(file)
		reader := bufio.NewReader(fp)
		for {
			line, err := reader.ReadString('\n')
			calculateWord(line, result)

			if err != nil {
				if err != io.EOF {
					log.Println("failed to finish reading the file: ", err)
				}
				break
			}
		}
	}
	return result
}
