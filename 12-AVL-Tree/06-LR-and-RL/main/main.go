package main

import (
	"sort"
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/06-LR-and-RL"
	"github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/06-LR-and-RL/AVLTree"
	"github.com/yangqinjiang/play-with-data-structures/Utils/FileOperation"
	"path/filepath"
	"time"
)


func main() {
	fmt.Println("------------------\nuse BSTMap")
	readFileAndTotalDifferentWordsCount("12-AVL-Tree/06-LR-and-RL/pride-and-prejudice.txt",NewBSTMap())
	fmt.Println("------------------\nuse AVLTree")
	readFileAndTotalDifferentWordsCount2("12-AVL-Tree/06-LR-and-RL/pride-and-prejudice.txt",AVLTree.NewAVLTree())
}
func readFileAndTotalDifferentWordsCount(path string,maper Map)  {
	if maper == nil{
		panic(" maper can't be nil")
	}
	fmt.Printf("Read File: %s ....\n",path)

	filename1, _ := filepath.Abs(path)
	words1,err1 := FileOperation.ReadFile(filename1)
	if err1 != nil{
		return
	}
	fmt.Println("Total words: ",len(words1))

	//将words1排序一下,bst会退化成链表
	sort.Strings(words1)
	fmt.Println("将words排序一下,BST会退化成链表")
	//从set.add开始,记录开始时间
	startTime := time.Now()
	for _,word := range words1 {
		if maper.Contains(word){ //存在,则更新 + 1
			maper.Set(word,maper.Get(word) + 1)
		}else{ 					//不存在,则设置 1
			maper.Add(word,1)
		}

	}
	//再查询一次
	for _,word := range words1 {
		maper.Contains(word)
	}
	//记录结束时间
	elapsed := time.Since(startTime)

	fmt.Printf("Frequency of PRIDE: %d \n" , maper.Get("pride"))
	fmt.Printf("Frequency of PREJUDICE: %d \n" , maper.Get("prejudice"))
	fmt.Printf("Total different words: %d  , took: %f s \n",maper.GetSize(),elapsed.Seconds())
	
}
func readFileAndTotalDifferentWordsCount2(path string,maper *AVLTree.AVLTree)  {
	if maper == nil{
		panic(" maper can't be nil")
	}
	fmt.Printf("Read File: %s ....\n",path)

	filename1, _ := filepath.Abs(path)
	words1,err1 := FileOperation.ReadFile(filename1)
	if err1 != nil{
		return
	}
	fmt.Println("Total words: ",len(words1))
	// 将words1排序一下,avlTree不会退化成链表,保持平稳
	fmt.Println("将words排序一下,AVLTree 不会 退化成链表")
	sort.Strings(words1)
	//从set.add开始,记录开始时间
	startTime := time.Now()
	for _,word := range words1 {
		if maper.Contains(word){ //存在,则更新 + 1
			maper.Set(word,maper.Get(word) + 1)
		}else{ 					//不存在,则设置 1
			maper.Add(word,1)
		}

	}
		//再查询一次
		for _,word := range words1 {
			maper.Contains(word)
		}
	//记录结束时间
	elapsed := time.Since(startTime)

	fmt.Printf("Frequency of PRIDE: %d \n" , maper.Get("pride"))
	fmt.Printf("Frequency of PREJUDICE: %d \n" , maper.Get("prejudice"))
	fmt.Printf("Total different words: %d  , took: %f s \n",maper.GetSize(),elapsed.Seconds())
	fmt.Printf("is BST :  %v  \n",maper.IsBST())
	fmt.Printf("is Balanced  :  %v  \n",maper.IsBalanced())
	
	ok := maper.IsBST() && maper.IsBalanced()
	if (!ok){
		panic("AVL Tree is not BST OR Balanced")
	}
}