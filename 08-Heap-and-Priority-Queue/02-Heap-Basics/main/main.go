package main

import (
	"fmt"
	. "github.com/yangqinjiang/play-with-data-structures/08-Heap-and-Priority-Queue/02-Heap-Basics"
)

func main() {
	fmt.Println("head basics")
	mh := NewMaxHeap()
	fmt.Println(mh.IsEmpty())
	fmt.Println(mh.Size())
}
