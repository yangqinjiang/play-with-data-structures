package FileOperation

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// 读取文件名称为filename中的内容，并将其中包含的所有词语放进words中
//来源于: https://github.com/Donng/Play-with-Data-Structures/blob/master/Utils/FileOperation/FileOperation.go
func ReadFile(filename string) ([]string,error) {

	// 文件读取
	fin,err0 := os.OpenFile(filename,os.O_RDONLY,0666)
	defer fin.Close()
	if err0 != nil{
		return nil,err0
	}

	buf := bufio.NewReader(fin)
	var words []string
	for{
		line ,err := buf.ReadString('\n')

		if err != nil || io.EOF == err {
			break
		}

		wordSlice := strings.Fields(line) //Fields
		for _, word := range wordSlice {
			if word = extractStr(strings.ToLower(word)); word != "" {
				words = append(words, word)
			}
		}
	}
	return words,nil
}
func extractStr(str string) string {
	var res []rune
	for _, letter := range str {
		if letter >= 'a' && letter <= 'z' {
			res = append(res, letter)
		}
	}
	return string(res)
}