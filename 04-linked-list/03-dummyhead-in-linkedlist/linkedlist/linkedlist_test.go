/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 02:37
  * Software: GoLand
*/

package linkedlist_test

import (
	"testing"
	"Play-with-Data-Structures/04-linked-list/02-add-elements-in-linkedlist/linkedlist"
)

func TestNewLinkedList(t *testing.T) {
	linked := linkedlist.NewLinkedList()
	t.Logf("linked: %+v\n",linked)
}