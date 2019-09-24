package main

import (
	"fmt"

	_7_Implement_Queue_in_LinkedList "github.com/yangqinjiang/play-with-data-structures/04-Linked-List/07-Implement-Queue-in-LinkedList"
	"math"
	"math/rand"
	"time"
)
// 测试使用q运行opCount个enqueueu和dequeue操作所需要的时间，单位：秒
func testQueue(q _7_Implement_Queue_in_LinkedList.Queue,opCount int) float64  {
	
	startTime := time.Now()
	//随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<opCount ;i++  {
		q.Enqueue(r.Intn(math.MaxInt32))
	}
	for i:=0;i<opCount ;i++  {
		q.Dequeue()
	}

	elapsed := time.Since(startTime)
	return elapsed.Seconds()
}

func main() {

	opCount := 100000
	//普通数组
	arrayQueue := _7_Implement_Queue_in_LinkedList.NewArrayQueueDefault()
	time1 := testQueue(arrayQueue,opCount)
	fmt.Printf("ArrayQueue, time:  %f s\n",time1)
	//循环队列
	loopQueue := _7_Implement_Queue_in_LinkedList.NewLoopQueueDefault()
	time2 := testQueue(loopQueue,opCount)
	fmt.Printf("LoopQueue, time:  %f s\n",time2)

	//基于链表的队列
	ll_queue := _7_Implement_Queue_in_LinkedList.NewLinkedListQueue()
	time3 := testQueue(ll_queue,opCount)
	fmt.Printf("LinkedListQueue, time:  %f s\n",time3)




}
