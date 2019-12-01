package main

import (
	"sort"
	"fmt"
	
	AVLTreeMap "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/Map/AVLTree"
	LinkedListMap "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/Map/LinkedList"
	BSTMap "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/Map/BST"
	Map "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/Map"

	AVLTreeSet "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/Set/AVLTree"
	LinkedListSet "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/Set/LinkedList"
	BSTSet "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/Set/BST"
	Set "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/Set"
	"github.com/yangqinjiang/play-with-data-structures/Utils/FileOperation"
	"path/filepath"
	"time"
)


func main() {
	filename := "12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/pride-and-prejudice.txt"
	fmt.Println("-------------Map------------------")
	//Map
	fmt.Println("use linkedListMap")
	testMap(filename,LinkedListMap.NewLinkedListMap())
	fmt.Println("use BSTMap")
	testMap(filename,BSTMap.NewBSTMap())
	fmt.Println("use AVLMap")
	testMap(filename,AVLTreeMap.NewAVLMap())

	fmt.Println("-------------Set------------------")

	//Set
	fmt.Println("use linkedListSet")
	testSet(filename,LinkedListSet.NewLinkedListSet())
	fmt.Println("use BSTSet")
	testSet(filename,BSTSet.NewBSTSet())
	fmt.Println("use AVLSet")
	testSet(filename,AVLTreeSet.NewAVLSet())

	fmt.Println("-------------OK-----------------")

}


func testMap(path string,maper Map.Map)  {
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
	sort.Strings(words1)
	fmt.Println("将words排序一下,AVLTree比BST快")
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


func testSet(path string,set Set.Set)  {
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
	sort.Strings(words1)
	fmt.Println("将words排序一下,AVLTree比BST快")
	//从set.add开始,记录开始时间
	startTime := time.Now()
	for _,word := range words1 {
		set.Add(word)
	}
	//再查询一次
	for _,word := range words1 {
		set.Contains(word)
	}
	//记录结束时间
	elapsed := time.Since(startTime)
	fmt.Printf("Total different words: %d  , took: %f s \n",set.GetSize(),elapsed.Seconds())
}