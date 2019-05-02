/**
  * Author: JeffreyBool
  * Date: 2019/4/26
  * Time: 21:58
  * Software: GoLand
*/

package linkedlist_sort_list

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


func SortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val:-1}
	dummyHead.Next = head

	flst := dummyHead.Next
	slow := dummyHead
	chanerrors := make(chan error, 0)
	for flst != nil {
		flst = flst.Next
	}

	return dummyHead.Next
}
