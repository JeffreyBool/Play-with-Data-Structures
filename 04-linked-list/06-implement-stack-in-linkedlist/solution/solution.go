/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 18:07
  * Software: GoLand
*/

package solution

import (
	implement "Play-with-Data-Structures/04-linked-list/06-implement-stack-in-linkedlist/stack"
	"Play-with-Data-Structures/04-linked-list/06-implement-stack-in-linkedlist/linkedlist_stack"
)

var brackets = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func IsValid(s string) bool {
	var stack implement.Stack
	stack = linkedlist_stack.NewStack()
	for _,char := range s {
		// 括号左半部直接入栈
		if char == '(' || char == '{' || char == '[' {
			stack.Push(char)
		}else{
			peek, _ := stack.Peek()
			if stack.GetSize() > 0 && peek == brackets[char] {
				stack.Pop()
			}else{
				return false
			}
		}
	}

	return stack.IsEmpty()
}


