/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 20:52
  * Software: GoLand
*/

package solution

import (
	"Play-with-Data-Structures/05-recursion/01-linked-list-problems-in-leetcode/listnode"
)

/// Leetcode 203. Remove Linked List Elements
/// https://leetcode.com/problems/remove-linked-list-elements/description/
func RemoveElements(head *listnode.ListNode,val int) *listnode.ListNode {
	for head != nil && head.Val == val {
		cur := head
		head = head.Next
		cur.Next = nil
	}

	if head == nil{
		return nil
	}

	prev := head
	for prev.Next != nil{
		if prev.Next.Val == val{
			cur := prev.Next
			prev.Next = cur.Next
			cur.Next = nil
		}else{
			prev = prev.Next
		}
	}

	return head
}
