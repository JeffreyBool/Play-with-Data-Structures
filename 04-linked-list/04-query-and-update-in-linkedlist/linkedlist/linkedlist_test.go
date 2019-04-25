/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 02:37
  * Software: GoLand
*/

package linkedlist_test

import (
	"testing"
	"Play-with-Data-Structures/04-linked-list/04-query-and-update-in-linkedlist/linkedlist"
)

//实例化链表
func TestNewLinkedList(t *testing.T) {
	linked := linkedlist.NewLinkedList()
	linked.PrintIn()
}

//获取链表长度
func TestLinkedList_GetSize(t *testing.T) {
	linked := linkedlist.NewLinkedList()
	t.Logf("linkedList len: %d \n",linked.GetSize())
}

//判断是否为空
func TestLinkedList_IsEmpty(t *testing.T) {
	linked := linkedlist.NewLinkedList()
	t.Logf("LinkedList empty: %t \n",linked.IsEmpty())
}

//在链表指定索引插入元素
func TestLinkedList_Add(t *testing.T) {
	linked := linkedlist.NewLinkedList()
	for i := 0 ;i < 10; i++{
		if err := linked.Add(i, i+1); err != nil{
			t.Error(err)
			return
		}
	}
	linked.PrintIn()
}

//在链表头部插入元素
func TestLinkedList_AddFirst(t *testing.T) {
	linked := linkedlist.NewLinkedList()
	linked.AddLast("我是链表第一个元素")
	linked.PrintIn()
}

//在链表尾部插入元素
func TestLinkedList_AddLast(t *testing.T) {
	linked := linkedlist.NewLinkedList()
	linked.AddFirst("我是链表第一个元素")
	linked.AddLast("我是链表最后一个元素")
	linked.PrintIn()
}

//根据索引获取链表元素
func TestLinkedList_Get(t *testing.T) {
	strs := []string{"A","B","C","D","E","F"}
	linked := linkedlist.NewLinkedList()
	for _,str := range strs{
		linked.AddFirst(str)
	}
	linked.PrintIn()

	if v, err := linked.Get(0); err != nil{
		t.Error(err)
	}else{
		t.Logf("linked v:%+v \n",v)
	}
}

//获取链表第一个元素
func TestLinkedList_GetFirst(t *testing.T) {
	strs := []string{"A","B","C","D","E","F"}
	linked := linkedlist.NewLinkedList()
	for _,str := range strs{
		linked.AddFirst(str)
	}
	linked.PrintIn()
	v, _ := linked.GetFirst()
	t.Logf("linked first: %+v\n",v)
}

//获取链表第一个元素
func TestLinkedList_GetLast(t *testing.T) {
	strs := []string{"A","B","C","D","E","F"}
	linked := linkedlist.NewLinkedList()
	for _,str := range strs{
		linked.AddFirst(str)
	}
	linked.PrintIn()
	v, _ := linked.GetLast()
	t.Logf("linked first: %+v\n",v)
}

//修改指定索引的数据
func TestLinkedList_Set(t *testing.T) {
	strs := []string{"A","B","C","D","E","F"}
	linked := linkedlist.NewLinkedList()
	for _,str := range strs{
		linked.AddFirst(str)
	}
	linked.PrintIn()
	linked.Set(1,"修改索引数据")
	linked.PrintIn()
}

//判断某个元素是否存在
func TestLinkedList_Contains(t *testing.T) {
	strs := []string{"A","B","C","D","E","F"}
	linked := linkedlist.NewLinkedList()
	for index,str := range strs{
		linked.Add(index,str)
	}
	linked.PrintIn()

	if linked.Contains("D"){
		t.Log("找到了")
	}else{
		t.Log("没找到")
	}
}

//查找某个元素的索引位置
func TestLinkedList_Find(t *testing.T) {
	strs := []string{"A","B","C","D","E","F"}
	linked := linkedlist.NewLinkedList()
	for index,str := range strs{
		linked.Add(index,str)
	}
	linked.PrintIn()

	if index  := linked.Find("E");index != -1 {
		t.Logf("找到了, 索引为: %d \n",index)
	}else{
		t.Log("没找到")
	}
}

//格式化输出
func TestLinkedList_PrintIn(t *testing.T) {
	strs := []string{"A","B","C","D","E","F"}
	linked := linkedlist.NewLinkedList()
	for index,str := range strs{
		linked.Add(index,str)
	}
	linked.PrintIn()
}