/**
  * Author: JeffreyBool
  * Date: 2019/4/26
  * Time: 01:37
  * Software: GoLand
*/

package emove_nth_node_from_end_of_list_test

import (
	"testing"
	"Play-with-Data-Structures/05-recursion/leetcode/emove-nth-node-from-end-of-list"
)

func TestRemoveNthFromEnd(t *testing.T) {
	listNode, _ := emove_nth_node_from_end_of_list.NewListNode([]int{1, 2, 3, 4, 5})
	listNode.PrintIn()
	emove_nth_node_from_end_of_list.RemoveNthFromEnd(listNode,3)
	listNode.PrintIn()
}
