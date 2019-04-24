/**
  * Author: JeffreyBool
  * Date: 2019/4/24
  * Time: 17:16
  * Software: GoLand
*/

package loop_queue

import (
	"Play-with-Data-Structures/03-stacks-and-queue/05-array-queue"
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
	return cap(queue.data) -1
}

//判断队列是否为空
func (queue *LoopQueue) IsEmpty() bool {
	return queue.front == queue.tail
}

func (queue *LoopQueue) Enqueue(v interface{}) (err error) {
	panic("implement me")
}

func (queue *LoopQueue) Dequeue() (interface{}, error) {
	panic("implement me")
}

func (queue *LoopQueue) GetFront() (interface{}, error) {
	panic("implement me")
}

func (queue *LoopQueue) PrintIn() string {
	panic("implement me")
}
