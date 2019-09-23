package _6_loop_queue

import (
	"fmt"
	"strings"
)

//使用数组实现的队列,

type LoopQueue struct {
	data []int
	front,tail int
	size int// 有兴趣的同学，在完成这一章后，可以思考一下：
			// LoopQueue中不声明size，如何完成所有的逻辑？
			// 这个问题可能会比大家想象的要难一点点：）
}

func NewLoopQueue(capacity int) *LoopQueue  {
	return &LoopQueue{
		data: []int{capacity+1},//比容量大1
		front:0,
		tail:0,
		size:0,
	}
}
//默认构造函数
func NewLoopQueueDefault() *LoopQueue {
	return NewLoopQueue(10)
}

func (this *LoopQueue)GetCapacity() int  {
	return len(this.data) -1
}

//实现Queue接口
func (this *LoopQueue)GetSize() int {
	return this.size
}

// 下一小节再做具体实现
func (this *LoopQueue)Enqueue(e int) {

}
// 下一小节再做具体实现
func (this *LoopQueue)Dequeue() int {
	return 0
}
// 下一小节再做具体实现
func (this *LoopQueue)getFront(e int) int {
	return 0
}
//打印数组字符串
func (this *LoopQueue)ToString() string  {
	var sb strings.Builder
	fmt.Fprintf(&sb,"Queue: ")
	sb.WriteString("front [")
	//下一小节再做具体实现
	sb.WriteString("] tail")
	return  sb.String()

}