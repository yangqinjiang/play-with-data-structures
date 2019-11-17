package main

import (
	"fmt"
	"github.com/yangqinjiang/play-with-data-structures/10-Trie/03-Searching-in-Trie"
	"github.com/yangqinjiang/play-with-data-structures/Utils/FileOperation"
	"path/filepath"
	"time"
)

func main() {

	//
	fmt.Println("测试基于BST的集合,统计不同的单词数量")
	readFileAndTotalDifferentWords("10-Trie/03-Searching-in-Trie/pride-and-prejudice.txt",_3_Searching_in_Trie.NewBSTSet())
	fmt.Println("----------------------")
	fmt.Println("测试基于Trie的集合,统计不同的单词数量")
	readFileAndTotalDifferentWords("10-Trie/03-Searching-in-Trie/pride-and-prejudice.txt",_3_Searching_in_Trie.NewTrieSet())

	fmt.Println("测试基于BST的集合,统计不同的单词数量")
	readFileAndTotalDifferentWords("10-Trie/03-Searching-in-Trie/a-tale-of-two-cities.txt",_3_Searching_in_Trie.NewBSTSet())
	fmt.Println("----------------------")
	fmt.Println("测试基于Trie的集合,统计不同的单词数量")
	readFileAndTotalDifferentWords("10-Trie/03-Searching-in-Trie/a-tale-of-two-cities.txt",_3_Searching_in_Trie.NewTrieSet())
}

func readFileAndTotalDifferentWords(path string,set _3_Searching_in_Trie.Set)  {
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
	for _,word := range words1 {
		set.Contains(word)
	}
	//记录结束时间
	elapsed := time.Since(startTime)
	fmt.Printf("Total different words: %d ,took: %f s \n",set.GetSize(),elapsed.Seconds())
}