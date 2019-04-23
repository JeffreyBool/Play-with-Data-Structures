/**
  * Author: JeffreyBool
  * Date: 2019/4/24
  * Time: 03:22
  * Software: GoLand
*/

package stack

import (
	"Play-with-Data-Structures/02-array"
	"fmt"
)

type ArrayStack struct {
	array *array.Array
}

/**
 * 实例化栈
 * 时间复杂度 O(1)
 */
func NewStack(capacity ...int) (stack Stack) {
	if len(capacity) > 0 && capacity[0] != 0 {
		stack = &ArrayStack{
			array.NewArray(capacity[0]),
		}
	} else {
		stack = &ArrayStack{
			array:array.NewArray(10),
		}
	}
	return
}

/**
 * 获取栈长度
 * 时间复杂度 O(1)
 */
func (stack *ArrayStack) GetSize() int {
	return stack.array.GetSize()
}

/**
 * 判断栈是否为空
 * 时间复杂度 O(1)
 */
func (stack *ArrayStack) IsEmpty() bool {
	return stack.array.IsEmpty()
}

/**
 * 压入栈顶元素
 * 时间复杂度 O(1)
 */
func (stack *ArrayStack) Push(v interface{}) error {
	return stack.array.AddLast(v)
}

/**
 * 弹出栈顶元素
 * 时间复杂度 O(1)
 */
func (stack *ArrayStack) Pop() (interface{}, error) {
	return stack.array.RemoveLast()
}

/**
 * 查看栈顶元素
 * 时间复杂度 O(1)
 */
func (stack *ArrayStack) Peek() (interface{}, error) {
	return stack.array.GetLast()
}

/**
 * 栈格式化输出
 * 时间复杂度 O(n)
 */
func (stack *ArrayStack) PrintIn() string {
	var format string
	format = fmt.Sprintf("Stack: size = %d , capacity = %d\n",stack.array.GetSize(), stack.array.GetCapacity())
	format += "["
	for i := 0; i < stack.GetSize(); i++{
		if v, err := stack.array.Get(i);err == nil{
			format += fmt.Sprintf("%+v",v)
		}
		if i != stack.GetSize() - 1 {
			format += ", "
		}
	}
	format += "] top"
	fmt.Println(format)
	return format
}
