/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 17:47
  * Software: GoLand
*/

package linkedlist_stack_test

import (
	"testing"
	Implement "Play-with-Data-Structures/04-linked-list/06-implement-stack-in-linkedlist/stack"
	"Play-with-Data-Structures/04-linked-list/06-implement-stack-in-linkedlist/linkedlist_stack"
)

//实例化
func TestNewStack(t *testing.T) {
	var stack Implement.Stack
	stack = linkedlist_stack.NewStack()
	stack.PrintIn()
}

//获取栈长度
func TestLinkedListStack_GetSize(t *testing.T) {
	var stack Implement.Stack
	stack = linkedlist_stack.NewStack()
	t.Logf("stack len: %d \n",stack.GetSize())
}

//判断栈是否为空
func TestLinkedListStack_IsEmpty(t *testing.T) {
	var stack Implement.Stack
	stack = linkedlist_stack.NewStack()
	t.Logf("stack empty: %t \n",stack.IsEmpty())
}

//压入栈顶元素
func TestLinkedListStack_Push(t *testing.T) {
	var stack Implement.Stack
	stack = linkedlist_stack.NewStack()
	if err := stack.Push("你好呀？"); err != nil{
		t.Error(err)
	}

	stack.PrintIn()
}

//弹出栈顶元素
func TestLinkedListStack_Pop(t *testing.T) {
	var stack Implement.Stack
	stack = linkedlist_stack.NewStack()
	if err := stack.Push("你好呀？"); err != nil{
		t.Error(err)
	}
	if err := stack.Push("hello"); err != nil{
		t.Error(err)
	}
	stack.PrintIn()
	if pop, err := stack.Pop(); err != nil{
		t.Error(err)
	}else{
		t.Log(pop)
	}
	stack.PrintIn()
}

//查看栈顶元素
func TestLinkedListStack_Peek(t *testing.T) {
	var stack Implement.Stack
	stack = linkedlist_stack.NewStack()
	strs := []string{"A","B","C","D","E","F"}
	for _,str := range strs{
		stack.Push(str)
	}
	if peek, err := stack.Peek();err != nil{
		t.Error(err)
	}else{
		t.Logf("stack peek: %+v \n",peek)
	}

	stack.PrintIn()
}

//格式化输出
func TestLinkedListStack_PrintIn(t *testing.T) {
	var stack Implement.Stack
	stack = linkedlist_stack.NewStack()
	strs := []string{"A","B","C","D","E","F"}
	for _,str := range strs{
		stack.Push(str)
	}

	stack.PrintIn()
}