/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 17:34
  * Software: GoLand
*/

package linkedlist_stack

import (
	"Play-with-Data-Structures/04-linked-list/06-implement-stack-in-linkedlist/linkedlist"
	"Play-with-Data-Structures/04-linked-list/06-implement-stack-in-linkedlist/stack"
	"fmt"
)

type LinkedListStack struct {
	linkedList *linkedlist.LinkedList
}

//实例化栈
func NewStack() (stack stack.Stack) {
	stack = &LinkedListStack{linkedList:linkedlist.NewLinkedList()}
	return
}

//获取栈元素长度
func (stack *LinkedListStack) GetSize() int {
	return stack.linkedList.GetSize()
}

//判断栈是否为空
func (stack *LinkedListStack) IsEmpty() bool {
	return stack.linkedList.IsEmpty()
}

//向栈头部压入一个元素
func (stack *LinkedListStack) Push(v interface{}) error {
	return stack.linkedList.AddFirst(v)
}

//弹出一个栈头部元素
func (stack *LinkedListStack) Pop() (interface{}, error) {
	return stack.linkedList.RemoveFirst()
}

//查看栈顶元素
func (stack *LinkedListStack) Peek() (interface{}, error) {
	return stack.linkedList.GetFirst()
}

//打印输出
func (stack *LinkedListStack) PrintIn() (format string) {
	format = fmt.Sprintf("Stack: size = %d\n",stack.linkedList.GetSize())
	format += "top ["
	for i:= 0; i< stack.linkedList.GetSize(); i++{
		v, _ := stack.linkedList.Get(i)
		format += fmt.Sprintf("%+v",v)
		if i != stack.GetSize() -1 {
			format += ", "
		}
	}
	format += "]"
	fmt.Println(format)
	return
}


