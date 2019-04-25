/**
  * Author: JeffreyBool
  * Date: 2019/4/26
  * Time: 02:14
  * Software: GoLand
*/

package merge_two_sorted_lists

import (
	"errors"
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(arr []int) (node *ListNode, err error) {
	if len(arr) == 0 {
		err = errors.New("arr can not be empty")
		return
	}
	node = &ListNode{Val:arr[0]}
	cur := node
	for i:= 1; i< len(arr); i++{
		cur.Next = &ListNode{Val:arr[i]}
		cur = cur.Next
	}

	return
}

func (linked *ListNode) PrintIn() {
	cur := linked
	var buffer bytes.Buffer
	for cur != nil{
		buffer.WriteString(fmt.Sprintf("%v->", cur.Val))
		cur = cur.Next
	}

	buffer.WriteString("Null")
	fmt.Println(buffer.String())
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil{
		return l1
	}

	// 类似归并排序中的合并过程
	dummyHead := &ListNode{}
	cur := dummyHead
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}else{
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}

	// 任一为空，直接连接另一条链表
	if l1 == nil{
		cur.Next = l2
	}else{
		cur.Next = l1
	}

	return dummyHead.Next
}
