/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 18:16
  * Software: GoLand
*/

package solution_test

import (
	"testing"
	"Play-with-Data-Structures/04-linked-list/06-implement-stack-in-linkedlist/solution"
)

func TestIsValid(t *testing.T) {
	t.Log(solution.IsValid("()[]{}"))
	t.Log(solution.IsValid("([)]"))
}
