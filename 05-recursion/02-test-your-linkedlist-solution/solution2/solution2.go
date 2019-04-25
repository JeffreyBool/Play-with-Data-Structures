/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 21:41
  * Software: GoLand
*/

package solution2

import (
	"Play-with-Data-Structures/05-recursion/02-test-your-linkedlist-solution/listnode"
)

/// Leetcode 203. Remove Linked List Elements
/// https://leetcode.com/problems/remove-linked-list-elements/description/
func RemoveElements(head *listnode.ListNode,val int) *listnode.ListNode {
	dummyHead := &listnode.ListNode{}
	dummyHead.Next = head

	prev := dummyHead
	for prev.Next != nil{
		if prev.Next.Val == val{
			prev.Next = prev.Next.Next
		}else{
			prev = prev.Next
		}
	}

	return dummyHead.Next
}
