/**
  * Author: JeffreyBool
  * Date: 2019/4/24
  * Time: 09:43
  * Software: GoLand
*/

package queue

type Queue interface {
	GetSize() int
	IsEmpty() bool
	Enqueue(v interface{}) error
	Dequeue() (interface{},error)
	GetFront() (interface{},error)
	PrintIn() string
}
