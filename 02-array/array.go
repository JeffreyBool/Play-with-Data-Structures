/**
  * Author: JeffreyBool
  * Date: 2019/4/23
  * Time: 23:30
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
		array.data[i+1] = array.data
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
	return array.Get(array.size -1)
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

func (array *Array) Set(index int, v interface{}) error {
	panic("implement me")
}

func (array *Array) Contains(v interface{}) bool {
	panic("implement me")
}

func (array *Array) Find(v interface{}) int {
	panic("implement me")
}

func (array *Array) RemoveFirst() (interface{}, error) {
	panic("implement me")
}

func (array *Array) RemoveLast() (interface{}, error) {
	panic("implement me")
}

func (array *Array) Remove(index int) (interface{}, error) {
	panic("implement me")
}

func (array *Array) RemoveElement(v interface{}) (int, error) {
	panic("implement me")
}

func (array *Array) Clear() {
	panic("implement me")
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
		newData[i] = array.data[9]
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
