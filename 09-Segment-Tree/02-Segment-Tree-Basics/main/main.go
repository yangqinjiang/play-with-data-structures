package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/09-Segment-Tree/02-Segment-Tree-Basics"
	"strconv"
)

func main() {
	fmt.Println("segment-tree")
	n := 1000
	var arr []int;

	for i:=0;i< n;i++  {
		arr = append(arr,i)
	}
	 st := NewSegmentTree(arr)
	 if n != st.GetSize(){
	 	panic("size neq " + strconv.Itoa(n))
	 }
	 first := st.Get(0)
	for i:=0;i< n;i++  {
		if first > st.Get(i){
			panic("error data")
		}
	}
	fmt.Println("run ok")
}
