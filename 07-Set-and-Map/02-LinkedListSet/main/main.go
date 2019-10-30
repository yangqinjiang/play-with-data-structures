package main

import (
	"fmt"
	_2_Linked_List_Set "github.com/yangqinjiang/play-with-data-structures/07-Set-and-Map/02-LinkedListSet"

	"github.com/yangqinjiang/play-with-data-structures/Utils/FileOperation"
	"path/filepath"
)
//读取文件的文本内容,添加到BST/LinkedList set集合内,并且统计不同的单词数量
func main() {

	//
	fmt.Println("测试基于BST的集合,统计不同的单词数量")
	readFileAndTotalDifferentWords("07-Set-and-Map/02-LinkedListSet/pride-and-prejudice.txt",_2_Linked_List_Set.NewBSTSet())
	fmt.Println("----------------------")
	readFileAndTotalDifferentWords("07-Set-and-Map/02-LinkedListSet/a-tale-of-two-cities.txt",_2_Linked_List_Set.NewBSTSet())
	fmt.Println("----------------------")
	fmt.Println("测试基于LinkedList的集合,统计不同的单词数量")
	readFileAndTotalDifferentWords("07-Set-and-Map/02-LinkedListSet/pride-and-prejudice.txt",_2_Linked_List_Set.NewLinkedListSet())
	fmt.Println("----------------------")
	readFileAndTotalDifferentWords("07-Set-and-Map/02-LinkedListSet/a-tale-of-two-cities.txt",_2_Linked_List_Set.NewLinkedListSet())

}

func readFileAndTotalDifferentWords(path string,set _2_Linked_List_Set.Set)  {
	if set == nil{
		panic(" set can't be nil")
	}
	fmt.Printf("Read File: %s ....\n",path)

	filename1, _ := filepath.Abs(path)
	words1,err1 := FileOperation.ReadFile(filename1)
	if err1 != nil{
		return
	}
	fmt.Println("Total words: ",len(words1))
	//set1 := _2_Linked_List_Set.NewBSTSet()
	for _,word := range words1 {
		set.Add(word)
	}
	fmt.Printf("Total different words: %d \n",set.GetSize())
}
