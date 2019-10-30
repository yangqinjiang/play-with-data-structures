package main

import (
	"fmt"
	_2_Linked_List_Set "github.com/yangqinjiang/play-with-data-structures/07-Set-and-Map/02-LinkedListSet"
	"time"

	"github.com/yangqinjiang/play-with-data-structures/Utils/FileOperation"
	"path/filepath"
)
//读取文件的文本内容,添加到BST/LinkedList set集合内,并且统计不同的单词数量,运行时间
func main() {
	//复杂度
	//BST( O(h)  )在某些情况下可退化成链表, 例如输入的数据是有序的
	//LinkedList (  O(h)  )
	//
	fmt.Println("测试基于BST的集合,统计不同的单词数量,运行时间")
	readFileAndTotalDifferentWordsAndTime("07-Set-and-Map/03-Time-Complexity-of-Set/pride-and-prejudice.txt",_2_Linked_List_Set.NewBSTSet())
	fmt.Println("----------------------")
	readFileAndTotalDifferentWordsAndTime("07-Set-and-Map/03-Time-Complexity-of-Set/a-tale-of-two-cities.txt",_2_Linked_List_Set.NewBSTSet())
	fmt.Println("----------------------")
	fmt.Println("测试基于LinkedList的集合,统计不同的单词数量,运行时间")
	readFileAndTotalDifferentWordsAndTime("07-Set-and-Map/03-Time-Complexity-of-Set/pride-and-prejudice.txt",_2_Linked_List_Set.NewLinkedListSet())
	fmt.Println("----------------------")
	readFileAndTotalDifferentWordsAndTime("07-Set-and-Map/03-Time-Complexity-of-Set/a-tale-of-two-cities.txt",_2_Linked_List_Set.NewLinkedListSet())

}

func readFileAndTotalDifferentWordsAndTime(path string,set _2_Linked_List_Set.Set)  {
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
	//从set.add开始,记录开始时间
	startTime := time.Now()
	for _,word := range words1 {
		set.Add(word)
	}
	//记录结束时间
	elapsed := time.Since(startTime)
	fmt.Printf("Total different words: %d  , took: %f s \n",set.GetSize(),elapsed.Seconds())
}