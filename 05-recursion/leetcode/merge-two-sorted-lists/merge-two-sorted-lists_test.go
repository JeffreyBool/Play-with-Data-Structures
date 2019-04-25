/**
  * Author: JeffreyBool
  * Date: 2019/4/26
  * Time: 02:18
  * Software: GoLand
*/

package merge_two_sorted_lists_test

import (
	"testing"
	"Play-with-Data-Structures/05-recursion/leetcode/merge-two-sorted-lists"
)

func TestMergeTwoLists(t *testing.T) {

	node1, _ := merge_two_sorted_lists.NewListNode([]int{1,3,4,5,9})
	node2, _ := merge_two_sorted_lists.NewListNode([]int{1,2,5,6,6})
	lists := merge_two_sorted_lists.MergeTwoLists(node1, node2)
	lists.PrintIn()
}
