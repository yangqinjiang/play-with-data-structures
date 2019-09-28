package main

import (
	"fmt"
	_12_Remove_Elements_in_BST "github.com/yangqinjiang/play-with-data-structures/06-Binary-Search-Tree/12-Remove-Elements-in-BST"
	"math/rand"
	"strconv"
	"time"
)

func shuffle(arr []int)  {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:= len(arr) -1;i>=0;i--{
		pos := r.Int() %len(arr)
		//fmt.Print(pos)
		//fmt.Print(",")
		//交换
		arr[i],arr[pos] = arr[pos],arr[i]
	}
}
func main() {
	bst := _12_Remove_Elements_in_BST.NewBST()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := 10
	order := make([]int,n)
	for i:=0;i<n;i++{
		// 注意, 由于随机生成的数据有重复, 所以bst中的数据数量大概率是小于n的
		bst.Add(r.Intn(n))
		// order数组中存放[0...n)的所有元素
		order[i] = i
	}
	fmt.Println(order)
	// 打乱order数组的顺序
	shuffle(order)
	fmt.Print("打乱order数组的顺序: ")
	fmt.Println(order)
	// 乱序删除[0...n)范围里的所有元素
	for i:=0;i<n;i++{
		if bst.Contains(order[i]){
			bst.Remove(order[i])
			fmt.Println("After remve "+ strconv.Itoa(order[i]) + " , size = " + strconv.Itoa(bst.Size()))
		}
	}
	// 最终整个二分搜索树应该为空
	if !bst.IsEmpty(){
		panic("最终整个二分搜索树应该为空")
	}
}
