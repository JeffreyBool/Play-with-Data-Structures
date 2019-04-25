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

// https://leetcode.com/problems/valid-parentheses/description/
// 括号匹配问题
//
// 使用LinkedListStack测试20号问题的代码在视频中没有提及
// 该代码主要用于使用Leetcode上的问题测试我们的LinkedListStack：）
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


