package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/02-Calculating-Balance-Factor"
	"github.com/yangqinjiang/play-with-data-structures/12-AVL-Tree/02-Calculating-Balance-Factor/AVLTree"
	"github.com/yangqinjiang/play-with-data-structures/Utils/FileOperation"
	"path/filepath"
	"time"
)


func main() {

	fmt.Println("use BSTMap")
	readFileAndTotalDifferentWordsCount("12-AVL-Tree/02-Calculating-Balance-Factor/pride-and-prejudice.txt",NewBSTMap())
	fmt.Println("use AVLTree")
	readFileAndTotalDifferentWordsCount("12-AVL-Tree/02-Calculating-Balance-Factor/pride-and-prejudice.txt",AVLTree.NewAVLTree())
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
	//从set.add开始,记录开始时间
	startTime := time.Now()
	for _,word := range words1 {
		if maper.Contains(word){ //存在,则更新 + 1
			maper.Set(word,maper.Get(word) + 1)
		}else{ 					//不存在,则设置 1
			maper.Add(word,1)
		}

	}
	//记录结束时间
	elapsed := time.Since(startTime)

	fmt.Printf("Frequency of PRIDE: %d \n" , maper.Get("pride"))
	fmt.Printf("Frequency of PREJUDICE: %d \n" , maper.Get("prejudice"))
	fmt.Printf("Total different words: %d  , took: %f s \n",maper.GetSize(),elapsed.Seconds())
}