package main

import (
	"fmt"
	"github.com/yangqinjiang/play-with-data-structures/13-Red-Black-Tree/05-Keep-Root-Black-and-Left-Rotation/AVLTree"
	. "github.com/yangqinjiang/play-with-data-structures/13-Red-Black-Tree/05-Keep-Root-Black-and-Left-Rotation/BST"
	"github.com/yangqinjiang/play-with-data-structures/Utils/FileOperation"
	"path/filepath"
	"sort"
	"time"
)


func main() {
	filename := "13-Red-Black-Tree/05-Keep-Root-Black-and-Left-Rotation/pride-and-prejudice.txt"

	fmt.Printf("Read File: %s ....\n",filename)

	filename1, _ := filepath.Abs(filename)
	words1,err1 := FileOperation.ReadFile(filename1)
	if err1 != nil{
		return
	}
	fmt.Println("Total words: ",len(words1))
	if !true{
		sort.Strings(words1)
		fmt.Println("将words排序一下,AVLTree比BST快")
	}


	//从set.add开始,记录开始时间
	//------------------
	startTime := time.Now()
	bst :=  NewBST()
	for _,word := range words1 {
		if bst.Contains(word){ //存在,则更新 + 1
			bst.Set(word,bst.Get(word) + 1)
		}else{ 					//不存在,则设置 1
			bst.Add(word,1)
		}

	}
	//再查询一次
	for _,word := range words1 {
		bst.Contains(word)
	}

	//记录结束时间
	elapsed := time.Since(startTime)
	fmt.Printf("BST: Frequency of PRIDE: %d , took: %f s \n" , bst.Get("pride"),elapsed.Seconds())

	//------------------
	startTime = time.Now()
	avl :=  AVLTree.NewAVLTree()
	for _,word := range words1 {
		if avl.Contains(word){ //存在,则更新 + 1
			avl.Set(word,avl.Get(word) + 1)
		}else{ 					//不存在,则设置 1
			avl.Add(word,1)
		}

	}
	//再查询一次
	for _,word := range words1 {
		avl.Contains(word)
	}

	//记录结束时间
	elapsed = time.Since(startTime)
	fmt.Printf("AVL: Frequency of PRIDE: %d , took: %f s \n" , avl.Get("pride"),elapsed.Seconds())
}

