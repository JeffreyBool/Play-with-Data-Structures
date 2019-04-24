/**
  * Author: JeffreyBool
  * Date: 2019/4/24
  * Time: 20:47
  * Software: GoLand
*/

package loop_queue_test

import (
	"testing"
	"Play-with-Data-Structures/03-stacks-and-queue/05-array-queue"
	"Play-with-Data-Structures/03-stacks-and-queue/07-Implementation-of-Loop-queue/loop_queue"
)

//初始化
func TestNewQueue(t *testing.T) {
	var arrayQueue queue.Queue
	arrayQueue = loop_queue.NewQueue()
	arrayQueue.PrintIn()
}

//获取队列长度
func TestArrayQueue_GetSize(t *testing.T) {
	var arrayQueue queue.Queue
	arrayQueue = loop_queue.NewQueue()
	t.Logf("queue len: %d \n",arrayQueue.GetSize())
}

//判断队列是否为空
func TestArrayQueue_IsEmpty(t *testing.T) {
	var arrayQueue queue.Queue
	arrayQueue = loop_queue.NewQueue()
	t.Logf("queue empty: %t \n",arrayQueue.IsEmpty())
}

//向队尾添加元素
func TestArrayQueue_Enqueue(t *testing.T) {
	var arrayQueue queue.Queue
	arrayQueue = loop_queue.NewQueue()
	arrayQueue.Enqueue("A")
	arrayQueue.Enqueue("B")
	arrayQueue.Enqueue("C")
	arrayQueue.Enqueue("D")
	arrayQueue.Enqueue("E")
	arrayQueue.PrintIn()
	arrayQueue.Dequeue()
	arrayQueue.PrintIn()
}

//弹出队首元素
func TestArrayQueue_Dequeue(t *testing.T) {
	var arrayQueue queue.Queue
	arrayQueue = loop_queue.NewQueue()
	arrayQueue.Enqueue("A")
	arrayQueue.Enqueue("B")
	arrayQueue.Enqueue("C")
	arrayQueue.Enqueue("D")
	arrayQueue.Enqueue("E")
	arrayQueue.PrintIn()
	arrayQueue.Dequeue()
	arrayQueue.Enqueue("D")
	arrayQueue.Enqueue("D")
	arrayQueue.Enqueue("D")
	arrayQueue.Enqueue("D")
	arrayQueue.Enqueue("D")
	arrayQueue.Enqueue("D")
	arrayQueue.Dequeue()
	arrayQueue.Enqueue("A")
	arrayQueue.PrintIn()
}

//查看队首元素
func TestArrayQueue_GetFront(t *testing.T) {
	var arrayQueue queue.Queue
	arrayQueue = loop_queue.NewQueue()
	arrayQueue.Enqueue("A")
	arrayQueue.Enqueue("B")
	arrayQueue.Enqueue("C")
	arrayQueue.Enqueue("D")
	arrayQueue.Enqueue("E")
	arrayQueue.PrintIn()
	if front, err := arrayQueue.GetFront(); err != nil{
		t.Error(err)
	}else{
		t.Logf("queue front: %+v\n",front)
	}
}
