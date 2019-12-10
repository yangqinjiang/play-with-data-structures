package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/08-Heap-and-Priority-Queue/05-Heapify-and-Replace-in-Heap"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("05-Heapify-and-Replace-in-Heap")
	n := 100000
	testData1 := make([]int,n)
	//随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<n ;i++  {
		testData1[i] = r.Intn(1<<31 -1)
	}
	testData2 := make([]int,n)
	copy(testData1, testData2)

	took := testHeap(testData1,false)
	fmt.Printf(" not isHeapify ,took: %f s \n",took)

	took = testHeap(testData2,true)
	fmt.Printf(" IsHeapify ,took: %f s \n",took)
}

func testHeap(testData []int,isHeapify bool)  float64{
	startTime := time.Now()
	 var mh *MaxHeap
	 if isHeapify{
	 	//全部添加,再siftDown
	 	mh = NewMaxHeapWithArray(testData)
	 }else{
		mh = NewMaxHeapByCapacity(len(testData))
	 	//一个一个添加
	 	for _,num := range testData{
			mh.Add(num)
		}
	 }
	arr := []int{}
	for i:=0;i<len(testData) ;i++  {
		arr = append(arr, mh.ExtractMax())
	}

	//测试ExtractMax导出的顺序
	//从第二个元素开始
	for i:=1;i<len(testData) ;i++  {
		if arr[i-1] < arr[i]{
			panic("sort Error")
		}
	}
	//记录结束时间
	elapsed := time.Since(startTime)

	return elapsed.Seconds()
}