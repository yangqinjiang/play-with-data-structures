package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/09-Segment-Tree/04-Query-in-Segment-Tree"
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
	fmt.Println("run ok")
}
