/**
  * Author: JeffreyBool
  * Date: 2019/4/24
  * Time: 09:49
  * Software: GoLand
*/

package queue

import (
	"Play-with-Data-Structures/02-array"
	"fmt"
)

// 数组队列的局限性：出队列的操作时间复杂度为 n
type ArrayQueue struct {
	array *array.Array
}

/**
 * 队列实例化
 * 时间复杂度 O(1)
 */
func NewQueue(capacity ...int) (queue Queue) {
	if len(capacity) > 0 && capacity[0] != 0 {
		queue = &ArrayQueue{
			array: array.NewArray(capacity[0]),
		}
	} else {
		queue = &ArrayQueue{
			array: array.NewArray(10),
		}
	}
	return
}

/**
 * 获取队列长度
 * 时间复杂度 O(1)
 */
func (queue *ArrayQueue) GetSize() int {
	return queue.array.GetSize()
}

/**
 * 判断队列是否为空
 * 时间复杂度 O(1)
 */
func (queue *ArrayQueue) IsEmpty() bool {
	return queue.array.IsEmpty()
}

/**
 * 向队尾插入元素
 * 时间复杂度 O(1)
 */
func (queue *ArrayQueue) Enqueue(v interface{}) error {
	return queue.array.AddLast(v)
}

/**
 * 弹出队首元素
 * 时间复杂度 O(n)
 */
func (queue *ArrayQueue) Dequeue() (interface{}, error) {
	return queue.array.RemoveFirst()
}

/**
 * 查看队尾元素
 * 时间复杂度 O(1)
 */
func (queue *ArrayQueue) GetFront() (interface{}, error) {
	return queue.array.GetFirst()
}

func (queue *ArrayQueue) PrintIn() string {
	var format string
	format = fmt.Sprintf("Queue: size = %d , capacity = %d\n",queue.array.GetSize(), queue.array.GetCapacity())
	format += "front ["
	for i:= 0; i < queue.GetSize(); i++{
		v,_ := queue.array.Get(i)
		format += fmt.Sprint(v)
		if i != queue.GetSize() -1 {
			format += ", "
		}
	}
	format += "] tail"
	fmt.Println(format)
	return format
}
