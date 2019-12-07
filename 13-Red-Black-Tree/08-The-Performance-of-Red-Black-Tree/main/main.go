package main

import (
	"fmt"
	"github.com/yangqinjiang/play-with-data-structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/AVLTree"
	. "github.com/yangqinjiang/play-with-data-structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/BST"
	"github.com/yangqinjiang/play-with-data-structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/RBTree"
	"github.com/yangqinjiang/play-with-data-structures/Utils/FileOperation"
	"math"
	"math/rand"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

func main()  {
	//main1()
	//main2()
	main3()
}

//测试文本
func main1() {
	fmt.Println("main1")
	filename := "13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/pride-and-prejudice.txt"

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
		rbTree.Contains(word)
	}

	//记录结束时间
	elapsed = time.Since(startTime)
	fmt.Printf("RBTree: Frequency of PRIDE: %d , took: %f s \n" , rbTree.Get("pride"),elapsed.Seconds())

}

//测试随机数,在教程视频的耗时: RBTree < BST < AVLTree
func main2()  {
	fmt.Println("main2")
	n := 20000000
	//随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	testData := make([]string,n)
	for i:=0;i<n ;i++  {
		testData[i] = string(r.Intn(math.MaxInt32))
	}
	//Test BST
	startTime := time.Now()
	bst :=  NewBST()
	for _,x := range testData {
		bst.Add(x,0)
	}
	//记录结束时间
	elapsed := time.Since(startTime)
	fmt.Printf("BST:  took: %f s \n" ,  elapsed.Seconds())

	//Test AVL
	startTime  = time.Now()
	avl := AVLTree.NewAVLTree()
	for _,x := range testData {
		avl.Add(x,0)
	}
	//记录结束时间
	elapsed = time.Since(startTime)
	fmt.Printf("AVL:  took: %f s \n" ,  elapsed.Seconds())

	//Test AVL
	startTime  = time.Now()
	rbt := RBTree.NewRBTree()
	for _,x := range testData {
		rbt.Add(x,0)
	}
	//记录结束时间
	elapsed = time.Since(startTime)
	fmt.Printf("RBTree:  took: %f s \n" ,  elapsed.Seconds())
}

//测试顺序数,在教程视频的耗时: RBTree < AVLTree, 实际是 AVLTree < RBTree
func main3()  {
	fmt.Println("main3")
	n := 20000000
	testData := make([]string,n)
	for i:=0;i<n ;i++  {
		testData[i] = strconv.Itoa(i)
	}
	//不测试BST, 因为它退化成了链表

	//Test AVL
	startTime  := time.Now()
	avl := AVLTree.NewAVLTree()
	for _,x := range testData {
		avl.Add(x,0)
	}
	//记录结束时间
	elapsed := time.Since(startTime)
	fmt.Printf("AVL:  took: %f s \n" ,  elapsed.Seconds())

	//Test AVL
	startTime  = time.Now()
	rbt := RBTree.NewRBTree()
	for _,x := range testData {
		rbt.Add(x,0)
	}
	//记录结束时间
	elapsed = time.Since(startTime)
	fmt.Printf("RBTree:  took: %f s \n" ,  elapsed.Seconds())
}

