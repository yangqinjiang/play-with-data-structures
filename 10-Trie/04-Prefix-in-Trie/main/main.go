package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/10-Trie/04-Prefix-in-Trie"
	"github.com/yangqinjiang/play-with-data-structures/Utils/FileOperation"
	"path/filepath"
	"time"
)

func main() {

	trie:=NewTrie()
	trie.Add("HELLO中方")
	trie.Add("HI")
	trie.Add("World")
	trie.Add("World") //重复的单词
	if 3 != trie.GetSize(){
		panic("单词个数不一致")
	}
	if false == trie.IsPrefix("HELLO"){
		panic("单词前缀HELLO应该存在的")
	}
	if false == trie.IsPrefix("H"){
		panic("单词前缀H应该存在的")
	}
	if false == trie.IsPrefix("W"){
		panic("单词前缀H应该存在的")
	}
	if true == trie.IsPrefix("NOT_PREFIX"){
		panic("单词前缀HELLO应该存在的")
	}
	//
	fmt.Println("测试基于BST的集合,统计不同的单词数量")
	readFileAndTotalDifferentWords("10-Trie/03-Searching-in-Trie/pride-and-prejudice.txt",NewBSTSet())
	fmt.Println("----------------------")
	fmt.Println("测试基于Trie的集合,统计不同的单词数量")
	readFileAndTotalDifferentWords("10-Trie/03-Searching-in-Trie/pride-and-prejudice.txt",NewTrieSet())

	fmt.Println("测试基于BST的集合,统计不同的单词数量")
	readFileAndTotalDifferentWords("10-Trie/03-Searching-in-Trie/a-tale-of-two-cities.txt",NewBSTSet())
	fmt.Println("----------------------")
	fmt.Println("测试基于Trie的集合,统计不同的单词数量")
	readFileAndTotalDifferentWords("10-Trie/03-Searching-in-Trie/a-tale-of-two-cities.txt",NewTrieSet())
}

func readFileAndTotalDifferentWords(path string,set Set)  {
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