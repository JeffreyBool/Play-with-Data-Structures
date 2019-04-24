/**
  * Author: JeffreyBool
  * Date: 2019/4/24
  * Time: 02:46
  * Software: GoLand
*/

package stack

//Stack
type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(v interface{}) error
	Pop() (interface{},error)
	Peek() (interface{},error)
	PrintIn() string
}
