package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/09-Segment-Tree/03-Building-Segment-Tree"
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
	fmt.Println("run ok")
}
