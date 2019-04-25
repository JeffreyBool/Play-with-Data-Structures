/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 19:05
  * Software: GoLand
*/

package linkedlist_queue

import (
	"Play-with-Data-Structures/04-linked-list/07-implement-queue-in-linkedlist/queue"
	"errors"
	"fmt"
)

type node struct {
	data interface{}
	next *node
}

func newNode(v interface{}, next ...*node) (n *node) {
	if len(next) > 0 && next[0] != nil {
		n = &node{data: v, next: next[0]}
	} else {
		n = &node{data: v}
	}
	return
}

func (node *node) GetNext() *node {
	return node.next
}

func (node *node) GetValue() interface{} {
	return node.data
}

type LinkedListQueue struct {
	head *node
	tail *node
	size int
}

func NewQueue() queue.Queue {
	return &LinkedListQueue{}
}

func (queue *LinkedListQueue) GetSize() int {
	return queue.size
}

func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *LinkedListQueue) Enqueue(v interface{}) error {
	//如果 tail 等于nil ,说明链表就是空
	if queue.tail == nil{
		queue.tail = newNode(v)
		queue.head = queue.tail
	}else{
		queue.tail.next = newNode(v)
		queue.tail = queue.tail.next
		queue.size ++
	}
	return nil
}

func (queue *LinkedListQueue) Dequeue() (v interface{},err error) {
	if queue.IsEmpty() {
		err = errors.New("Cannot dequeue from an empty queue.")
		return
	}
	v = queue.head.data
	oldNode := queue.head
	queue.head = queue.head.next
	oldNode.next = nil
	if queue.head == nil{
		queue.tail = nil
	}
	queue.size --
	return
}

func (queue *LinkedListQueue) GetFront() (v interface{},err error) {
	if queue.IsEmpty() {
		err = errors.New("Cannot dequeue from an empty queue.")
		return
	}
	v = queue.head.data
	return
}

func (queue *LinkedListQueue) PrintIn() (format string) {
	format += fmt.Sprintf("Queue: size = %d\n",queue.size)
	format += "front ["
	for cur := queue.head; cur != nil; cur = cur.next{
		format += fmt.Sprintf("%+v",cur.data)
		if cur.next != nil{
			format += ", "
		}
	}
	format += "] tail"
	fmt.Println(format)
	return
}



