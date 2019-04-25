/**
  * Author: JeffreyBool
  * Date: 2019/4/24
  * Time: 04:03
  * Software: GoLand
*/

package main

import (
	"fmt"
	"Play-with-Data-Structures/03-stacks-and-queue/stack"
)

//有效括号检查 https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	var arrayStack stack.Stack
	arrayStack = stack.NewStack()
	brackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _,char := range []rune(s){
		if char == '{' || char == '[' || char == '(' {
			arrayStack.Push(char)
		}else{
			if arrayStack.IsEmpty() {return false}
			v,_ := arrayStack.Pop()
			topChar := v.(rune)
			if  val,ok := brackets[char]; ok && topChar != val{
				return false
			}
		}
	}

	return arrayStack.IsEmpty()
}

// go 的字符串实际是 byte 类型组成的切片
func isVaildRune(s string) bool {
	var stack []rune
	brackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, char := range []rune(s) {
		if char == '(' || char == '[' || char == '{' {
			// 括号左半部直接入栈
			stack = append(stack, char)
		}else if len(stack) > 0 && stack[len(stack) -1 ] == brackets[char] {
			// 括号右半部：栈中有数据，且此元素与栈尾元素相同，出栈
			stack = stack[:len(stack)-1]
		}else{
			return false
		}
	}

	return len(stack) == 0
}

func main()  {
	fmt.Println(isVaildRune("([])"))
}
