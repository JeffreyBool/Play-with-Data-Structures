/**
  * Author: JeffreyBool
  * Date: 2019/4/24
  * Time: 03:44
  * Software: GoLand
*/

package stack_test

import (
	"testing"
	"Play-with-Data-Structures/03-stacks-and-queue/stack"
)

//实例化
func TestNewStack(t *testing.T) {
	var arrayStack stack.Stack
	arrayStack = stack.NewStack()
	arrayStack.PrintIn()
}

//获取栈长度
func TestArrayStack_GetSize(t *testing.T) {
	var arrayStack stack.Stack
	arrayStack = stack.NewStack()
	t.Logf("stack len: %d \n",arrayStack.GetSize())
}

//判断栈是否为空
func TestArrayStack_IsEmpty(t *testing.T) {
	var arrayStack stack.Stack
	arrayStack = stack.NewStack()
	t.Logf("stack empty: %t \n",arrayStack.IsEmpty())
}

//压入栈顶元素
func TestArrayStack_Push(t *testing.T) {
	var arrayStack stack.Stack
	arrayStack = stack.NewStack()
	if err := arrayStack.Push("你好呀？"); err != nil{
		t.Error(err)
	}

	arrayStack.PrintIn()
}

//弹出栈顶元素
func TestArrayStack_Pop(t *testing.T) {
	var arrayStack stack.Stack
	arrayStack = stack.NewStack()
	if err := arrayStack.Push("你好呀？"); err != nil{
		t.Error(err)
	}
	if err := arrayStack.Push("hello"); err != nil{
		t.Error(err)
	}
	arrayStack.PrintIn()
	if pop, err := arrayStack.Pop(); err != nil{
		t.Error(err)
	}else{
		t.Log(pop)
	}
}

//查看栈顶元素
func TestArrayStack_Peek(t *testing.T) {
	var arrayStack stack.Stack
	arrayStack = stack.NewStack()
	strs := []string{"A","B","C","D","E","F"}
	for _,str := range strs{
		arrayStack.Push(str)
	}
	if peek, err := arrayStack.Peek();err != nil{
		t.Error(err)
	}else{
		t.Logf("stack peek: %+v \n",peek)
	}

	arrayStack.PrintIn()
}

//格式化输出
func TestArrayStack_PrintIn(t *testing.T) {
	var arrayStack stack.Stack
	arrayStack = stack.NewStack()
	strs := []string{"A","B","C","D","E","F"}
	for _,str := range strs{
		arrayStack.Push(str)
	}

	arrayStack.PrintIn()
}
