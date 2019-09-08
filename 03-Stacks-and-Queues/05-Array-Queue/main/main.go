package main

import (
	"fmt"
	_5_Array_Queue "github.com/yangqinjiang/play-with-data-structures/03-Stacks-and-Queues/05-Array-Queue"
	"strings"
)

func main() {
	queue := _5_Array_Queue.NewArrayQueue(20)
	for i:=0; i<10;i++  {
		queue.Enqueue(i)
		fmt.Println(queue.ToString())
		if i%3 ==2{
			queue.Dequeue()
			fmt.Println(queue.ToString())
		}
	}

	//
	if len(_5_Array_Queue.LevelOrder(nil)) != 0{
		panic("should be fail")
	}
	/**
	    3
	   / \
	  9  20
	    /  \
	   15   7
	 */
	_20 := &_5_Array_Queue.TreeNode{Val:20,
		Left:&_5_Array_Queue.TreeNode{Val:15},
		Right:&_5_Array_Queue.TreeNode{Val:7}}
	root := &_5_Array_Queue.TreeNode{Val:3,Left:&_5_Array_Queue.TreeNode{Val:9},Right:_20}
	var sb strings.Builder
	fmt.Fprintf(&sb,"%d",_5_Array_Queue.LevelOrder(root))
	//输出数组
	if "[[3] [9 20] [15 7]]" != sb.String(){
		panic("should success")
	}

}
