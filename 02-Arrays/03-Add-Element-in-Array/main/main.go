package main

import (
	"fmt"
	_3_Add_Element_in_Array "github.com/yangqinjiang/play-with-data-structures/02-Arrays/03-Add-Element-in-Array"
)

func main() {
	//fmt.Println(_3_Add_Element_in_Array.NewArray(100))
	arr := _3_Add_Element_in_Array.NewArrayDefault()
	fmt.Println(arr)
	arr.AddFirst(1)
	arr.AddFirst(1)
	arr.Add(2,5)
	arr.AddLast(8)
	arr.Add(4,5)
	arr.AddLast(8)
	fmt.Println(arr.GetSize())
	fmt.Println(arr)
}
