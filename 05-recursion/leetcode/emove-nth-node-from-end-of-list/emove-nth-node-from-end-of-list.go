/**
  * Author: JeffreyBool
  * Date: 2019/4/26
  * Time: 01:24
  * Software: GoLand
*/

package emove_nth_node_from_end_of_list

import (
	"fmt"
	"errors"
	"bytes"
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
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}

	return
}

func (linked *ListNode) PrintIn() {
	cur := linked
	var buffer bytes.Buffer
	for cur != nil {
		buffer.WriteString(fmt.Sprintf("%v->", cur.Val))
		cur = cur.Next
	}
	buffer.WriteString("Null")
	fmt.Println(buffer.String())
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Val:-1}
	dummyHead.Next = head
	second := dummyHead
	first := dummyHead.Next
	for i:= 1;i<= n;i++{
		first = first.Next
	}
	for first != nil{
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummyHead.Next
}
