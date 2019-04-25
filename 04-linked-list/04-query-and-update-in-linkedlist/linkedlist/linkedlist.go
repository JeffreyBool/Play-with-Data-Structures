/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 02:30
  * Software: GoLand
*/

package linkedlist

import (
	"errors"
	"fmt"
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
	size      int
}

//实例化
func NewLinkedList() *LinkedList {
	return &LinkedList{dummyHead: newNode(nil)}
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
func (linked *LinkedList) Add(index int, v interface{}) (err error) {
	if index < 0 || index > linked.size {
		err = errors.New("add failed. Illegal index.")
		return
	}

	prev := linked.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	prev.next = newNode(v, prev.next)
	linked.size ++
	//node := newNode(v)
	//node.next = prev.next
	//prev.next = node
	return
}

//在链表头部插入元素
func (linked *LinkedList) AddFirst(v interface{}) error {
	return linked.Add(0, v)
}

//在链表末尾添加元素
func (linked *LinkedList) AddLast(v interface{}) error {
	return linked.Add(linked.size, v)
}

// 获得链表的第index(0-based)个位置的元素
// 在链表中不是一个常用的操作，练习用：）
func (linked *LinkedList) Get(index int) (v interface{}, err error) {
	if index < 0 || index >= linked.size || linked.IsEmpty() {
		err = errors.New("Get failed. Illegal index.")
		return
	}
	cur := linked.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	v = cur.data
	return
}

// 获得链表的第一个元素
func (linked *LinkedList) GetFirst() (v interface{}, err error) {
	return linked.Get(0)
}

// 获得链表的最后一个元素
func (linked *LinkedList) GetLast() (v interface{}, err error) {
	return linked.Get(linked.size - 1)
}

func (linked *LinkedList) Set(index int, v interface{}) (err error) {
	if index < 0 || index >= linked.size {
		err = errors.New("set failed. Illegal index.")
		return
	}
	cur := linked.dummyHead.next
	for i:=0;i < index; i++{
		cur = cur.next
	}
	cur.data = v
	return
}

// 查找链表中是否有元素e
func (linked *LinkedList) Contains(v interface{}) bool {
	for cur := linked.dummyHead.next; cur != nil; cur = cur.next{
		if cur.data == v {
			return true
		}
	}
	return false
}

//查找链表指定元素位置，找不到返回 -1
func (linked *LinkedList) Find(v interface{}) int {
	cur := linked.dummyHead.next
	for i := 0; i < linked.size; i++{
		if cur.data == v {
			return i
		}
		cur = cur.next
	}
	return -1
}

//格式化输出
func (linked *LinkedList) PrintIn() (format string) {
	format = fmt.Sprintf("LinkedList: size = %d\n",linked.size)
	format += ""
	for cur := linked.dummyHead.next; cur != nil;cur = cur.next{
		format += fmt.Sprintf("%+v->",cur.data)
	}
	format += "Null"
	fmt.Println(format)
	return
}
