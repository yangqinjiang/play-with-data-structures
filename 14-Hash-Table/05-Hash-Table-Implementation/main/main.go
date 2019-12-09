package main

import (
	"fmt"
	"github.com/yangqinjiang/play-with-data-structures/14-Hash-Table/05-Hash-Table-Implementation/HashTable"
	"github.com/yangqinjiang/play-with-data-structures/14-Hash-Table/05-Hash-Table-Implementation/AVLTree"
	. "github.com/yangqinjiang/play-with-data-structures/14-Hash-Table/05-Hash-Table-Implementation/BST"
	"github.com/yangqinjiang/play-with-data-structures/14-Hash-Table/05-Hash-Table-Implementation/RBTree"
	"github.com/yangqinjiang/play-with-data-structures/Utils/FileOperation"

	"path/filepath"
	"sort"
	"time"
)

func main()  {
	main1()
}

//测试文本
func main1() {
	fmt.Println("main1")
	filename := "14-Hash-Table/05-Hash-Table-Implementation/pride-and-prejudice.txt"

	fmt.Printf("Read File: %s ....\n",filename)

	filename1, _ := filepath.Abs(filename)
	words1,err1 := FileOperation.ReadFile(filename1)
	if err1 != nil{
		return
	}
	fmt.Println("Total words: ",len(words1))
	if true{
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
		if !bst.Contains(word){
			panic("bst should be contain "+word+" ,error")
		}
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
		if !avl.Contains(word){
			panic("avl should be contain "+word+" ,error")
		}
	}

	//记录结束时间
	elapsed = time.Since(startTime)
	fmt.Printf("AVL: Frequency of PRIDE: %d , took: %f s \n" , avl.Get("pride"),elapsed.Seconds())
	//------------------
	startTime = time.Now()
	rbTree :=  RBTree.NewRBTree()
	for _,word := range words1 {
		if rbTree.Contains(word){ //存在,则更新 + 1
			rbTree.Set(word,rbTree.Get(word) + 1)
		}else{ 					//不存在,则设置 1
			rbTree.Add(word,1)
		}

	}
	//再查询一次
	for _,word := range words1 {
		if !rbTree.Contains(word){
			panic("rbTree should be contain "+word+" ,error")
		}
	}

	//记录结束时间
	elapsed = time.Since(startTime)
	fmt.Printf("RBTree: Frequency of PRIDE: %d , took: %f s \n" , rbTree.Get("pride"),elapsed.Seconds())

	//------------------
	startTime = time.Now()
	hashtable :=  HashTable.NewHashTable2()
	for _,word := range words1 {
		if hashtable.Contains(word){ //存在,则更新 + 1
			hashtable.Set(word,hashtable.Get(word) + 1)
		}else{ 					//不存在,则设置 1
			hashtable.Add(word,1)
		}

	}
	//再查询一次
	for _,word := range words1 {
		if !hashtable.Contains(word){
			panic("hashTable should be contain "+word+" ,error")
		}
	}

	//记录结束时间
	elapsed = time.Since(startTime)
	fmt.Printf("HashTable: Frequency of PRIDE: %d , took: %f s \n" , hashtable.Get("pride"),elapsed.Seconds())
}


