/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 21:25
  * Software: GoLand
*/

package solution1

import (
	"Play-with-Data-Structures/05-recursion/01-linked-list-problems-in-leetcode/listnode"
)

/// Leetcode 203. Remove Linked List Elements
/// https://leetcode.com/problems/remove-linked-list-elements/description/
func RemoveElements(head *listnode.ListNode,val int) *listnode.ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	if head == nil {
		return nil
	}

	prev := head
	for prev.Next != nil{
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		}else{
			prev = prev.Next
		}
	}

	return head
}
