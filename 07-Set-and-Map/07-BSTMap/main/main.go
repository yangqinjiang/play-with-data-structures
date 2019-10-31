package main

import (
	"fmt"
	_6_LinkedListMap "github.com/yangqinjiang/play-with-data-structures/07-Set-and-Map/06-LinkedListMap"

	_7_BSTMap "github.com/yangqinjiang/play-with-data-structures/07-Set-and-Map/07-BSTMap"
	"github.com/yangqinjiang/play-with-data-structures/07-Set-and-Map/07-BSTMap/BSTMap"
	"github.com/yangqinjiang/play-with-data-structures/Utils/FileOperation"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("use linkedListMap")
	readFileAndTotalDifferentWordsCount("07-Set-and-Map/07-BSTMap/pride-and-prejudice.txt",_7_BSTMap.NewLinkedListMap())
	fmt.Println("use BSTMap")
	readFileAndTotalDifferentWordsCount("07-Set-and-Map/07-BSTMap/pride-and-prejudice.txt",BSTMap.NewBSTMap())
}

func readFileAndTotalDifferentWordsCount(path string,maper _6_LinkedListMap.Map)  {
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