package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/08-Heap-and-Priority-Queue/03-Add-and-Sift-Up-in-Heap"
)

func main() {
	fmt.Println("head basics")
	mh := NewMaxHeap()
	mh.Add(1)
	mh.Add(2)
	mh.Add(3)
	mh.Add(4)

	fmt.Println(mh.IsEmpty())
	fmt.Println(mh.Size())
}
