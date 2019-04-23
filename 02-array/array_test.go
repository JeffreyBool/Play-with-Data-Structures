/**
  * Author: JeffreyBool
  * Date: 2019/4/24
  * Time: 00:49
  * Software: GoLand
*/

package array_test

import (
	"testing"
	"Play-with-Data-Structures/02-array"
)

//数组初始化
func TestNewArray(t *testing.T) {
	arr := array.NewArray()
	arr.PrintIn()
}


func TestArray_Add(t *testing.T) {
	arr := array.NewArray()
	for i := 0; i < 10; i++{
		if err := arr.Add(i, i+1); err != nil{
			t.Error(err)
			break
		}
	}

	arr.PrintIn()
	if err := arr.Add(10, 11); err != nil{
		t.Error(err)
	}

	arr.PrintIn()
}