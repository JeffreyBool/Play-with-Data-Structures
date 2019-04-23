/**
  * Author: JeffreyBool
  * Date: 2019/4/23
  * Time: 23:30
  * Software: GoLand
*/

package array

type ArrayInterface interface {
	GetCapacity() int
	GetSize() int
	IsEmpty() bool
	AddFirst(v interface{}) error
	AddLast(v interface{}) error
	Add(index int,v interface{}) error
	GetFirst() (interface{},error)
	GetLast() (interface{},error)
	Get(index int) (interface{},error)
	Set(index int, v interface{}) error
	Contains(v interface{}) bool
	Find(v interface{}) int
	RemoveFirst() (interface{},error)
	RemoveLast() (interface{},error)
	Remove(index int) (interface{},error)
	RemoveElement(v interface{}) (int,error)
	Clear()
	PrintIn() string
}

type Array struct {
	data []interface{}
	size int
}

func NewArray(capacity ...int)  {

}

func (array *Array) GetCapacity() int {
	panic("implement me")
}

func (array *Array) GetSize() int {
	panic("implement me")
}

func (array *Array) IsEmpty() bool {
	panic("implement me")
}

func (array *Array) AddFirst(v interface{}) error {
	panic("implement me")
}

func (array *Array) AddLast(v interface{}) error {
	panic("implement me")
}

func (array *Array) Add(index int, v interface{}) error {
	panic("implement me")
}

func (array *Array) GetFirst() (interface{}, error) {
	panic("implement me")
}

func (array *Array) GetLast() (interface{}, error) {
	panic("implement me")
}

func (array *Array) Get(index int) (interface{}, error) {
	panic("implement me")
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

func (array *Array) PrintIn() string {
	panic("implement me")
}


