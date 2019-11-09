package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/08-Heap-and-Priority-Queue/04-Extract-and-Sift-Down-in-Heap"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("head basics")
	mh := NewMaxHeap()
	n := 100000
	//随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<n ;i++  {
		mh.Add(r.Intn(1<<31 -1))
	}
	fmt.Println("after add ,size=",mh.Size())
	arr := []int{}
	for i:=0;i<n ;i++  {
		arr = append(arr, mh.ExtractMax())
	}
	fmt.Println("after ExtractMax ,size=",mh.Size())

	//测试ExtractMax导出的顺序
	//从第二个元素开始
	for i:=1;i<n ;i++  {
		if arr[i-1] < arr[i]{
			panic("sort Error")
		}
	}
	fmt.Println(len(arr))
}
