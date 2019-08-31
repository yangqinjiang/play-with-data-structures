package main

import (
	"fmt"
	_7_Dynamic_Array "github.com/yangqinjiang/play-with-data-structures/02-Arrays/07-Dynamic-Array"
)

func main() {
	arr := _7_Dynamic_Array.NewArray(10)
	for i:=0 ;i<10;i++{
		arr.AddLast(i)
	}

	fmt.Println(arr.ToString())


	fmt.Println("AddLast-->")
	for i:=0 ;i<10;i++{
		arr.AddLast(i)
	}
	fmt.Println(arr.ToString())

	fmt.Println("RemoveFirst-->")
	for i:=0 ;i<4;i++{
		arr.RemoveFirst()
	}
	fmt.Println(arr.ToString())
}
