package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/09-Segment-Tree/06-Update-Single-Element-in-Segment-Tree"
)

type MyMerger struct {

}

func (this *MyMerger)Merge(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("segment-tree")

	arr := []int{-2, 0, 3, -5, 2, -1}

	 st := NewSegmentTree(arr,&MyMerger{})

	fmt.Println(st.ToString())
	 // arr 的0至2 三个数两加  -2, 0, 3
	 if 1 != st.Query(0,2){
	 	panic("error")
	 }
	// arr 的2至5 四个数两加
	if -1 != st.Query(2,5){
		panic("error")
	}
	// arr 的0至5  6个数两加
	if -3 != st.Query(0,5){
		panic("error")
	}

	 //update
	 st.Set(0,100)
	// update之后, arr 的0至2 -2, 0, 3 ==> 100, 0, 3  结果应该是103
	if 103 != st.Query(0,2){
		panic("error")
	}
	// update之后, arr 的0至2 100, 0, 3 ==> 100, -99, 3  结果应该是4
	st.Set(1,-99)
	if 4 != st.Query(0,2){
		panic("error")
	}
	fmt.Println("run ok")
}
