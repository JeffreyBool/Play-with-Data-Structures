/**
  * Author: JeffreyBool
  * Date: 2019/4/24
  * Time: 21:06
  * Software: GoLand
*/
package array

import (
	"errors"
	"fmt"
)

type ArrayInterface interface {
	GetCapacity() int
	GetSize() int
	IsEmpty() bool
	AddFirst(v interface{}) error
	AddLast(v interface{}) error
	Add(index int, v interface{}) error
	GetFirst() (interface{}, error)
	GetLast() (interface{}, error)
	Get(index int) (interface{}, error)
	Set(index int, v interface{}) error
	Contains(v interface{}) bool
	Find(v interface{}) int
	RemoveFirst() (interface{}, error)
	RemoveLast() (interface{}, error)
	Remove(index int) (interface{}, error)
	RemoveElement(v interface{}) (int, error)
	Clear()
	PrintIn() string
}

type Array struct {
	data []interface{}
	size int
}

const size = 10

/**
 * 数组实例化
 * 时间复杂度 O(1)
*/
func NewArray(capacity ...int) (array *Array) {
	if len(capacity) > 0 && capacity[0] != 0 {
		array = &Array{
			data: make([]interface{}, capacity[0]),
		}
	} else {
		array = &Array{
			data: make([]interface{}, size),
		}
	}

	return
}

/**
 * 获取数组容量大小
 * 时间复杂度 O(1)
*/
func (array *Array) GetCapacity() int {
	return cap(array.data)
}

/**
 * 获取数组大小
 * 时间复杂度 O(1)
*/
func (array *Array) GetSize() int {
	return array.size
}

/**
 * 判断数组是否为空
 * 时间复杂度 O(1)
*/
func (array *Array) IsEmpty() bool {
	return array.size == 0
}

/**
 * 根据索引插入元素
 * 时间复杂度 O(m + n)
*/
func (array *Array) Add(index int, v interface{}) (err error) {
	if array.checkIndex(index) {
		err = errors.New("Add failed. Require index >= 0 and index <= size.")
		return
	}

	//因为插入元素，所以在插入的位置数据不动，它后面索引的数据都往后挪动一位。从后往前挪动，如果从前往后挪动会导致后面的数据被覆盖。
	for i := array.size - 1; i >= index; i -- {
		array.data[i+1] = array.data[i]
	}

	//如果数组当前长度等于数组容量就将当前容积扩容 2 倍, 超过1024 大小就扩容 1.25倍，防止数组长度过大，内容不够
	if array.size == array.GetCapacity() {
		var capacity int
		if array.size >= 1024 {
			capacity = array.size * 2 * array.size / 4
		} else {
			capacity = array.size * 2
		}
		array.resize(capacity)
	}

	array.data[index] = v
	array.size ++
	return
}

/**
 * 判断数组是否为空
 * 时间复杂度 O(1)
*/
func (array *Array) AddFirst(v interface{}) error {
	return array.Add(0, v)
}

/**
 * 在数组尾部插入元素
 * 时间复杂度 O(1)
*/
func (array *Array) AddLast(v interface{}) error {
	return array.Add(array.size, v)
}

/**
 * 获取index索引位置的元素
 * 时间复杂度 o(1)
 */
func (array *Array) Get(index int) (v interface{}, err error) {
	if index < 0 || index >= array.size {
		err = errors.New("Get failed. Index is illegal.")
		return
	}
	v = array.data[index]
	return
}

/**
 * 获取数组第一个元素
 * 时间复杂度 o(1)
 */
func (array *Array) GetFirst() (interface{}, error) {
	return array.Get(0)
}

/**
 * 获取数组最后一个元素
 * 时间复杂度 o(1)
 */
func (array *Array) GetLast() (interface{}, error) {
	return array.Get(array.size - 1)
}

/**
 * 修改index索引位置的元素为e
 * 时间复杂度 o(1)
 */
func (array *Array) Set(index int, v interface{}) (err error) {
	if index < 0 || index >= array.size {
		err = errors.New("Set failed. Index is illegal.")
		return
	}
	array.data[index] = v
	return
}

/**
 * 修改index索引位置的元素为e
 * 时间复杂度 o(n)
 */
func (array *Array) Contains(v interface{}) bool {
	for i := 0; i < array.size; i ++ {
		if array.data[i] == v {
			return true
		}
	}
	return false
}

/**
 * 查找数组中元素e所在的索引，如果不存在元素e，则返回-1
 * 时间复杂度 o(n)
 */
func (array *Array) Find(v interface{}) int {
	for i := 0; i < array.size; i ++ {
		if array.data[i] == v {
			return i
		}
	}
	return -1
}

/**
 * 从数组中删除index位置的元素, 返回删除的元素
 * 时间复杂度 o(m+n)
 */
func (array *Array) Remove(index int) (v interface{}, err error) {
	if index < 0 || index >= array.size {
		err = errors.New("Remove failed. Index is illegal.")
		return
	}

	v = array.data[index]
	//解释：因为需要删除 index 索引的数据，所以需要将大于 index 的数据都往前挪动一位。需要从前往后挪动，如果从后往前挪动会覆盖前面的数据。
	for i := index + 1; i < array.size; i ++ {
		array.data[i-1] = array.data[i]
	}

	array.size --
	array.data[array.size] = nil

	//数据空间维护懒模式lazy
	cap := cap(array.data)
	if array.size == cap/4 && cap/2 != 0 {
		array.resize(cap / 2)
	}

	return
}

/**
 * 从数组中删除第一个元素, 返回删除的元素
 * 时间复杂度 o(m+n)
 */
func (array *Array) RemoveFirst() (interface{}, error) {
	return array.Remove(0)
}

/**
 * 从数组中删除最后一个元素, 返回删除的元素
 * 时间复杂度 o(m+n)
 */
func (array *Array) RemoveLast() (interface{}, error) {
	return array.Remove(array.size - 1)
}

/**
 * 删除数组中元素e所在的索引，如果不存在元素e，则返回-1
 * 时间复杂度 o(m+n)
 */
func (array *Array) RemoveElement(v interface{}) (index int, err error) {
	index = array.Find(v)
	if index != -1 {
		_, err = array.Remove(index)
	}
	return
}

/**
 * 清空数组
 * 时间复杂度 o(1)
 */
func (array *Array) Clear() {
	array.data = make([]interface{}, array.size)
	array.size = 0
}

//格式化打印输出
func (array *Array) PrintIn() string {
	var format string
	format = fmt.Sprintf("Array: size = %d , capacity = %d\n", array.size, cap(array.data))
	format += "["
	for i := 0; i < array.GetSize(); i++ {
		format += fmt.Sprintf("%+v", array.data[i])
		if i != array.size-1 {
			format += ", "
		}
	}
	format += "]"
	fmt.Println(format)
	return format
}

//将数组空间的容量变成newCapacity大小
func (array *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < array.size; i++ {
		newData[i] = array.data[i]
	}

	array.data = newData
	newData = nil
}

//索引检查
func (array *Array) checkIndex(index int) bool {
	if index < 0 || index > array.size {
		return true
	}

	return false
}
