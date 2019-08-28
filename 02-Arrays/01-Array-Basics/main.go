package main

import (
	"fmt"
	"github.com/yangqinjiang/play-with-data-structures/02-Arrays/03-Add-Element-in-Array"
)

//切片(数组的基本使用)
func main() {
	var arr []int;
	for i:=0;i< len(arr);i++  {
		arr[i] = i
	}

	scores := []int{100,99,66}
	for i:=0;i< len(scores);i++  {
		fmt.Println(scores[i])
	}

	for _,score := range scores{
		fmt.Println(score)
	}
	scores[0] = 96
	for i:=0;i< len(scores);i++  {
		fmt.Println(scores[i])
	}



}
