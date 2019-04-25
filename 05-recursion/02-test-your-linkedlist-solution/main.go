/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 21:01
  * Software: GoLand
*/

package main

import (
	"Play-with-Data-Structures/05-recursion/02-test-your-linkedlist-solution/listnode"
	"Play-with-Data-Structures/05-recursion/02-test-your-linkedlist-solution/solution2"
)

func main() {
	//listNode := &listnode.ListNode{Val: 1, Next: &listnode.ListNode{Val: 2, Next: &listnode.ListNode{Val: 6, Next: &listnode.ListNode{Val: 3, Next: &listnode.ListNode{Val: 4, Next: &listnode.ListNode{Val: 5, Next: &listnode.ListNode{Val: 6}}}}}}}
	//listNode = solution2.RemoveElements(listNode, 6)
	//
	//var format string
	//for cur := listNode; cur != nil; cur = cur.Next {
	//	format += fmt.Sprintf("%+v->", cur.Val)
	//}
	//format += "Null"
	//fmt.Println(format)

	nums := []int{1, 2, 6, 3, 4, 5, 6}
	head, _ := listnode.NewListNode(nums)
	head.PrintIn()
	head = solution2.RemoveElements(head, 6)
	head.PrintIn()
}
