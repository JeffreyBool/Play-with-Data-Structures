/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 19:29
  * Software: GoLand
*/

package linkedlist_queue_test

import (
	"testing"
	"Play-with-Data-Structures/04-linked-list/07-implement-queue-in-linkedlist/linkedlist_queue"
	"Play-with-Data-Structures/04-linked-list/07-implement-queue-in-linkedlist/queue"
)

func TestNewQueue(t *testing.T) {
	queue := linkedlist_queue.NewQueue()
	queue.PrintIn()
}

//获取队列长度
func TestLoopQueue_GetSize(t *testing.T) {
	var queue queue.Queue
	queue = linkedlist_queue.NewQueue()
	t.Logf("queue len: %d \n",queue.GetSize())
}

//判断队列是否为空
func TestLoopQueue_IsEmpty(t *testing.T) {
	var queue queue.Queue
	queue = linkedlist_queue.NewQueue()
	t.Logf("queue empty: %t \n",queue.IsEmpty())
}

//向队尾添加元素
func TestLoopQueue_Enqueue(t *testing.T) {
	var queue queue.Queue
	queue = linkedlist_queue.NewQueue()
	queue.Enqueue("A")
	queue.Enqueue("B")
	queue.Enqueue("C")
	queue.Enqueue("D")
	queue.Enqueue("E")
	queue.PrintIn()
	queue.Dequeue()
	queue.PrintIn()
}

//弹出队首元素
func TestLoopQueue_Dequeue(t *testing.T) {
	var queue queue.Queue
	queue = linkedlist_queue.NewQueue()
	queue.Enqueue("A")
	queue.Enqueue("B")
	queue.Enqueue("C")
	queue.Enqueue("D")
	queue.Enqueue("E")
	queue.PrintIn()
	queue.Dequeue()
	queue.Enqueue("D")
	queue.Enqueue("D")
	queue.Enqueue("D")
	queue.Enqueue("D")
	queue.Enqueue("D")
	queue.Enqueue("D")
	queue.PrintIn()
	queue.Dequeue()
	queue.Enqueue("A")
	queue.PrintIn()
	queue.Dequeue()
	queue.Enqueue("B")
	queue.PrintIn()
}

//查看队首元素
func TestLoopQueue_GetFront(t *testing.T) {
	var queue queue.Queue
	queue = linkedlist_queue.NewQueue()
	queue.Enqueue("A")
	queue.Enqueue("B")
	queue.Enqueue("C")
	queue.Enqueue("D")
	queue.Enqueue("E")
	queue.PrintIn()
	if front, err := queue.GetFront(); err != nil{
		t.Error(err)
	}else{
		t.Logf("queue front: %+v\n",front)
	}
}

//格式化输出
func TestLoopQueue_PrintIn(t *testing.T) {
	var queue queue.Queue
	queue = linkedlist_queue.NewQueue()
	queue.Enqueue("A")
	queue.Enqueue("B")
	queue.Enqueue("C")
	queue.Enqueue("D")
	queue.Enqueue("E")
	queue.PrintIn()
}
