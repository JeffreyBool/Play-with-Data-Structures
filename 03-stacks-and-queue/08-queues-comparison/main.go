/**
  * Author: JeffreyBool
  * Date: 2019/4/24
  * Time: 22:40
  * Software: GoLand
*/

package main

import (
	"Play-with-Data-Structures/03-stacks-and-queue/08-queues-comparison/queue"
	"time"
	"math/rand"
	"Play-with-Data-Structures/03-stacks-and-queue/08-queues-comparison/array_queue"
	"fmt"
	"sync"
	"Play-with-Data-Structures/03-stacks-and-queue/08-queues-comparison/loop_queue"
)

func TestQueue(queue queue.Queue, opCount int) float64 {
	startTime := time.Now()
	for i := 0; i < opCount; i++ {
		queue.Enqueue(rand.Int())
	}
	queue.PrintIn()
	for i := 0; i < opCount; i++ {
		queue.Dequeue()
	}

	queue.PrintIn()
	return time.Now().Sub(startTime).Seconds()
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	opCount := 100000
	go func() {
		defer wg.Done()
		arrayQueue := array_queue.NewQueue(100000)
		sumTime := TestQueue(arrayQueue, opCount)
		fmt.Println("ArrayQueue, time:", sumTime, "s")
	}()

	go func() {
		defer wg.Done()
		LoopQueue := loop_queue.NewQueue(100000)
		sumTime := TestQueue(LoopQueue, opCount)
		fmt.Println("LoopQueue, time:", sumTime, "s")
	}()
	wg.Wait()

	//Queue: size = 0 , capacity = 0
	//front [] tail
	//LoopQueue, time: 0.019569912 s
	//Queue: size = 0 , capacity = 1
	//front [] tail
	//ArrayQueue, time: 7.335745758 s
}

