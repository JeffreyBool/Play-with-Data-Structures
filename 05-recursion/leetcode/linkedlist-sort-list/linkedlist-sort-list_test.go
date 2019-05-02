/**
  * Author: JeffreyBool
  * Date: 2019/4/26
  * Time: 21:59
  * Software: GoLand
*/

package linkedlist_sort_list_test

import (
	"testing"
	"Play-with-Data-Structures/05-recursion/leetcode/linkedlist-sort-list"
)

func TestSortList(t *testing.T) {
	list, _ := linkedlist_sort_list.NewListNode([]int{1,3,4,5,9})
	list.PrintIn()
	list = linkedlist_sort_list.SortList(list)
	list.PrintIn()
}
