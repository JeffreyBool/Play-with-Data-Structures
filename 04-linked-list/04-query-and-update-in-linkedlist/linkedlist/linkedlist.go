/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 02:30
  * Software: GoLand
*/

package linkedlist

import (
	"errors"
)

type node struct {
	data interface{}
	next *node
}

//实例化
func newNode(v interface{}, next ...*node) (n *node) {
	if len(next) > 0 && next[0] != nil {
		n = &node{data: v, next: next[0]}
	} else {
		n = &node{data: v}
	}
	return
}

type LinkedList struct {
	dummyHead *node
	size int
}

//实例化
func NewLinkedList() *LinkedList {
	return &LinkedList{dummyHead:newNode(nil)}
}

//获取链表元素长度
func (linked *LinkedList) GetSize() int {
	return linked.size
}

//判断链表是否为空
func (linked *LinkedList) IsEmpty() bool {
	return linked.size == 0
}

// 在链表的index(0-based)位置添加新的元素e
// 在链表中不是一个常用的操作，练习用：）
func (linked *LinkedList) Add(index int,v interface{}) (err error) {
	if index < 0 || index > linked.size {
		err = errors.New("add failed. Illegal index.")
		return
	}

	prev := linked.dummyHead
	for i:= 0; i < index; i++{
		prev = prev.next
	}

	prev.next = newNode(v,prev.next)
	linked.size ++
	//node := newNode(v)
	//node.next = prev.next
	//prev.next = node
	return
}

//在链表头部插入元素
func (linked *LinkedList) AddFirst(v interface{}) error {
	return linked.Add(0,v)
}

//在链表末尾添加元素
func (linked *LinkedList) AddLast(v interface{}) error {
	return linked.Add(linked.size,v)
}

