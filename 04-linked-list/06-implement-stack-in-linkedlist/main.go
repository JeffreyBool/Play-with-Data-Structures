/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 18:31
  * Software: GoLand
*/

package main

import (
	"Play-with-Data-Structures/04-linked-list/06-implement-stack-in-linkedlist/stack"
	"math/rand"
	"time"
	"fmt"
	"Play-with-Data-Structures/04-linked-list/06-implement-stack-in-linkedlist/array_stack"
	"Play-with-Data-Structures/04-linked-list/06-implement-stack-in-linkedlist/linkedlist_stack"
	"sync"
)

func main()  {
	opCount := 10000000
	wg := sync.WaitGroup{}
	wg.Add(2)
	arrayStack := array_stack.NewStack()
	go func() {
		defer wg.Done()
		time1 := testStack(arrayStack, opCount)
		fmt.Println("ArrayStack, time: ", time1, " s")
	}()

	linkedListStack := linkedlist_stack.NewStack()
	go func() {
		defer wg.Done()
		time2 := testStack(linkedListStack,opCount)
		fmt.Println("LinkedList, time: ", time2, " s")
	}()

	wg.Wait()
}

func testStack(stack stack.Stack, opCount int) float64 {
	startTime := time.Now()
	for i := 0; i < opCount; i++ {
		stack.Push(rand.Int())
	}

	for i := 0; i < opCount; i++ {
		stack.Pop()
	}

	return time.Now().Sub(startTime).Seconds()
}