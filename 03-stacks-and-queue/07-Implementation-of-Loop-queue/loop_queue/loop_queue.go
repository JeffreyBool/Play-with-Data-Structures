/**
  * Author: JeffreyBool
  * Date: 2019/4/24
  * Time: 17:16
  * Software: GoLand
*/

package loop_queue

import (
	"Play-with-Data-Structures/03-stacks-and-queue/05-array-queue"
	"errors"
	"fmt"
)

// 有兴趣的同学，在完成这一章后，可以思考一下：
// LoopQueue中不声明size，如何完成所有的逻辑？
// 这个问题可能会比大家想象的要难一点点：）
type LoopQueue struct {
	data              []interface{}
	size, front, tail int
}

const size  = 10

//初始化循环队列
// 队列存在以空位，所以参数为20就需要开辟21的位置(因为会浪费一个空间用来做循环队列计算)
func NewQueue(capacity ...int) (queue queue.Queue) {
	if len(capacity) > 0 && capacity[0] != 0 {
		queue = &LoopQueue{
			data:make([]interface{},capacity[0] + 1),
			size:0,front:0,tail:0,
		}
	}else{
		queue = &LoopQueue{
			data:make([]interface{},size + 1),
			size:0,front:0,tail:0,
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
	// 与上对应，需要删除空位
	return len(queue.data) -1
}

//判断队列是否为空
func (queue *LoopQueue) IsEmpty() bool {
	return queue.front == queue.tail
}

//向队尾插入元素
func (queue *LoopQueue) Enqueue(v interface{}) (err error) {
	if queue.front == (queue.tail + 1) % len(queue.data) {
		queue.resize(queue.GetCapacity() * 2)
	}
	queue.data[queue.tail] = v
	queue.tail = (queue.tail + 1) % len(queue.data)
	queue.size ++
	return
}

//弹出队首元素
func (queue *LoopQueue) Dequeue() (v interface{},err error) {
	if queue.IsEmpty() {
		err = errors.New("Cannot dequeue from an empty queue.")
		return
	}
	v = queue.data[queue.front]
	queue.front = (queue.front + 1) % len(queue.data)
	queue.size --
	if queue.size == queue.GetCapacity() / 4 && queue.GetCapacity() / 2 != 0{
		queue.resize(queue.GetCapacity() / 2)
	}
	return
}

//查看队首元素
func (queue *LoopQueue) GetFront() (v interface{},err error) {
	if queue.IsEmpty() {
		err = errors.New("Queue is empty.")
		return
	}
	v = queue.data[queue.front]
	return
}

//格式化输出
func (queue *LoopQueue) PrintIn() string {
	var format string
	format = fmt.Sprintf("Queue: size = %d , capacity = %d\n",queue.GetSize(), queue.GetCapacity())
	format += "front ["
	for i:= queue.front; i != queue.tail; i = (i + 1) % len(queue.data){
		format += fmt.Sprintf("%+v",queue.data[i])
		if ( i + 1) % len(queue.data) != queue.tail {
			format += ", "
		}
	}
	format += "] tail"
	fmt.Println(format)
	return format
}

//队列扩容
func (queue *LoopQueue) resize(newCapacity int) {
	newData := make([]interface{},newCapacity + 1)
	for i:= 0; i < queue.size; i ++ {
		newData[i] = queue.data[(queue.front + i) % len(queue.data)]
	}

	queue.data = newData
	queue.front = 0
	queue.tail = queue.size
	newData = nil
}