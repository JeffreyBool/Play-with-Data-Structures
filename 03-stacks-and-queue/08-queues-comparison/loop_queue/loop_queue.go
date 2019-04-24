/**
  * Author: JeffreyBool
  * Date: 2019/4/24
  * Time: 21:07
  * Software: GoLand
*/

package loop_queue

import (
	"Play-with-Data-Structures/03-stacks-and-queue/08-queues-comparison/queue"
	"errors"
	"fmt"
)

type LoopQueue struct {
	data  []interface{}
	size  int
	front int
	tail  int
}

//实例化
func NewQueue(capacity ...int) (queue queue.Queue) {
	if len(capacity) > 0 && capacity[0] != 0 {
		queue = &LoopQueue{
			data:  make([]interface{}, capacity[0]+1),
			size:  0,
			front: 0,
			tail:  0,
		}
	} else {
		queue = &LoopQueue{
			data:  make([]interface{}, 11),
			size:  0,
			front: 0,
			tail:  0,
		}
	}
	return
}

//获取队列长度
func (queue *LoopQueue) GetSize() int {
	return queue.size
}

//获取队列容量
func (queue *LoopQueue) GetCapacity() int {
	return cap(queue.data) - 1
}

//判断队列是否为空
func (queue *LoopQueue) IsEmpty() bool {
	return queue.front == queue.tail
}

//向队尾插入元素
func (queue *LoopQueue) Enqueue(v interface{}) error {
	// (tail + 1) % 队列容量 == front 队列满
	if (queue.tail+1)%cap(queue.data) == queue.front {
		queue.resize(queue.GetCapacity() * 2)
	}

	queue.data[queue.tail] = v
	queue.tail = (queue.tail + 1) % cap(queue.data)
	queue.size ++
	return nil
}

//弹出队首元素
func (queue *LoopQueue) Dequeue() (v interface{}, err error) {
	if queue.IsEmpty() {
		err = errors.New("Cannot dequeue from an empty queue.")
		return
	}

	v = queue.data[queue.front]
	queue.data[queue.front] = nil
	queue.front = (queue.front + 1) % cap(queue.data)
	queue.size --
	//缩荣 lazy机制
	if cap(queue.data)/4 == queue.size && cap(queue.data)/2 != 0 {
		queue.resize(queue.GetCapacity() / 2)
	}
	return
}

//查看队首元素
func (queue *LoopQueue) GetFront() (v interface{}, err error) {
	if queue.IsEmpty() {
		err = errors.New("Queue is empty.")
		return
	}
	v = queue.data[queue.front]
	return
}

//格式化输出
func (queue *LoopQueue) PrintIn() (format string) {
	format += fmt.Sprintf("Queue: size = %d , capacity = %d\n",queue.size, cap(queue.data))
	format += "front ["
	for i := 0; i < queue.size; i++{
		format += fmt.Sprintf("%+v",queue.data[(queue.front + i) %cap(queue.data)])
		//注意： 这里说明一下为啥还需要加一， 因为我们数组预留了一个空间，也就是说 当 front + i + 1 == tail 就说明这是数组末尾了。不需要循环了
		if queue.tail != (queue.front + i + 1) % cap(queue.data) {
			format += ", "
		}
	}
	format += "] tail"
	fmt.Println(format)
	return
}

//数组动态扩容机制
func (queue *LoopQueue) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < queue.size; i ++ {
		newData[i] = queue.data[(queue.front + i) % cap(queue.data)]
	}

	queue.data = newData
	queue.front = 0
	queue.tail = queue.size
	newData = nil //手动回收内存，gcc垃圾回收机制调用
}
