/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 21:01
  * Software: GoLand
*/

package main

import (
	"Play-with-Data-Structures/05-recursion/01-linked-list-problems-in-leetcode/listnode"
	"fmt"
	"Play-with-Data-Structures/05-recursion/01-linked-list-problems-in-leetcode/solution2"
)

func main() {
	listNode := &listnode.ListNode{Val: 1, Next: &listnode.ListNode{Val: 2, Next: &listnode.ListNode{Val: 6, Next: &listnode.ListNode{Val: 3, Next: &listnode.ListNode{Val: 4, Next: &listnode.ListNode{Val: 5, Next: &listnode.ListNode{Val: 6}}}}}}}
	listNode = solution2.RemoveElements(listNode, 6)

	var format string
	for cur := listNode; cur != nil; cur = cur.Next {
		format += fmt.Sprintf("%+v->", cur.Val)
	}
	format += "Null"
	fmt.Println(format)
}
