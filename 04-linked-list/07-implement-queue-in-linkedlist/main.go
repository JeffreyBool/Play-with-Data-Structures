/**
  * Author: JeffreyBool
  * Date: 2019/4/25
  * Time: 20:18
  * Software: GoLand
*/

package main

import (
	"Play-with-Data-Structures/04-linked-list/07-implement-queue-in-linkedlist/queue"
	"math/rand"
	"time"
	"fmt"
	"sync"
	"Play-with-Data-Structures/04-linked-list/07-implement-queue-in-linkedlist/array_queue"
	"Play-with-Data-Structures/04-linked-list/07-implement-queue-in-linkedlist/loop_queue"
	"Play-with-Data-Structures/04-linked-list/07-implement-queue-in-linkedlist/linkedlist_queue"
)

func main()  {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	opCount := 100000

	go func() {
		defer wg.Done()
		arrayQueue := array_queue.NewQueue(100000)
		sumTime := testQueu(arrayQueue, opCount)
		fmt.Println("ArrayQueue, time:", sumTime, "s")
	}()

	go func() {
		defer wg.Done()
		LoopQueue := loop_queue.NewQueue(100000)
		sumTime := testQueu(LoopQueue, opCount)
		fmt.Println("LoopQueue, time:", sumTime, "s")
	}()

	go func() {
		defer wg.Done()
		linkedListQueue := linkedlist_queue.NewQueue()
		sumTime := testQueu(linkedListQueue, opCount)
		fmt.Println("linkedListQueue, time:", sumTime, "s")
	}()

	wg.Wait()

	//LoopQueue, time: 0.028870041 s
	//linkedListQueue, time: 0.036870903 s
	//ArrayQueue, time: 7.182285503 s
}

func testQueu(queue queue.Queue, opCount int) float64 {
	startTime := time.Now()
	for i:= 0; i< opCount;i++{
		queue.Enqueue(rand.Int())
	}
	for i:= 0; i< opCount; i++{
		queue.Dequeue()
	}
	return time.Now().Sub(startTime).Seconds()
}
