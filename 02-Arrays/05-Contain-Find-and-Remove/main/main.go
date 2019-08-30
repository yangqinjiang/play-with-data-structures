package main

import (
	"fmt"
	_5_Contain_Find_and_Remove "github.com/yangqinjiang/play-with-data-structures/02-Arrays/05-Contain-Find-and-Remove"
)

func main() {
	arr := _5_Contain_Find_and_Remove.NewArray(20)
	for i:=0 ;i<10;i++{
		arr.AddLast(i)
	}

	fmt.Println(arr.ToString())

	arr.Add(1,100)
	fmt.Println(arr.ToString())

	arr.AddFirst(-1)
	fmt.Println(arr.ToString())

	arr.Remove(2)
	fmt.Println(arr.ToString())

	arr.RemoveElement(4)
	fmt.Println(arr.ToString())

	arr.RemoveFirst()
	fmt.Println(arr.ToString())
}
