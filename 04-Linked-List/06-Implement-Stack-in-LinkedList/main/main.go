package main

import (
	"fmt"

	_6_implement_stack_in_linkedlist "github.com/yangqinjiang/play-with-data-structures/04-Linked-List/06-Implement-Stack-in-LinkedList"
	"math"
	"math/rand"
	"time"
)

// 测试使用stack运行opCount个push和pop操作所需要的时间，单位：秒
func testStack(q _6_implement_stack_in_linkedlist.Stack,opCount int) float64  {

	startTime := time.Now()
	//随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<opCount ;i++  {
		q.Push(r.Intn(math.MaxInt32))
	}
	for i:=0;i<opCount ;i++  {
		q.Pop()
	}

	elapsed := time.Since(startTime)
	return elapsed.Seconds()
}
func main() {

	opCount := 100000
	array_stack := _6_implement_stack_in_linkedlist.NewArrayStackDefault()
	time1 := testStack(array_stack,opCount)
	fmt.Printf("ArrayStack, time:  %f s\n",time1)


	ll_stack := _6_implement_stack_in_linkedlist.NewLinkedListStack()
	time2 := testStack(ll_stack,opCount)
	fmt.Printf("LinkedListStack, time:  %f s\n",time2)

	// 其实这个时间比较很复杂，因为LinkedListStack中包含更多的new操作




}
