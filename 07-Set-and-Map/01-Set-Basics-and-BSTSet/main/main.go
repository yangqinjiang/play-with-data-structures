package main

import (
	"fmt"
	_1_Set_Basics_and_BSTSet "github.com/yangqinjiang/play-with-data-structures/07-Set-and-Map/01-Set-Basics-and-BSTSet"
	"github.com/yangqinjiang/play-with-data-structures/Utils/FileOperation"
	"path/filepath"
)
//读取文件的文本内容,添加到BST set集合内,并且统计不同的单词数量
func main() {


	readFileAndTotalDifferentWords("07-Set-and-Map/01-Set-Basics-and-BSTSet/pride-and-prejudice.txt")

	readFileAndTotalDifferentWords("07-Set-and-Map/01-Set-Basics-and-BSTSet/a-tale-of-two-cities.txt")

}

func readFileAndTotalDifferentWords(path string)  {
	fmt.Printf("Read File: %s ....\n",path)

	filename1, _ := filepath.Abs(path)
	words1,err1 := FileOperation.ReadFile(filename1)
	if err1 != nil{
		return
	}
	fmt.Println("Total words: ",len(words1))
	set1 := _1_Set_Basics_and_BSTSet.NewBSTSet()
	for _,word := range words1 {
		set1.Add(word)
	}
	fmt.Printf("Total different words: %d \n",set1.GetSize())
}
