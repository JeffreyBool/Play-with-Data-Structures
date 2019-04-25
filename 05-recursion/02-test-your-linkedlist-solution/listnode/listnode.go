/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 20:51
  * Software: GoLand
*/

package listnode

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
